<template>
  <div>
    <div>
        <el-button type="primary" icon="el-icon-plus" size="mini" style="float:right" @click="addESConn">新建连接信息</el-button>
    </div>
    <el-table
        :loading="loading"
        :data="tableData"
        border
        stripe
    >
        <el-table-column label="id" min-width="60" prop="id"></el-table-column>
        <el-table-column label="Host" min-width="280" prop="es_conn"></el-table-column>
        <el-table-column label="版本" min-width="150" prop="version"></el-table-column>
        <el-table-column label="用户名" min-width="150" prop="user"></el-table-column>
        <el-table-column label="描述" min-width="150" prop="describe"></el-table-column>
        <el-table-column label="操作" min-width="300" fixed="right">
            <template slot-scope="scope">
                <el-button size="small" type="success" icon="el-icon-connection" @click="esping(scope.row)">连接测试</el-button>
                <el-button size="mini" type="primary" icon="el-icon-edit" @click="esConnEdit(scope.row)">编辑</el-button>
                <el-button size="mini" type="danger" icon="el-icon-delete" @click="deleteESConn(scope.row)">删除</el-button>
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
    ></el-pagination>

        <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
            <el-form ref="esConnForm" :inline="false" :model="form" :rules="rules" label-width="80px">
                <el-form-item label="ES连接" prop="es_conn">
                    <el-input v-model="form.es_conn" auto-complete="off" placeholder="http://<IP>:<PORT>,http://<IP>:<PORT>"></el-input>
                </el-form-item>
                <el-form-item label="用户名" prop="user">
                    <el-input v-model="form.user" auto-complete="off"></el-input>
                </el-form-item>
                <el-form-item label="密码" prop="passWord">
                    <el-input v-model="form.passWord" auto-complete="off"></el-input>
                </el-form-item>
                <el-form-item label="描述" prop="describe">
                    <el-input v-model="form.describe" auto-complete="off"></el-input>
                </el-form-item>
                <el-form-item label="版本" prop="method">
                    <el-select v-model.number="form.version" placeholder="请选择">
                        <el-option
                            v-for="item in versionOptions"
                            :key="item.value"
                            :label="`${item.label}`"
                            :value="item.value"
                        >
                        </el-option>
                    </el-select>
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
import infoList from "@/mixins/infoList"
import {
    esConnList,
    esPing,
    deleteESConn,
    esConnDetail,
    createESConn,
    updateESConn
} from "@/api/esManager/conn/conn.js"
const versionOptions = [
  {
    value: 7,
    label: "7.XX"
  },
  {
    value: 6,
    label: "6.XX"
  }
]
export default {
    name: "ConnManager",
    mixins: [infoList],
    data(){
        return {
            listApi: esConnList,
            dialogFormVisible: false,
            dialogTitle: "",
            type:"",
            versionOptions: versionOptions,
            form: {
                es_conn: "",
                user: "",
                passWord: "",
                describe: "",
                version:7,
            },
            rules: {
                es_conn: [
                    { required: true, message:"输入ES连接", trigger: 'burl'}
                ],
                version: [
                    {type: "number", message:"必须是数字", trigger: 'blur'}
                ],
                describe: [
                    { required: true, message:"输入ES连接描述", trigger: 'burl'}
                ]
            }
        }
    },
    async created() {
        await this.getTableData()
    },
    methods: {
        async esping(data) {
            const res = await esPing(data)
            if (res.code === 200) {
                this.$message({
                    type:"success",
                    message: res.msg + "  ES版本为:" + res.data.version.number
                })
            }
        },
        async deleteESConn(data) {
            const res = await deleteESConn(data)
            if (res.code === 200) {
                this.$message({
                    type: "success",
                    message: res.msg
                })
                if (this.tableData.length === 1 && this.page > 1) {
                    this.page--
                }
                await this.getTableData()
            }
        },
        //编辑创建
        addESConn() {
            this.openDialog("addESConn")
        },
        async esConnEdit(data) {
            const res = await esConnDetail(data)
            if (res.code === 200) {
                this.form = res.detail
                this.openDialog("editESConn")
            }
        },
        openDialog(type) {
            switch (type) {
                case 'addESConn':
                    this.dialogTitle = "新增ES连接"
                    break
                case 'editESConn':
                    this.dialogTitle = "编辑ES连接"
                    break
                default:
                    break
            }
            this.type = type
            this.dialogFormVisible = true
        },
        closeDialog() {
            this.initForm()
            this.dialogFormVisible = false
        },
        initForm() {
            this.$refs.esConnForm.resetFields()
            this.form = {
                es_conn: "",
                user: "",
                passWord: "",
                describe: "",
                version: 7,
            }
        },
        enterDialog() {
            this.$refs.esConnForm.validate(async valid => {
                if (valid) {
                    switch (this.type) {
                        case "addESConn":
                            {
                                const res = await createESConn(this.form)
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
                        case "editESConn": 
                            {
                                const res = await updateESConn(this.form)
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

<style>

</style>