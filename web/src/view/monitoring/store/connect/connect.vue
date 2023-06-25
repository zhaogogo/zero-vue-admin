<template>
    <div>
        <div>
            <el-button type="primary" icon="el-icon-plus" size="mini" style="float:right;margin-bottom:5px" @click="addConn">新建连接</el-button>
        </div>
        <el-table
            :loading="loading"
            :data="tableData"
            :border="true"
            :stripe="true"
        >
            <el-table-column label="id" min-width="60" prop="id"></el-table-column>
            <el-table-column label="type" min-width="60" prop="type"></el-table-column>
            <el-table-column label="env" min-width="60">
                <template slot-scope="scope">
                    <el-tag
                        :key="scope.row.hostVisibleFilter"
                        :type="scope.row.env|tagTypeFilter"
                        effect="dark"
                        size="mini"
                    >
                        {{ scope.row.env|hostVisibleFilter }}
                    </el-tag>
                </template>
            </el-table-column>
            <el-table-column label="host" min-width="60" prop="host"></el-table-column>
            <el-table-column label="access_key" min-width="60" prop="accessKey"></el-table-column>
            <el-table-column label="secret_key" min-width="60" prop="secretKey"></el-table-column>
            <el-table-column label="操作">
                <template slot-scope="scope">
                    <el-button size="small" type="primary" icon="el-icon-edit" @click="editConn(scope.row)">编辑</el-button>
                    <el-button size="small" type="danger" icon="el-icon-delete" @click="deleteConn(scope.row)">删除</el-button>
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
            <el-form ref="dialogForm" :inline="false" :model="form" :rules="rules" label-width="80px">
                <el-form-item label="连接" prop="host">
                    <el-input v-model="form.host" auto-complete="off" placeholder="连接IP地址"></el-input>
                </el-form-item>
                <el-form-item label="accessKey" prop="access_key">
                    <el-input v-model="form.accessKey" auto-complete="off"></el-input>
                </el-form-item>
                <el-form-item label="secretKey" prop="secret_key">
                    <el-input v-model="form.secretKey" auto-complete="off"></el-input>
                </el-form-item>
                <el-form-item label="存储类型">
                    <el-select v-model="form.type" placeholder="请选择">
                        <el-option
                            v-for="item in dialogFormOption.storeType"
                            :key="item.value"
                            :label="`${item.label}`"
                            :value="item.value"
                        >
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="环境">
                    <el-select v-model="form.env" placeholder="请选择">
                        <el-option
                            v-for="item in dialogFormOption.env"
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
import infoList from '@/mixins/infoList'
import { 
    connectList,
    connectCreate,
    connectDel,
    connectDetail,
    connectUpdate
} from '@/api/monitoring/store/connectManager'
const envOptions = [
  {
    value: "zw",
    label: "兆维"
  },
  {
    value: "yz",
    label: "亦庄"
  }
]
const storeOptions = [
  {
    value: "s3",
    label: "s3"
  }
]

const hostVisibleOptions = [
  {
    value: 'zw',
    label: '兆维',
    type: 'success'
  },
  {
    value: 'yz',
    label: '亦庄',
    type: 'warning'
  }
]

export default {
    name:"monitoringStoreS3",
    mixins: [infoList],
    data() {
        return {
            listApi: connectList,
            dialogFormVisible: false,
            dialogAction: "",
            dialogTitle:"",
            dialogFormOption: {
                env: envOptions,
                storeType: storeOptions
            },
            form: {
                host:"",
                env: "",
                type: "",
                accessKey: "",
                secretKey:"",
            },
            rules:{
                host :[
                    { required: true, message:"输入连接", trigger: 'burl'}
                ],
                accessKey :[
                    { required: true, message:"输入连接", trigger: 'burl'}
                ],
                secretKey :[
                    { required: true, message:"输入连接", trigger: 'burl'}
                ],
                env :[
                    { required: true, message:"输入连接", trigger: 'burl'}
                ],
                type :[
                    { required: true, message:"输入连接", trigger: 'burl'}
                ]
            }
        }
    },
    filters: {
        hostVisibleFilter(v) {
            console.log(v)
            const target = hostVisibleOptions.filter(item => item.value === v)[0]
            return target && `${target.label}`
        },
        tagTypeFilter(v) {
            const target = hostVisibleOptions.filter(item => item.value === v)[0]
            console.log(target)
            return target && `${target.type}`
        }
    },
    async created() {
        await this.getTableData()
    },
    methods:{
        addConn(){
            this.openDialog("add")
        },
        async deleteConn(data) {
            const res = await connectDel(data)
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
        async editConn(data) {
            const res = await connectDetail(data)
            if (res.code === 200) {
                this.form = res.detail
                this.openDialog("edit")   
            }
        },
        openDialog(dialogAction) {
            switch (dialogAction){
                case "add":
                    this.dialogTitle = "新建连接"
                    break
                case "edit":
                    this.dialogTitle = "编辑连接"
                    break
            }
            this.dialogAction = dialogAction
            this.dialogFormVisible = true
        },
        enterDialog() {
            this.$refs.dialogForm.validate(async valid => {
                if (valid) {
                    switch(this.dialogAction) {
                        case "add":
                            {
                                const res = await connectCreate(this.form)
                                if (res.code === 200) {
                                    this.$message({
                                        type: 'success',
                                        message: res.msg,
                                        showClose: true
                                    })
                                    this.getTableData()
                                }
                            }
                            break
                        case "edit": 
                            {
                                const res = await connectUpdate(this.form)
                                if (res.code === 200) {
                                    this.$message({
                                        type: 'success',
                                        message: res.msg,
                                        showClose: true
                                    })
                                    this.getTableData()
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
                    this.closeDialog()
                }
            })
        },
        //dialog
        closeDialog() {
            this.initForm()
            this.dialogFormVisible = false
        },
        initForm() {
            this.$refs.dialogForm.resetFields()
            this.form = {
                host:"",
                env: "",
                type: "",
                accessKey: "",
                secretKey:"",
            }
        }
    }
}
</script>

<style>

</style>