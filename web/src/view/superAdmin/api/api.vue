<template>
    <div>
        <div class="searchContext">
            <el-form :inline="true" :model="searchInfo">
                <el-form-item label="api">
                    <el-input v-model="searchInfo.api" placeholder="api"></el-input>
                </el-form-item>
                <el-form-item label="描述">
                    <el-input v-model="searchInfo.describe" placeholder="描述"></el-input>
                </el-form-item>
                <el-form-item label="api组">
                    <el-input v-model="searchInfo.group" placeholder="api组"></el-input>
                </el-form-item>
                <el-form-item label="请求">
                    <el-select v-model="searchInfo.method" clearable placeholder="请选择">
                        <el-option
                            v-for="item in methodOptions"
                            :key="item.value"
                            :label="`${item.label}(${item.value})`"
                            :value="item.value"
                        ></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-button size="mini" type="primary" icon="el-icon-search" @click="submitQuery">查询</el-button>
                    <el-button size="moni" icon="el-icon-refresh" @click="onResetSearch">重置</el-button>
                </el-form-item>
            </el-form>
        </div>
        <div class="dataContent">
            <div class="actionContent">
                <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog('addApi')">新增</el-button>
                <el-popover v-model="deleteVisible" placement="top" width="160">
                    <p>确定要删除吗？</p>
                    <div style="text-align: right; margin-top: 8px;">
                        <el-button size="mini" link @click="deleteVisible = false">取消</el-button>
                        <el-button size="mini" type="primary" @click="deleteMulit">确定</el-button>
                    </div>
                    <el-button slot="reference" icon="el-icon-delete" size="mini" :disabled="!apis.length" style="margin-left: 10px;">删除</el-button>
                </el-popover>
            </div>
            <el-table
                loading="loading"
                :data="tableData"
                border
                stripe
                @sort-change="sortChange"
                @selection-change="selectChange"
            >
                <el-table-column type="selection" width="55"></el-table-column>
                <el-table-column label="id" min-width="60" prop="id" sortable="custom"></el-table-column>
                <el-table-column label="api接口" min-width="280" prop="api" sortable="custom"></el-table-column>
                <el-table-column label="api分组" min-width="150" prop="group" sortable="custom"></el-table-column>
                <el-table-column label="描述" min-width="150" prop="describe" sortable="custom"></el-table-column>
                <el-table-column label="请求方法" min-width="150" prop="method" sortable="custom">
                    <template slot-scope="scope">
                        <div>
                            {{scope.row.method}}
                            <el-tag
                                :key="scope.row.methodFilter"
                                :type="scope.row.method|tagTypeFilter"
                                effect="dark"
                                size="mini"
                            >
                                {{scope.row.method|methodFilter}}
                            </el-tag>
                        </div>
                    </template>
                </el-table-column>
                <el-table-column label="操作" min-width="200" fixed="right">
                    <template slot-scope="scope">
                        <el-button size="small" type="primary" icon="el-icon-edit" @click="editApi(scope.row)">编辑</el-button>
                        <el-button size="mini" type="danger" icon="el-icon-delete" @click="deleteAPI(scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination
                :current-page="page"
                :page-size="pageSize"
                :page-sizes="[10, 30, 50, 100]"
                :style="{float:'right',padding:'20px'}"
                :total="total"
                layout="total, sizes, prev, pager, next, jumper"
                @current-change="handlePageChange"
                @size-change="handlePageSizeChange"
            />
        </div>

        <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
            <div class="warning">
                <i class="el-icon-warning-outline"><span>新增Api需要在角色管理内配置权限才可使用</span></i>
            </div>
            <el-form ref="apiForm" :inline="false" :model="form" :rules="rules" label-width="80px">
                <el-form-item label="API" prop="api">
                    <el-input v-model="form.api" auto-complete="off"></el-input>
                </el-form-item>
                <el-form-item label="描述" prop="describe">
                    <el-input v-model="form.describe" auto-complete="off"></el-input>
                </el-form-item>
                <el-form-item label="请求方法" prop="method">
                    <el-select v-model="form.method" placeholder="请选择">
                        <el-option
                            v-for="item in methodOptions"
                            :key="item.value"
                            :label="`${item.label}(${item.value})`"
                            :value="item.value"
                        >
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="分组" prop="group">
                    <el-input v-model="form.group" autocomplete="off"></el-input>
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
import {
    pagingAPI,
    deleteAPI,
    createApi,
    detailAPI,
    editApi,
    deleteMulitAPI
} from "@/api/api/api.js"
import infoList from '@/mixins/infoList'

const methodOptions = [
  {
    value: 'POST',
    label: '创建',
    type: 'success'
  },
  {
    value: 'GET',
    label: '查看',
    type: ''
  },
  {
    value: 'PUT',
    label: '更新',
    type: 'warning'
  },
  {
    value: 'DELETE',
    label: '删除',
    type: 'danger'
  }
]
export default {
    name: "APIManager",
    mixins: [infoList],
    filters: {
        methodFilter(v) {
            const target = methodOptions.filter(item => item.value === v)[0]
            return target && `${target.label}`
        },
        tagTypeFilter(v) {
            const target = methodOptions.filter(item => item.value === v)[0]
            return target && `${target.type}`
        }
    },
    data(){
        return {
            listApi: pagingAPI,
            apis: [],
            deleteVisible: false,
            dialogTitle: "",
            dialogFormVisible: false,
            methodOptions: methodOptions,
            form:{
                api: "",
                describe:"",
                method: "",
                group: ""
            },
            rules:{
                api: [
                    { required: true, message:"输入api路径", trigger: 'burl'}
                ],
                group: [
                    { required: true, message:"输入api组名", trigger: 'burl'}
                ],
                method: [
                    { required: true, message:"输入api方法", trigger: 'burl'}
                ],
                describe: [
                    { required: true, message:"输入api描述", trigger: 'burl'}
                ]
            }
        }
    },
    created(){
        this.getTableData()
    },
    methods:{
        //删除api
        deleteAPI(row){
            this.$confirm('此操作将永久删除所有角色下该api, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            })
                .then(async () => {
                    const res =  await deleteAPI(row)
                    if (res.code === 200){
                        this.$message({
                            type:"success",
                            message:res.msg
                        })
                        if (this.tableData.length === 1 && this.page > 1) {
                            this.page--
                        }
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
        async deleteMulit(){
            console.log("===>",this.apis)
            const res = await deleteMulitAPI({apis: this.apis})
            if (res.code === 200) {
                this.$message({
                    type: "success",
                    message: res.msg,
                })
                if (this.tableData.length === this.apis.length && this.page > 1) {
                    this.page--
                }
                this.deleteVisible = false
                this.getTableData()
            }
        },
        async editApi(row){
            const res = await detailAPI(row)
            if (res.code === 200) {
                this.form = res.detail
                this.openDialog("editApi")
            }
        },
        onResetSearch(){
            this.searchInfo = {}
        },
        //多选选api
        selectChange(rows) {
            this.apis = rows
        },
        //排序查询
        sortChange({ prop, order }){
            if (prop) {
                this.searchInfo.orderKey = prop
                this.searchInfo.order = order
            }
            this.getTableData()
        },
        //过滤查询
        async submitQuery(){
            this.page = 1
            const res = await this.getTableData()
            if (res.code === 200) {
                this.$message({
                    type:"success",
                    message:res.msg
                })
            }
        },
        initForm() {
            this.$refs.apiForm.resetFields()
            this.form = {
                api: "",
                describe:"",
                method: "",
                group: ""
            }
        },
        closeDialog() {
            this.initForm()
            this.dialogFormVisible = false
        },
        openDialog(type) {
            switch (type) {
                case 'addApi':
                    this.dialogTitle = "新增API"
                    break
                case 'editApi':
                    this.dialogTitle = "编辑API"
                    break
                default:
                    break
            }
            this.type = type
            this.dialogFormVisible = true
        },
        async enterDialog() {
            this.$refs.apiForm.validate(async valid => {
                if (valid) {
                    switch (this.type) {
                        case "addApi":
                            {
                                const res = await createApi(this.form)
                                if (res.code === 200) {
                                    this.$message({
                                        type: 'success',
                                        message: res.msg,
                                        showClose: true
                                    })
                                    this.getTableData()
                                    this.closeDialog()
                                }
                            }
                            break
                        case "editApi": 
                            {
                                const res = await editApi(this.form)
                                if (res.code === 200) {
                                    this.$message({
                                        type: 'success',
                                        message: res.msg,
                                        showClose: true
                                    })
                                    this.getTableData()
                                    this.closeDialog()
                                }
                            }
                            break
                        default:
                            {
                                this.$message({
                                type: 'error',
                                message: '未知操作',
                                showClose: true
                                })
                            }
                            break
                    }
                }
            })
        }
    }
}
</script>

<style lang="scss">
.warning {
    font-size: 14px;
    padding: 4px 14px;
    display: flex;
    align-items: center;
    border-radius: 3px;
    margin-bottom: 12px;
}
.actionContent {
    margin-bottom: 12px;
}
</style>