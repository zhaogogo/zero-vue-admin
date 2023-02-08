<template>
    <div>
        <el-table
            loading="loading"
            :data="tableData"
            border
            stripe
        >
            <el-table-column label="name" prop="name">
                <template slot-scope="scope">
                    <div>{{ scope.row.name }}{{ scope.row.operator }}{{ scope.row.value }}</div>
                </template>
            </el-table-column>
            <el-table-column label="tag" prop="tag"></el-table-column>
            <el-table-column label="type" prop="type"></el-table-column>
            <el-table-column label="group" prop="group"></el-table-column>
            <el-table-column label="to" prop="to">
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