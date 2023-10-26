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
                    <el-tag  size="mini" v-for="item in scope.row.tags" :key="item">{{ item.key }}={{ item.value }}</el-tag>
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
                <el-table
                    :data="hostSlienceRule.list"
                >
                    <el-table-column label="slience_name" prop="slience_name"></el-table-column>
                </el-table>
                <div class="demo-drawer__footer">
                    <el-button @click="cancelForm">取 消</el-button>
                    <el-button type="primary" @click="$refs.drawer.closeDrawer()" :loading="loading">{{ loading ? '提交中 ...' : '确 定' }}</el-button>
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
        }
    },
    data() {
        return {
            listApi: hosts,
            hostLocationOptions: hostLocationOptions,
            slienceDialogVisible: false,
            slienceTitle: "",
            hostSlienceRule: {},
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
            this.slienceTitle = host + " slience"
            this.hostSlienceRule = await hostSlienceRule(host)
            this.slienceDialogVisible = true
        },
        closeSlienceDialog() {
            this.slienceDialogVisible = false
        }
    }
}
</script>

<style>
</style>