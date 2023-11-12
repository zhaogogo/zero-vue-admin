<template>
    <div>
        <div style="margin-bottom: 10px;">
            <el-button type="info" size="mini" @click="getSlienceJson">获取静默视图json</el-button>
            <el-button style="float: right;" type="primary" size="mini" icon="el-icon-plus" @click="addHost">添加主机</el-button>
        </div>
        <div style="margin-bottom: 10px;">
            <el-button type="primary" size="mini" @click="mulSlienceHostsHandler">批量静默主机</el-button>
            <el-button type="warning" size="mini" @click="mulDelSlienceHostsHandler">批量删除静默主机</el-button>
        </div>
        <el-table
            :loading="loading"
            :data="tableData"
            :border="true"
            :stripe="true"
            @selection-change="handleSlienceHostsChange"
        >
            <el-table-column
                type="selection"
                width="55">
            </el-table-column>
            <el-table-column label="id" prop="id" min-width="60"></el-table-column>
            <el-table-column label="主机" min-width="120">
                <template slot-scope="scope">
                    <el-button type="text" @click="openSlienceDialog(scope.row.host)">{{ scope.row.host }}</el-button>
                </template>
            </el-table-column>
            <el-table-column label="机房">
                <template slot-scope="scope">
                    <el-tag
                        :key="scope.row.to"
                        :type="scope.row.to|tagTypeFilter"
                        effect="dark"
                        size="mini"
                    >
                        {{ scope.row.to|hostLocationFilter }}
                    </el-tag>
                </template>
            </el-table-column>
            <el-table-column label="标签">
                <template slot-scope="scope">
                    <el-tag  size="mini" v-for="item in scope.row.tags" :key="item.id">{{ item.key }}={{ item.value }}</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="操作" width="350">
                <template slot-scope="scope">
                    <el-button size="mini" type="primary" @click="slienceHostsHandler(scope.row.host)">静默</el-button>
                    <el-button size="mini" type="warning" @click="delSlienceHostsHandler(scope.row.host)">删除静默</el-button>
                    <el-button size="mini" type="primary" icon="el-icon-edit" @click="editHost(scope.row)">编辑</el-button>
                    <el-button size="mini" type="primary" icon="el-icon-document-copy" @click="copyHost(scope.row)">复制</el-button>
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
<!-- 添加主机dialog -->
        <el-dialog :title="hostTitle" :visible.sync="hostDialogVisible" :before-close="closeHostDialogDefault">
            <el-form ref="hostForm" :model="hostFormData" :rules="hostFormRules">
                <el-form-item v-show="hostTitle!=='添加主机'?true:false" label="ID" prop="id" label-width="80px">
                    <el-input v-model.number="hostFormData.id" :disabled="true"></el-input>
                </el-form-item>
                <el-form-item label="主机" prop="host" label-width="80px">
                    <el-input v-model="hostFormData.host"></el-input>
                </el-form-item>
                <el-form-item label="机房" prop="to" label-width="80px">
                    <el-radio-group v-model="hostFormData.to" size="mini">
                        <el-radio-button :label="1">兆维</el-radio-button>
                        <el-radio-button :label="2">亦庄</el-radio-button>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="标签" label-width="80px">
                    <el-table :data="hostFormData.tags" border>
                        <el-table-column label="key" prop="key">
                            <template slot-scope="scope">
                                <el-input v-model="scope.row.key" placeholder="key"></el-input>
                            </template>
                        </el-table-column>
                        <el-table-column label="value" prop="value">
                            <template slot-scope="scope">
                                <el-input v-model="scope.row.value" ></el-input>
                            </template>
                        </el-table-column>
                        <el-table-column align="center" :render-header="renderHostHeader" width="80px">
                            <template slot-scope="scope">
                                <el-button type="danger" icon="el-icon-delete" size="mini" circle @click="delHostTag(scope.$index)"></el-button>
                            </template>
                        </el-table-column>
                    </el-table>
                </el-form-item>
                <el-form-item size="large">
                    <el-button type="primary" @click="enterHostDialog">提交</el-button>
                    <el-button @click="closeHostDialog">取消</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
<!-- 自定义静默规则dialog -->
        <el-drawer
            :title="slienceTitle"
            :before-close="closeSlienceDialog"
            :visible.sync="slienceDialogVisible"
            direction="rtl"
            custom-class="demo-drawer"
            ref="hostSlienceRule"
            size="60%"
            >
            <div class="demo-drawer__content">
                <el-form ref="slienceForm" label-width="100px">
                    <el-form-item 
                        v-for="item,hostSlienceIndex in hostSlienceRules.sliences"
                        :key="hostSlienceIndex"
                        :prop="item.slience_name"
                        label="静默描述"
                    >
                        <el-col>
                            <el-input v-model="item.slience_name" placeholder="静默描述" style="width: 80%;"></el-input>
                            <el-button @click.prevent="removeSlience(item)">删除</el-button>
                        </el-col>
                        <el-col>
                        <el-form-item label="是否为默认">
                            <el-checkbox v-model="item.default" value="false"></el-checkbox>
                            
                            <el-radio-group v-model.number="item.to" size="mini" style="margin-left: 30px;">
                                <el-radio-button :label="1">兆维</el-radio-button>
                                <el-radio-button :label="2">亦庄</el-radio-button>
                            </el-radio-group>
                            
                        </el-form-item>
                        </el-col>
                        <el-tag
                            :key="index"
                            v-for="matcher,index in item.matchers"
                            closable
                            :disable-transitions="false"
                            @close="handleCloseMatcher(item.matchers,index)">
                            {{matcher | matcherFilter }}
                        </el-tag>
                        <el-input
                            :ref="'saveTagInput'+item.slience_name"
                            class="input-new-tag"
                            v-show="inputVisible[item.slience_name]"
                            v-model="inputValue"
                            size="small"
                            @keyup.enter.native="handleInputConfirm(hostSlienceIndex)"
                            @blur="handleInputConfirm(hostSlienceIndex)"
                            >
                        </el-input>
                        <el-button class="button-new-tag" size="small" @click="showInput(item.slience_name)">+ New Tag</el-button>
                        
                    </el-form-item>
                </el-form>
                <el-button style="margin-left: 5%;" @click="addSlience(hostSlienceRules.id)">新增静默规则</el-button>
                <div class="demo-drawer__footer">
                    <!-- <el-button @click="cancelForm">取 消</el-button> -->
                    <el-button type="primary" @click="$refs.hostSlienceRule.closeDrawer()" :loading="slienceRuleLoading">{{ slienceRuleLoading ? '提交中 ...' : '确 定' }}</el-button>
                </div>
            </div>
        </el-drawer>

<!-- 自定义静默规则json dialog -->
        <el-dialog
            title="自定义静默规则"
            :visible.sync="allSlienceJsonVisible"
            width="90%"
            :before-close="allSlienceJsonClose"
        >
            <vue-json-editor
                :show-btns="false"
                :expandedOnStart="true"
                v-model="allSlienceJsonData"
            >
            </vue-json-editor>
        </el-dialog>
    </div>
</template>

<script>
import vueJsonEditor from 'vue-json-editor'
import infoList from '@/mixins/infoList'
import {
    hosts,
    createHost,
    updateHost,
    getAllSlience,
    handlerHostsSlience,
    hostSlienceRule,
    putHostSlienceRule
} from '@/api/monitoring/hosts/hosts'
import { matcherExpParse } from "@/utils/slienceParse"

const hostLocationOptions = [
    {
        value: 1,
        label: "兆维",
        type: "success"
    },
    {
        value: 2,
        label: "亦庄",
        type: ""
    },
    {
        value: 3,
        label: "通用",
        type: "success"
    }
]

export default {
    name: "MonitorHosts",
    components: { vueJsonEditor },
    mixins: [infoList],
    filters: {
        hostLocationFilter(v) {
            const target = hostLocationOptions.filter(item => item.value === v)[0]
            return target && `${target.label}`
        },
        tagTypeFilter(v) {
            const target = hostLocationOptions.filter(item => item.value === v)[0]
            return target && `${target.type}`
        },
        matcherFilter(v) {
            if (v.is_equal == true ) {
                if (v.is_regex == true) {
                    return v.name + "=~" + "\""+ v.value +"\""
                }else {
                    return v.name + "=" + "\""+ v.value +"\""
                }
            }else {
                if (v.is_regex == true) {
                    return v.name + "!~" + "\""+ v.value +"\""
                }else {
                    return v.name + "!=" + "\""+ v.value +"\""
                }
            }
        }
    },
    data() {
        return {
            listApi: hosts,
            hostLocationOptions: hostLocationOptions,
            
            allSlienceJsonVisible: false,
            slienceDialogVisible: false,
            hostDialogVisible: false,
            inputVisible: {},

            slienceTitle: "",
            hostTitle: "",
            
            slienceRuleLoading: false,
            hostSlienceRules: {},
            inputValue:"",
            mulSlienceHosts: [],

            allSlienceJsonData: {},
            hostFormData:{
                id:0,
                host: "",
                to:0,
                tags: []
            },
            hostFormRules: {
                host: [
                    {required: true, message:"主机名必须添加", target: "blur"}
                ],
                to: [
                    {required: true, type: "number",message: "必选项", target: "change"},
                ]
            }
        }
    },
    async created() {
        await this.getTableData()
    },
    methods: {
        initHostForm() {
            this.$refs.hostForm.resetFields()
            this.hostFormData = {
                id:0,
                host: "",
                to:0,
                tags: []
            }
        },
        closeHostDialogDefault(done) {
            this.initHostForm()
            done()
        },
        closeHostDialog(){
            this.initHostForm()
            this.hostDialogVisible = false
        },
        addHost() {
            this.hostTitle = "添加主机"
            this.hostDialogVisible = true
        },
        editHost(data){
            this.hostTitle = "编辑主机"
            this.hostDialogVisible = true
            this.hostFormData =  data
        },
        copyHost(data) {
            this.hostTitle = "添加主机"
            this.hostFormData = data
            this.hostFormData.id = 0
            let i
            for (i=0;i<this.hostFormData.tags.length;i++){
                this.hostFormData.tags[i].id = 0
            }
            this.hostDialogVisible = true
        },
        renderHostHeader(h, a) {
            let that = this
            // 逻辑是 h() 括号里包裹标签 第一个参数是标签名 第二个是属性  第三个是标签内容  如果是多个标签需要包裹数组
            // <el-button icon="el-icon-plus" type="success" size="mini" circle onClick={ this.addHostTag }></el-button>
            return h("div",[
                h("el-button",{
                    props: {
                        type: "success",
                        size: "mini",
                        circle: true,
                        icon:"el-icon-plus"
                    },
                    on: {
                        click: function() {
                            that.addHostTag()
                        }
                    }
                })
            ])
        },
        addHostTag() {
            if (!this.hostFormData.tags) {
                this.$set(this.hostFormData,"tags",[])
            }
            this.hostFormData.tags.push({key:"",value:""})
        },
        delHostTag(index) {
            this.hostFormData.tags.splice(index,1)
        },
        enterHostDialog(){
            this.$refs.hostForm.validate(async(valid)=>{
                if (valid) {
                    let res 
                    switch(this.hostTitle) {
                        case "添加主机":
                            res = await createHost(this.hostFormData)
                            break
                        case "编辑主机":
                            res = await updateHost(this.hostFormData)
                    }
                    if (res.code === 200) {
                        this.$message({
                            type: 'success',
                            message: res.msg
                        })
                        this.getTableData()
                        this.closeHostDialog()
                    }
                }
            })
        },
        slienceHostsHandler(host) {
            let hosts = []
            hosts.push(host)

            this.$prompt('请输入静默时间', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                inputPattern: /[0-9]+[smhd]{1}/,
                inputErrorMessage: '时间格式不正确'})
                .then(async ({ value }) => {
                    const res = await handlerHostsSlience({"hosts":hosts,"duration": value, "op_type": "active"}) 
                    if (res.code === 200) {
                        this.$message({
                            type: "success",
                            message: res.msg,
                        })
                    }
                })
                .catch(_ => {
                    this.$message({
                        type: 'info',
                        message: '取消输入'
                    });
                });
        },
        mulSlienceHostsHandler() {
            if (this.mulSlienceHosts.length <= 0) {
                this.$message({
                    type: "warning",
                    message: "必须选择主机"
                })
                return
            }
            this.$prompt('请输入静默时间', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                inputPattern: /[0-9]+[smhd]{1}/,
                inputErrorMessage: '时间格式不正确'})
                .then(async ({ value }) => {
                    const res = await handlerHostsSlience({"hosts":this.mulSlienceHosts,"duration": value, "op_type": "active"}) 
                    if (res.code === 200) {
                        this.$message({
                            type: "success",
                            message: res.msg,
                        })
                    }
                })
                .catch(_ => {
                    this.$message({
                        type: 'info',
                        message: '取消输入'
                    });
                });
        },
        delSlienceHostsHandler(host) {
            let hosts = []
            hosts.push(host)

            this.$confirm('是否取消自定义静默', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            })
                .then(async _ => {
                    const res = await handlerHostsSlience({"hosts":hosts, "op_type": "expired"}) 
                    if (res.code === 200) {
                        this.$message({
                            type: "success",
                            message: res.msg,
                        })
                    }
                })
                .catch(_ => {
                    this.$message({
                        type: 'info',
                        message: '取消输入'
                    });
                });
        },
        mulDelSlienceHostsHandler() {
            if (this.mulSlienceHosts.length <= 0) {
                this.$message({
                    type: "warning",
                    message: "必须选择主机"
                })
                return
            }
            this.$confirm('是否取消自定义静默', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            })
                .then(async _ => {
                    const res = await handlerHostsSlience({"hosts":this.mulSlienceHosts, "op_type": "expired"}) 
                    if (res.code === 200) {
                        this.$message({
                            type: "success",
                            message: res.msg,
                        })
                    }
                })
                .catch(_ => {
                    this.$message({
                        type: 'info',
                        message: '取消输入'
                    });
                });
        },
        async handleSlienceHostsChange(row) {
            let host = []
            var i 
            for (i=0;i<row.length;i++) {
                host.push(row[i].host)
            }
            this.mulSlienceHosts = host
        },
        async getSlienceJson() {
            this.allSlienceJsonVisible = true
            const res = await getAllSlience()
            this.allSlienceJsonData = res.data
        },
        allSlienceJsonClose() {
            this.allSlienceJsonData = {}
            this.allSlienceJsonVisible = false
        },
        async openSlienceDialog(host) {
            this.slienceTitle = host + " 静默规则"
            const res = await hostSlienceRule(host)
            this.hostSlienceRules = res.hostSliences
            this.slienceDialogVisible = true
        },
        async closeSlienceDialog() {
            this.slienceDialogVisible = false
            this.slienceRuleLoading = true
            await putHostSlienceRule(this.hostSlienceRules)

            this.slienceRuleLoading = false
            this.hostSlienceRules = {}
            
        },
        handleCloseMatcher(matchers, index) {
            matchers.splice(index,1)
            // var f = this.hostSlienceRules.sliences[hostSlienceIndex].matchers.filter(v => v.id !== matcher.id)
            // this.hostSlienceRules.sliences[hostSlienceIndex].matchers = f 
        },
        removeSlience(hostSlience) {
            this.hostSlienceRules.sliences = this.hostSlienceRules.sliences.filter(v => v.id !== hostSlience.id)
        },
        showInput(name) {
            this.$set(this.inputVisible, name, true)
            this.$nextTick(_ => {
                this.$refs["saveTagInput"+name][0].$refs.input.focus();
            });
        },
        handleInputConfirm(hostSlienceIndex) {
            let inputValue = this.inputValue;
            if (inputValue) {
                let hostID = this.hostSlienceRules.sliences[hostSlienceIndex].host_id
                let slienceNameID 
                if (this.hostSlienceRules.sliences[hostSlienceIndex].hasOwnProperty("id")) {
                    slienceNameID = this.hostSlienceRules.sliences[hostSlienceIndex].id
                }else {
                    slienceNameID = 0
                }
                this.hostSlienceRules.sliences[hostSlienceIndex].matchers.push(matcherExpParse(inputValue,hostID,slienceNameID));
            }
            this.inputVisible = {};
            this.inputValue = '';        
        },
        addSlience(hostID){
            this.hostSlienceRules.sliences.push({
                "id": 0,
                "host_id": hostID,
                "slience_name": "",
                "default": false,
                "matchers": []
            });
        }
    }
}
</script>

<style>
</style>