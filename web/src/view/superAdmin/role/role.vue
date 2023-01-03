<template>
  <div>
    <div>
        <el-button size="mini" type="primary" icon="el-icon-plus" @click="addRole">新增角色</el-button>
        <el-button size="mini" type="primary" icon="el-icon-plus" :loading="casbinbtnloading" @click="refreshPermission">Casbin权限刷新</el-button>
    </div>
    <el-table
        :data="tableData"
        border
        stripe
        style="width: 100%"
    >
        <el-table-column label="id" prop="id"></el-table-column>
        <el-table-column label="角色" prop="role"></el-table-column>
        <el-table-column label="名字" prop="name"></el-table-column>
        <el-table-column label="创建人" prop="createBy"></el-table-column>
        <el-table-column label="修改人" prop="updateBy"></el-table-column>
        <el-table-column label="状态" prop="delete_time">
            <template slot-scope="scope">
                <el-tooltip :content="scope.row.delete_time<=0?'active':parseTimeStamp(scope.row.delete_time)" placement="top">
                    <el-switch
                        v-model="scope.row.state"
                        active-color="#13ce66"
                        inactive-color="#ff4949"
                        @change="softDeleteRole($event,scope.row)"
                        active-value="deleted"
                        inactive-value="resume"
                    >
                    </el-switch>
                </el-tooltip>
            </template>
        </el-table-column>
        <el-table-column label="操作" fixed="right"  width="300">
            <template slot-scope="scope">
                <el-button type="text" icon="el-icon-setting" size="mini" @click="opendrawer(scope.row)">设置权限</el-button>
                <el-button type="text" icon="el-icon-edit" size="mini" @click="editRole(scope.row)">编辑</el-button>
                <el-button type="text" icon="el-icon-delete" size="mini" @click="deleteRole(scope.row)">删除</el-button>
            </template>
        </el-table-column>
    </el-table>

    <el-dialog :before-close="handleClose" :title="dialogTitle" :visible.sync="dialogVisible">
        <el-form ref="roleForm" :model="form" :rules="rules" label-width="80px">
            <el-form-item label="id" prop="id" v-show="dialogType == 'edit'">
                <el-input v-model="form.id" type="text" disabled></el-input>
            </el-form-item>
            <el-form-item label="角色" prop="role">
                <el-input v-model="form.role" type="text"></el-input>
            </el-form-item>
            <el-form-item label="角色名字" prop="name">
                <el-input v-model="form.name" type="text"></el-input>
            </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>

    <el-drawer v-if="drawerVisible" :visible.sync="drawerVisible" :with-header="false" size="40%" title="角色配置">
        <el-tabs class="role-box" type="border-card">
            <el-tab-pane label="菜单">
                <menu-permission ref="menus" :row="activeRow"></menu-permission>
            </el-tab-pane>
            <el-tab-pane label="按钮">
                <api-permission ref="apis" :row="activeRow"></api-permission>
            </el-tab-pane>
        </el-tabs>
    </el-drawer>
  </div>
</template>

<script>
import { 
    getAllRole,
    createRole,
    refreshPermission,
    deleteSoftRole,
    detailRole,
    updateRole,
    deleteRole
} from '@/api/role/role'
import infoList from '@/mixins/infoList'
import menuPermission from '@/view/superAdmin/role/components/menupermission.vue'
import apiPermission from '@/view/superAdmin/role/components/apipermission.vue'
export default {
    name: "RoleManager",
    mixins: [infoList],
    components: {
        menuPermission,
        apiPermission
    },
    data(){
        return {
            listApi: getAllRole,
            casbinbtnloading: false,
            dialogTitle: "",
            dialogVisible: false,
            dialogType: "",
            drawerVisible: false,
            activeRow: {},
            form: {
                id: 0,
                role: "",
                name: ""
            },
            rules: {
                role: [
                    { required: true, message: '请输入角色', trigger: 'blur' },
                ],
                name: [
                    { required: true, message: '请输入角色名', trigger: 'blur' },
                ]
            }
        }
    },
    async created(){
        this.getTableData()
    },
    methods: {
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
        async refreshPermission(){
            this.casbinbtnloading = true
            const res = await refreshPermission()
            if (res.code === 200) {
                this.$message({
                    type:"success",
                    message: res.msg
                })
                this.casbinbtnloading = false
            }
        },
        //软删除用户
        async softDeleteRole(e,row){
            const res = await deleteSoftRole({id: row.id, state: e})
            if (res.code === 200) {
                this.$message({
                    type: "success",
                    message: res.msg
                })
            }
            this.getTableData()
        },
        deleteRole(row){
            this.$confirm('此操作将永久删除该角色相关信息, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            })
            .then(async () => {
                console.log("===>", row)
                const res = await deleteRole(row)
                if (res.code === 200){
                    this.$message({
                        type:"success",
                        message:res.msg
                    })

                    this.getTableData()
                }
            })
            .catch(() => {
                this.$message({
                    type:"info",
                    message:"已取消删除"
                })
            })
        },
        async editRole(row) {
            this.dialogTitle = "编辑角色"
            this.dialogType = "edit"
            const res = await detailRole(row)
            if (res.code === 200) {
                this.form = JSON.parse(JSON.stringify(res.detail))
                this.dialogVisible = true
            }

        },
        addRole() {
            this.dialogTitle = "新增角色"
            this.dialogType = "add"
            this.dialogVisible = true
        },
        enterDialog() {
            this.$refs.roleForm.validate(async valid => {
                if (valid) {
                    switch (this.dialogType){
                        case 'add':
                            const res = await createRole(this.form)
                            if (res.code === 200) {
                                this.$message({
                                    type:"success",
                                    message: res.msg,
                                })
                                this.initForm()
                                this.dialogVisible = false
                                this.getTableData()
                            }
                            break
                        case 'edit':
                            console.log("===>",this.form)
                            const ress = await updateRole(this.form)
                            if (ress.code === 200) {
                                this.$message({
                                    type:"success",
                                    message: ress.msg,
                                })
                                this.initForm()
                                this.dialogVisible = false
                                this.getTableData()
                            }
                            break
                        default:
                            this.$message({
                                type:"waring",
                                message: "未知操作" + this.dialogType
                            })
                    }

                }
            })
        },
        closeDialog(){
            this.initForm()
            this.dialogVisible = false
        },
        handleClose() {
            this.initForm()
            this.dialogVisible = false
        },
        initForm(){
            if (this.$refs.roleForm) {
                this.$refs.roleForm.resetFields()
            }
            this.form = {
                id: "",
                role: "",
                name: ""
            }
        },
        opendrawer(row){
            this.drawerVisible = true
            this.activeRow = row
        }
    }
}
</script>

<style lang="scss" scoped>
.role-box {
  .el-tabs__content {
    height: calc(100vh - 150px);
    overflow: auto;
  }
}
</style>