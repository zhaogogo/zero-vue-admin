<template>
  <div>
    <div class="button-box clearflex">
        <el-input style="width:30%" placeholder="搜索用户" v-model="searchInfo.nameX" clearable @change="searchUser"></el-input>
        <el-button size="mini" type="primary" icon="el-icon-plus" @click="addUser">新增用户</el-button>
    </div>
    <el-table :data="tableData" border stripe row-key="id" @cell-click="rowClick">
        <el-table-column label="ID"  prop="id"></el-table-column>
        <el-table-column label="用户名"  prop="name" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column label="昵称"  prop="nick_name" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column label="密码"  prop="password" :show-overflow-tooltip="true">
            <template slot-scope="scope">
                <span v-if="!isEdit[scope.row.id]">{{scope.row.password}}</span>
                <el-input ref="inputpassword" v-if="isEdit[scope.row.id]" v-model="scope.row.password" type="password" @blur="changePassword(scope.row)"></el-input>
            </template>
        </el-table-column>
        <el-table-column label="用户角色" prop="role" min-width="170">
            <template slot-scope="scope">
                <el-cascader
                    v-model="scope.row.roleList"
                    :options="roleOptions"
                    :props="{ checkStrictly: true,label:'roleName',value:'roleid',disabled:'disabled', emitPath:false, multiple: true}"
                    clearable
                    filterable
                    @change="changeRole(scope.row)"
                >
                </el-cascader>
            </template>
        </el-table-column>
        <el-table-column label="类型"  prop="type">
            <template slot-scope="scope">
                <el-tag v-show="(scope.row.type === 0)">本地用户</el-tag>
                <el-tag v-show="(scope.row.type === 1)" type="success">LDAP用户</el-tag>
            </template>
        </el-table-column>
        <el-table-column label="邮箱"  prop="email" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column label="电话"  prop="phone" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column label="部门"  prop="department" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column label="职位"  prop="position" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column label="创建人"  prop="createBy" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column label="创建时间"  prop="create_time" :show-overflow-tooltip="true">
            <template slot-scope="scope">
                <div>{{parseTimeStamp(scope.row.create_time)}}</div>
            </template>
        </el-table-column>
        <el-table-column label="修改人"  prop="updateBy" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column label="修改时间"  prop="update_time" :show-overflow-tooltip="true">
            <template slot-scope="scope">
                <div>{{parseTimeStamp(scope.row.update_time)}}</div>
            </template>
        </el-table-column>
        <el-table-column label="删除人"  prop="deleteBy" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column label="软删除"  prop="delete_time">
            <template slot-scope="scope">
                <el-tooltip :content="scope.row.delete_time<=0?'active':parseTimeStamp(scope.row.delete_time)" placement="top">
                    <el-switch
                        v-model="scope.row.state"
                        active-color="#13ce66"
                        inactive-color="#ff4949"
                        @change="softDeleteUser($event,scope.row,scope)"
                        active-value="deleted"
                        inactive-value="resume"
                    >
                    </el-switch>
                </el-tooltip>
            </template>
        </el-table-column>
        <el-table-column label="操作" min-width="150" fixed="right">
            <template slot-scope="scope">
                <el-popover
                    placement="top"
                    width="160"
                    trigger="hover"
                    v-model="scope.row.visible"
                >
                    <p>确定要删除此用户吗</p>
                    <div style="text-align: right; margin: 0">
                        <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
                        <el-button size="mini" type="primary" @click="deleteUser(scope.row)">确定</el-button>
                    </div>
                    <el-button size="mini" type="text" icon="el-icon-delete" slot="reference">删除</el-button>
                </el-popover>
                <el-button size="mini" type="text" icon="el-icon-edit" @click="editUser(scope.row)">编辑</el-button>
            </template>
        </el-table-column>
    </el-table>
    <el-pagination
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[10,30,50,100]"
        :style="{float:'right',padding:'20px'}"
        :total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="handlePageChange"
        @size-change="handlePageSizeChange"
    ></el-pagination>

    <el-dialog :visible.sync="dialogVisible" custom-class="user-dialog" :title="dialogTitle" :before-close="closeDialog">
        <el-form ref="userInfoFrom" :rules="rulesadduser" :model="userInfo">
            <el-form-item label="用户ID" label-width="80px" prop="id">
                <el-input v-model.number="userInfo.id" :disabled="true"></el-input>
            </el-form-item>
            <el-form-item label="用户名" label-width="80px" prop="name">
                <el-input v-model="userInfo.name"></el-input>
            </el-form-item>
            <el-form-item label="昵称" label-width="80px" prop="nick_name">
                <el-input v-model="userInfo.nick_name"></el-input>
            </el-form-item>
            <el-form-item v-if="dialogFlag === 'adduser'" label="用户类型" prop="type">
                <el-radio-group v-model.number="userInfo.type">
                    <el-radio :label="0">本地用户</el-radio>
                    <el-radio :label="1">LDAP用户</el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item v-if="dialogFlag === 'adduser'" label="密码" label-width="80px" prop="password">
                <el-input show-password v-model="userInfo.password"></el-input>
            </el-form-item>
            <el-form-item v-if="dialogFlag === 'adduser'" label="确认密码" label-width="80px" prop="confirmPassword">
                <el-input show-password v-model="userInfo.confirmPassword"></el-input>
            </el-form-item>
            <el-form-item v-if="dialogFlag === 'adduser'" label="角色" label-width="80px" prop="roleList">
                <el-cascader
                    v-model="userInfo.roleList"
                    :options="roleOptions"
                    :props="{ checkStrictly: true,label:'roleName',value:'roleid',disabled:'disabled', emitPath:false, multiple: true}"
                    clearable
                    filterable
                >
                </el-cascader>
            </el-form-item>
            <el-form-item label="邮箱" label-width="80px" prop="email">
                <el-input v-model="userInfo.email"></el-input>
            </el-form-item>
            <el-form-item label="电话" label-width="80px" prop="phone">
                <el-input v-model.number="userInfo.phone"></el-input>
            </el-form-item>
            <el-form-item label="部门" label-width="80px" prop="department">
                <el-input v-model="userInfo.department"></el-input>
            </el-form-item>
            <el-form-item label="职位" label-width="80px" prop="postition">
                <el-input v-model="userInfo.postition"></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
            <el-button @click="closeDialog">取 消</el-button>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
    </el-dialog>
  </div>
</template>

<script>
import infoList from '@/mixins/infoList'
import {
    detail,
    addUser,
    paging,
    deleteSoft,
    updatepassword,
    updateRole,
    deleteUser,
    editUser
} from '@/api/user/user'

import {
    getAllRole
} from '@/api/role/role'
export default {
    name: "User",
    mixins: [infoList],
    data(){
        var validatePhone = (rule, value, callback) => {
            const v = value + "a"
            if (v.length !== 12) {
                callback(new Error('电话长度为11位'));
            } else {
                callback();
            }
        };
        return {
            listApi: paging,
            isEdit: [],
            roleOptions: [],
            roleProps: { multiple: true },
            softDeleteValue: "todelete",
            dialogVisible: false,
            dialogFlag: "",
            dialogTitle: "",
            userInfo:{
                id: "",
                name: "",
                nick_name:"",
                type: "",
                password: "",
                confirmPassword: "",
                email: "",
                phone: "",
                department: "",
                postition: "",
                roleList:[]
            },
            rulesadduser: {
                name:[
                    { required: true, message: "输入用户名", trigger: "blur" },
                    { min: 5, message: "大于5个字符", trigger: "blur" }
                ],
                nick_name:[
                    { required: true, message: "输入用户昵称", trigger: "blur" },
                    { min: 2, message: "大于5个字符", trigger: "blur" }
                ],
                type: [
                    { required: true, message: "选择用户类型", trigger: "change" },
                ],
                password: [
                    { required: true, message: "输入用户密码", trigger: "blur" },
                    { min: 5, message: "大于5个字符", trigger: "blur" }
                ],
                confirmPassword: [
                    { required: true, message: "确认用户密码", trigger: "blur" },
                    { min: 5, message: "大于5个字符", trigger: "blur" },
                    { 
                        validator: (rule, value, callback) => {
                            if (value !== this.userInfo.password) {
                                callback(new Error("两次密码不一致"))
                            }else {
                                callback()
                            }
                        },
                        trigger: "blur" 
                    }
                ],
                email: [
                    { required: true, message: "请输入邮箱地址", trigger: "blur"},
                    { type: "email", message: "请输入正确的邮箱地址", trigger:["blur","change"]}
                ],
                phone: [
                    { required: true, message: "输入电话", trigger: "blur" },
                    { type: "number", message: "请输入正确的电话", trigger: "blur"},
                    { validator: validatePhone, trigger: "blur"}
                ]
            }
        }
    },
    async created(){
        this.getTableData()
        const res = await getAllRole()
        this.setOptions(res.list)
    },
    methods:{
        parseTimeStamp(timestamp) {
            var time = new Date(timestamp*1000)
            var y = time.getFullYear()
            var m = time.getMonth()+1
            var d = time.getDate()
            var h = time.getHours()
            var mm = time.getMinutes()
            var s = time.getSeconds()
            return y+'-'+this.add0(m)+'-'+this.add0(d)+' '+this.add0(h)+':'+this.add0(mm)+':'+this.add0(s);
        },
        add0(m){
            return m<10?'0'+m:m
        },
        setRoleOptions(AuthorityData, optionsData) {
            AuthorityData &&
                AuthorityData.map(item => {
                    if (item.children && item.children.length) {
                        const option = {
                            roleid: item.id,
                            roleName: item.role + "-" + item.name,
                            children: []
                        }
                        this.setAuthorityOptions(item.children, option.children)
                        optionsData.push(option)
                    } else {
                        const option = {
                            roleid: item.id,
                            roleName: item.role + "-" + item.name,
                        }
                        optionsData.push(option)
                    }
                })
        },
        setOptions(roleData){
            this.roleOptions = []
            this.setRoleOptions(roleData, this.roleOptions)
        },
        initForm(){
            this.$refs.userInfoFrom.resetFields()
            this.dialogFlag = ""
            this.dialogTitle = ""
            this.userInfo = {
                id: "",
                name: "",
                nick_name:"",
                type: "",
                password: "",
                confirmPassword: "",
                email: "",
                phone: "",
                department: "",
                postition: "",
                roleList:[]
            }
        },
        handleClose(){
            this.initForm()
        },
        closeDialog(){
            this.initForm()
            this.dialogVisible = false
        },
        rowClick(row, column, cell, event) {
            //判断是否是需要编辑的列 再改变对应的值
            if(column.property == 'password') {
                this.$set(this.isEdit,row.id,true)
                this.$nextTick(function () {
                    this.$refs.inputpassword.focus()
                })
                return
            }
        },
        async changePassword(row){
            console.log(row)
            const res = await updatepassword({id: row.id, password: row.password})
            if (res.code === 200) {
                this.$set(this.isEdit,row.id,false)
                this.getTableData()
            }else {
                this.$message({
                    type: "error",
                    message: res.msg
                })
                this.$set(this.isEdit,row.id,false)
            }
        },
        async changeRole(row) {
            const res = await updateRole({id: row.id, roleList: row.roleList})
            if (res.code !== 200) {
                this.$message({
                    type: "error",
                    message: res.msg
                })
            }
            this.getTableData()
        },
        //软删除用户
        async softDeleteUser(e,row,index){
            const res = await deleteSoft({id: row.id, state: e})
            if (res.code === 200) {
                this.$message({
                    type: "success",
                    message: res.msg
                })
            }else {
                this.$message({
                    type: "error",
                    message: res.msg
                })
            }
            this.getTableData()
        },
        //删除用户
        async deleteUser(row) {
            const res = await deleteUser({id: row.id})
            if (res.code === 200) {
                this.getTableData()
                row.visible = false
            }else {
                this.$message({
                    type:"error",
                    message: res.msg
                })
            }
        },
        //添加用户
        addUser(){
            this.dialogFlag = "adduser"
            this.dialogTitle = "新增用户"
            this.userInfo.type = 0
            this.dialogVisible = true
        },
        async editUser(row) {
            const res = await detail(row.id)
            if (res.code === 200) {
                // this.userInfo = JSON.parse(JSON.stringify(row))
                this.userInfo = JSON.parse(JSON.stringify(res.userInfo))
                this.userInfo.phone = parseInt(this.userInfo.phone)
                this.dialogFlag = "edituser"
                this.dialogTitle = "编辑用户"
                this.dialogVisible = true
            }

        },
        searchUser() {
            console.log("searchInfo", this.searchUserInfo)
            this.getTableData()
        },
        async enterDialog(){
            this.$refs.userInfoFrom.validate(async valid => {
                if (valid) {
                    if (this.dialogFlag === "adduser") {
                        const res = await addUser(this.userInfo)      
                        if (res.code === 200) {
                            this.$message({type:"success", message: res.msg})
                            this.getTableData()
                            this.closeDialog()
                        }
                    }else if (this.dialogFlag === "edituser"){
                        const res = await editUser(this.userInfo) 
                        if (res.code === 200) {
                            this.$message({type:"success", message: res.msg})
                            this.closeDialog()
                            this.getTableData()
                        }
                    }
                }
            })
        }
    }
}
</script>

<style lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
</style>