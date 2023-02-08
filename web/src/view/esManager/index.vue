<template>
    <div>
        <el-select v-if="$route.path !== '/layout/esManager/esConn'" v-model.number="esConnID" placeholder="选择ES连接" @change.passive="esConnChange">
            <el-option
                v-for="item in options"
                :key="item.id"
                :label="item.describe"
                :value="item.id"
            >
            </el-option>
        </el-select>
        <router-view></router-view>
    </div>
</template>

<script>
import { esConnAll } from '@/api/esManager/conn/conn.js'
export default {
    name: "esManager",
    data() {
        return {
            options: [],
            esConnID: ""
        }
    },
    async created(){
        var esopts = []
        const res =  await esConnAll()
        res.list && res.list.map(item => {
            esopts.push({
                id: item.id,
                describe: item.describe
            })
        })
        this.options = esopts
    },
    methods:{
        esConnChange(v) {
            this.$store.commit("esconn/SETESCONN",v)
        }
    }
}
</script>

<style>

</style>