<template>
    <div>
        <div class="clearflex">
            <el-input style="width:75%;margin:0px 12px 12px" class="fl-left" placeholder="过滤" v-model="filterText"></el-input>
            <el-button class="fl-right" size="small" type="primary" @click="relation">确 定</el-button>
        </div>
        <el-tree
            v-loading="loading"
            ref="menuTree"
            :data="menuTreeData"
            :default-checked-keys="menuTreeIds"
            :props="menuDefaultProps"
            :filter-node-method="filterNode"
            default-expand-all
            highlight-current
            show-checkbox
            node-key="id"
            @check="nodeChange"
            
        >

        </el-tree>
    </div>
</template>

<script>
import {
    allMenu
} from "@/api/menu/menu"

import {
    getroleMenu,
    replaceRoleMenuPermission
} from "@/api/role/role"
export default {
    name:"menuPermission",
    props:{
        row: {
            default: function(){
                return {}
            },
            type: Object
        }
    },
    data(){
        return {
            filterText: "",
            menuTreeData: [],
            menuTreeIds: [],
            needConfirm: false,
            loading: true,
            menuDefaultProps: {
                children: "children",
                label: function(data) {
                    return data.meta.title
                }
            }
        }
    },
    watch: {
      filterText(val) {
        this.$refs.menuTree.filter(val);
      }
    },
    async created(){        
        const res2 = await getroleMenu(this.row)
        const menus = res2.list 

        const res = await allMenu()
        this.menuTreeData = res.list
        const arr = []
        menus.map(item=>{
            // 防止直接选中父级造成全选
            if (!menus.some(same => same.parentId === item.id)) {
                arr.push(Number(item.id))
            }
        })
        this.menuTreeIds = arr
        this.loading = false
    },
    methods: {
        nodeChange() {
            this.needConfirm = true
        },
        enterAndNext() {
            this.relation()
        },
        filterNode(value, data) {
            if (!value) return true;
            return data.meta.title.indexOf(value) !== -1;
        },
        async relation(){
            const checkArr = this.$refs.menuTree.getCheckedNodes(false,true)
            var menuids = []
            checkArr.map(item => {
                menuids.push(item.id)
            })
            console.log(menuids)
            console.log(checkArr)
            console.log({"id": this.row.id, "menuIdList": menuids})
            const res = await replaceRoleMenuPermission({"id": this.row.id, "menuIdList": menuids})
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