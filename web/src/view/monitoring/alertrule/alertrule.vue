<template>
    <div>
        <el-table
            loading="loading"
            :data="tableData"
            border
            stripe
        >
            <el-table-column label="name" prop="name" min-width="60">
                <template slot-scope="scope">
                    <div>{{ scope.row.name }}{{ scope.row.operator }}{{ scope.row.value }}</div>
                </template>
            </el-table-column>
            <el-table-column label="tag" min-width="50" prop="tag"></el-table-column>
            <el-table-column label="type" min-width="50">
                <template slot-scope="scope">
                    <div>{{ scope.row.type }}/</div>
                    <el-tag size="mini">{{ scope.row.group }}</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="to" min-width="40" prop="to">
                <template slot-scope="scope">
                    <el-tag v-if="scope.row.to === 1" type="info">通用</el-tag>
                    <el-tag v-if="scope.row.to === 2" type="success">兆维</el-tag>
                    <el-tag v-if="scope.row.to === 3">亦庄</el-tag>
                </template>
            </el-table-column>
            <el-table-column label="expr" prop="expr" :show-overflow-tooltip="true">
                <template slot-scope="scope">
                    <div>{{ scope.row.expr }} + {{ scope.row.operator }} + {{ scope.row.value }}</div>
                </template>
            </el-table-column>
            <el-table-column label="操作" fixed="right">
                <template slot-scope="scope">
                    <el-button size="small" type="text" icon="el-icon-edit">编辑</el-button>
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
    </div>
</template>

<script>
import infoList from '@/mixins/infoList'
import {
    alertRules
} from '@/api/monitoring/alertrule/alertrule'
export default {
    name: "AlertRuleManager",
    mixins: [infoList],
    data() {
        return {
            listApi: alertRules
        }
    },
    async created() {
        await this.getTableData()
        console.log("===>", this.tableData)
    }
}
</script>

<style>

</style>