<template>
    <div>
        <el-cascader
            v-model="currentState.env"
            :options="envOptions"
            @change="selectEnv">
        </el-cascader>

        <div>
            <el-link v-show="currentState.host.length != 0" @click.prevent="selectEnv([this.currentState.env,currentState.host])">{{ currentState.host }}/</el-link>
            <el-link 
                v-for="(item,index ) in currentState.path" 
                :key="index"
                @click.prevent="toFile(index)"
            >{{ item }}</el-link>
        </div>
        <el-table
            :data="tableData"
            v-loading="loading"
            border
            stripe
            @sort-change="sortChange"
        >
            <el-table-column label="size" min-width="60" prop="size"></el-table-column>
            <el-table-column label="lastModified" min-width="60" prop="lastModified">
                <template slot-scope="scope">
                    <div>{{ Unix(scope.row.lastModified) }}</div>
                </template>
            </el-table-column>
            <el-table-column label="name" min-width="60" prop="name">
                <template slot-scope="scope">
                    <el-link type="primary" :disabled="scope.row.isFile" @click="file(scope.row.name)">{{ scope.row.name }}</el-link>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>

<script>
import infoList from '@/mixins/infoList'
import {
    fileEnvSelect,
    fileList
} from '@/api/monitoring/store/s3FileManager'
import { Unix } from '@/utils/formatTimeStamp'
export default {
    name: "s3FileManager",
    mixins: [infoList],
    data(){
        return {
            listApi: fileList,
            envOptions: [],
            currentState: {
                env: "",
                host: "",
                bucket: "",
                path: []
            }
        }
    },
    async created(){
        const selectOption = await fileEnvSelect()
        if (selectOption.code === 200) {
            this.envOptions = selectOption.list
        }
    },
    methods: {
        selectEnv(v){
            this.currentState.host = v[1]
            this.currentState.env = v[0]
            this.searchInfo.host = this.currentState.host
            this.searchInfo.path = this.currentState.path.join("")
            this.getTableData()
        },
        sortChange({ prop, order }){
            if (prop) {
                this.searchInfo.orderKey = prop
                this.searchInfo.order = order
            }
            this.getTableData()
        },
        Unix(sec) {
            return Unix(sec)
        },
        file(path) {
            this.currentState.path.push(path)
            this.searchInfo.host = this.currentState.host
            this.searchInfo.path = "s3://" + this.currentState.path.join("")
            console.log("---->", this.searchInfo.path)
            this.getTableData()
        },
        toFile(index) {
            this.searchInfo.host = this.currentState.host
            this.currentState.path = this.currentState.path.slice(0,index+1)
            this.searchInfo.path = "s3://" + this.currentState.path.join("")
            this.getTableData()
        }
    }
}
</script>

<style>

</style>