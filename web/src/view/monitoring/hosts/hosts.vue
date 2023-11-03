<template>
    <div>
        <el-table
            loading="loading"
            :data="tableData"
            :border="true"
            :stripe="true"
        >
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
                        <!-- <el-form-item label="静默机房">
                            
                        </el-form-item> -->
                        </el-col>
                        <el-tag
                            :key="matcher.id"
                            v-for="matcher in item.matchers"
                            closable
                            :disable-transitions="false"
                            @close="handleCloseMatcher(matcher,hostSlienceIndex)">
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
    </div>
</template>

<script>
import infoList from '@/mixins/infoList'
import {
    hosts,
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
            slienceRuleLoading: false,
            slienceDialogVisible: false,
            slienceTitle: "",
            hostSlienceRules: {},
            inputVisible: {},
            inputValue:"",
            form: {
                name: '',
                region: '',
                date1: '',
                date2: '',
                delivery: false,
                type: [],
                resource: '',
                desc: ''
            }
        }
    },
    async created() {
        await this.getTableData()
    },
    methods: {
        async openSlienceDialog(host) {
            this.slienceTitle = host + " 静默规则"
            const res = await hostSlienceRule(host)
            this.hostSlienceRules = res.hostSliences
            this.slienceDialogVisible = true
        },
        async closeSlienceDialog() {
            this.slienceDialogVisible = false
            this.slienceRuleLoading = true
            console.log("####",this.hostSlienceRules)
            await putHostSlienceRule(this.hostSlienceRules)

            this.slienceRuleLoading = false
            this.hostSlienceRules = {}
            
        },
        handleCloseMatcher(matcher, hostSlienceIndex) {
            var f = this.hostSlienceRules.sliences[hostSlienceIndex].matchers.filter(v => v.id !== matcher.id)
            this.hostSlienceRules.sliences[hostSlienceIndex].matchers = f 
        },
        removeSlience(hostSlience) {
            console.log(hostSlience)
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
                console.log("@@@@@@",matcherExpParse(inputValue,hostID,slienceNameID),slienceNameID)
                this.hostSlienceRules.sliences[hostSlienceIndex].matchers.push(matcherExpParse(inputValue,hostID,slienceNameID));
            }
            this.inputVisible = {};
            this.inputValue = '';        
        },
        addSlience(hostID){
            // this.$set(this.hostSlienceRues.slienceRules,)
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