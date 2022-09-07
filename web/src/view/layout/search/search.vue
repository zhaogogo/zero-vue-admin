<template>
    <div>
        <transition name="el-fade-in-linear">
            <div v-show="show" style="display: inline-block;float: left;">
                <el-select
                    ref="search-input"
                    v-model="value"
                    filterable
                    placeholder="请选择"
                    @blur="hiddenSearch"
                    @change="changeRouter"
                >
                    <el-option
                        v-for="item in routerList"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    ></el-option>
                </el-select>
            </div>
        </transition>
        <div
            v-show="!show"
            :style="{display:'inline-block',float:'left'}"
        >
            <i :style="{cursor: 'pointer'}" class="el-icon-search" @click="showSearch()"></i>
        </div>
        <div
            v-show="!show"
            style="display:inline-block;float: left;"
        >
            <i :style="{cursor:'pointer',paddingLeft:'1px'}" class="el-icon-refresh reload" :class="[reload?'reloading':'']" @click="handleReload"></i>
        </div>
    </div>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
    name: "Search",
    data() {
        return {
            value: "",
            reload: false,
            show: false
        }
    },
    computed: {
        ...mapGetters("router",["routerList"]),
    },
    methods:{
        changeRouter() {
            this.$router.push({name: this.value})
            this.value = ''
        },
        hiddenSearch(){
            this.show = false
        },
        showSearch(){
            this.show = true
            this.$nextTick(()=>{
                this.$refs["search-input"].focus()
            })
        },
        handleReload(){
            // this.reload = true
            // this.$bus.$emit('reload')
            // setTimeout(()=>{
            //     this.reload = false
            // },500)
        }
    }
}
</script>

<style>
.el-select .el-input__inner {
    padding-right: 0;
}
</style>