<template>
    <div>
        <div class="clearflex">
            <el-button class="fl-right" size="small" type="primary" @click="relation">确 定</el-button>
        </div>
        <el-tree
            ref="menuTree"
            :data="menuTreeData"
            :default-checked-keys="menuTreeIds"
            :props="menuDefaultProps"
            default-expand-all
            highlight-current
            show-checkbox
            node-key="id"
        >

        </el-tree>
    </div>
</template>

<script>
import {
    allMenu
} from "@/api/menu/menu"

import {
    getroleMenu
} from "@/api/role/role"
export default {
    name:"menus",
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
            menuTreeData: [],
            menuTreeIds: [],
            menuDefaultProps: {
                children: "children",
                label: function(data) {
                    return data.meta.title
                }
            }
        }
    },
    async created(){
        const res = await allMenu()

        this.menuTreeData = res.list
        
        const res2 = await getroleMenu(this.row)
        const menus = res2.list 
        const arr = []
        menus.map(item=>{
            arr.push(Number(item.id))
        })
        this.menuTreeIds = arr
    },
    methods: {
        relation(){
            const a = this.$refs.menuTree.getCheckedKeys(false)
            const b =  this.$refs.menuTree.getHalfCheckedKeys()
            console.log("--->",a,b)
        }
    }
}
</script>

<style>

</style>