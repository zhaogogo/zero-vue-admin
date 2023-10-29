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
                        v-for="item,hostSlienceId in hostSlienceRules"
                        :key="item.slience_name"
                        :prop="item.slience_name"
                        label="静默描述"
                    >
                        <el-col>
                            <el-input v-model="item.slience_name" ref="test" style="width: 80%;"></el-input>
                            <el-checkbox v-model="item.default" value="false"></el-checkbox>
                            <el-button @click.prevent="removeSlience(item)">删除</el-button>
                        </el-col>
                        <el-tag
                            :key="matcher.id"
                            v-for="matcher in item.matchers"
                            closable
                            :disable-transitions="false"
                            @close="handleCloseMatcher(matcher,hostSlienceId)">
                            {{matcher | matcherFilter }}
                        </el-tag>
                        <el-input
                            :ref="'saveTagInput'+item.slience_name"
                            class="input-new-tag"
                            v-show="inputVisible[item.slience_name]"
                            v-model="inputValue"
                            size="small"
                            @keyup.enter.native="handleInputConfirm(hostSlienceId)"
                            @blur="handleInputConfirm(hostSlienceId)"
                            >
                        </el-input>
                        <el-button class="button-new-tag" size="small" @click="showInput(item.slience_name)">+ New Tag</el-button>
                        
                    </el-form-item>
                </el-form>
                <div class="demo-drawer__footer">
                    <!-- <el-button @click="cancelForm">取 消</el-button> -->
                    <!-- <el-button type="primary" @click="$refs.drawer.closeDrawer()" :loading="loading">{{ loading ? '提交中 ...' : '确 定' }}</el-button> -->
                </div>
            </div>
        </el-drawer>
    </div>
</template>

<script>
import infoList from '@/mixins/infoList'
import {
    hosts,
    hostSlienceRule
} from '@/api/monitoring/hosts/hosts'
import { del } from 'vue'

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
            console.log("---->v",v)
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
            slienceDialogVisible: false,
            slienceTitle: "",
            hostSlienceRules: [],
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
        console.log(this)
    },
    methods: {
        async openSlienceDialog(host) {
            this.slienceTitle = host + " slience"
            const res = await hostSlienceRule(host)
            this.hostSlienceRules = res.list
            this.slienceDialogVisible = true
        },
        closeSlienceDialog() {
            this.slienceDialogVisible = false
        },
        handleCloseMatcher(matcher, hostSlienceId) {
            var f = this.hostSlienceRules[hostSlienceId].matchers.filter(v => v.id !== matcher.id)
            this.hostSlienceRules[hostSlienceId].matchers = f 
        },
        removeSlience(hostSlience) {
            this.hostSlienceRules = this.hostSlienceRules.filter(v => v.id !== hostSlience.id)
        },
        showInput(name) {
            this.$set(this.inputVisible, name, true)
            // this.inputVisible[name]= true
            console.log(this.$refs)
            console.log(this.$refs["saveTagInput"+name])
            this.$nextTick(_ => {
                console.log("111111",this.$refs)
                console.log("222222",this.$refs["saveTagInput"+name])
                console.log(this.inputVisible[name])
                this.$refs["saveTagInput"+name][0].$refs.input.focus();
            });
        },
        handleInputConfirm(hostSlienceId) {
            console.log("111-->",hostSlienceId,this.hostSlienceRules)
            let inputValue = this.inputValue;
            if (inputValue) {
                this.hostSlienceRules[hostSlienceId].matchers.push(inputValue);
            }
            console.log("222-->")
            this.inputVisible = {};
            console.log("333-->")
            this.inputValue = '';
            console.log("444-->")
        
        }
    }
}
</script>

<style>
</style>