export default {
    data() {
        return {
            page: 1,
            total: 0,
            pageSize: 10,
            tableData: [],
            searchInfo: {}
        }
    },
    methods: {
        handlePageSizeChange(val) {
            this.pageSize = val
            this.getTableData()
        },
        handlePageChange(val){
            this.page = val
            this.getTableData()
        },
        async getTableData(page = this.page, pageSize = this.pageSize) {
            const table = await this.listApi({page, pageSize, ...this.searchInfo})
            if (table.code === 200) {
                this.tableData = table.list
                this.total = table.total
                this.page = table.page
                this.pageSize = table.pageSize
            }
        }
    },
}