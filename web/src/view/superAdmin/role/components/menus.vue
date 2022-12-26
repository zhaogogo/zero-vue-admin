<template>
    <div>

        <el-tree
            ref="menuTree"
            :data="menuTreeData"
            show-checkbox
        >

        </el-tree>
    </div>
</template>

<script>
import {
    allMenu
} from "@/api/menu/menu"
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
            menuTreeData: []
        }
    },
    async created(){
        const res = await allMenu()

        this.menuTreeData = res.list
        

    }
}
</script>

<style>

</style>