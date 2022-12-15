<template>
  <div>
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
        <el-table-column label="用户角色" prop="role">
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
        <el-table-column label="邮箱"  prop="email"></el-table-column>
        <el-table-column label="电话"  prop="phone"></el-table-column>
        <el-table-column label="部门"  prop="department"></el-table-column>
        <el-table-column label="职位"  prop="position"></el-table-column>
        <el-table-column label="创建人"  prop="createBy"></el-table-column>
        <el-table-column label="创建时间"  prop="create_time">
            <template slot-scope="scope">
                <div>{{parseTimeStamp(scope.row.create_time)}}</div>
            </template>
        </el-table-column>
        <el-table-column label="修改人"  prop="updateBy"></el-table-column>
        <el-table-column label="修改时间"  prop="update_time">
            <template slot-scope="scope">
                <div>{{parseTimeStamp(scope.row.update_time)}}</div>
            </template>
        </el-table-column>
        <el-table-column label="删除人"  prop="deleteBy"></el-table-column>
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
  </div>
</template>

<script>
import infoList from '@/mixins/infoList'
import { 
    getPagingUser,
    softDeleteUser,
    changeUserPassword,
    updateUserRole
} from '@/api/user/user'
import {
    getAllRole
} from '@/api/role/role'
export default {
    name: "User",
    mixins: [infoList],
    data(){
        return {
            isEdit: [],
            roleOptions: [],
            roleOptions: [],
            roleProps: { multiple: true },
            softDeleteValue: "todelete",
            listApi: getPagingUser
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
            console.log("---->", row.password)
            const res = await changeUserPassword({id: row.id, password: row.password})
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
            const res = await updateUserRole({userId: row.id, roleList: row.roleList})
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
            const res = await softDeleteUser({userId: row.id, state: e})
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
        }
    }
}
</script>

<style>

</style>