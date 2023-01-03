<template>
    <div>
        <div class="clearflex">
            <el-input style="width:75%;margin:0px 12px 12px" class="fl-left" placeholder="过滤" v-model="filterText"></el-input>
            <el-button class="fl-right" size="small" type="primary" @click="relation">确 定</el-button>
        </div>
        <el-tree
            ref="apiTree"
            :data="apiTreeData"
            :default-checked-keys="apiTreeIds"
            :props="apiDefaultProps"
            :filter-node-method="filterNode"
            default-expand-all
            highlight-current
            node-key="onlyid"
            show-checkbox
            @check="nodeChange"
        ></el-tree>
  </div>
</template>

<script>
import {
    allAPI
} from '@/api/api/api.js'

import {
    getRoleAPI,
    updateRoleApiPermission
} from '@/api/role/role.js'

export default {
    name: "apiPermission",
    props: {
        row: {
            default: function(){
                return {}
            },
            type: Object
        }
    },
    data(){
        return {
            filterText:"",
            apiTreeData: [],
            apiTreeIds: [],
            needConfirm: false,
            apiDefaultProps: {
                children: 'children',
                label: 'describe'
            }
        }
    },
    watch: {
      filterText(val) {
        this.$refs.apiTree.filter(val);
      }
    },
    async created(){
        const allapi = await allAPI()
        const apis = allapi.list
        this.apiTreeData = this.buildApiTree(apis)

        const roleapi = await getRoleAPI(this.row)
        const arr = []
        roleapi.policy &&  roleapi.policy.map(item => {
            arr.push("p:"+item.api+"m:"+item.method)
        })
        this.apiTreeIds = arr
    },
    methods:{
        filterNode(value, data) {
            if (!value) return true;
            return data.describe.indexOf(value) !== -1;
        },
        //创建api树
        buildApiTree(apis){
            const apiObj = {}
            apis && apis.map(item => {
                // item.describe = item.describe + " - (" + item.api + ")"
                item.onlyid = "p:" + item.api + "m:" + item.method
                if (Object.prototype.hasOwnProperty.call(apiObj,item.group)){
                    apiObj[item.group].push(item)
                }else {
                    Object.assign(apiObj,{ [item.group]: [item] })
                }
            })
            const apiTree = []
            for (const key in apiObj) {
                const treeNode = {
                    ID: key,
                    describe: key + "权限",
                    children: apiObj[key]
                }
                apiTree.push(treeNode)
            }
            return apiTree
        },
        nodeChange() {
            this.needConfirm = true
        },
        async relation() {
            const checkArr = this.$refs.apiTree.getCheckedNodes(true)
            const res = await updateRoleApiPermission({"id":this.row.id, "casbinRules": checkArr})
            if (res.code === 200) {
                this.$message({
                    type: "success",
                    message: res.msg
                })
            }
        }
    }
}
</script>

<style>

</style>