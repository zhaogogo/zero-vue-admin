<template>
  <div style="height:100vh" :style="{backgroundColor:userPageSet.sideMode}">
    <el-scrollbar>
      <el-menu
          :collapse="isCollapse"
          :collapse-transition="true"
          :default-active="active" 
          class="el-menu-vertical"
          :background-color="userPageSet.sideMode"
          :active-text-color="userPageSet.activeTextColor"
          :text-color="userPageSet.textColor"
          unique-opened
          mode="vertical"
          @select="selectMenuItem"
      >
        <template v-for="item in asyncRouters[0].children">
          <aside-component v-if="!item.hidden" :key="item.name" :router-info="item" />
        </template>
      </el-menu>
    </el-scrollbar>
  </div>
</template>

<script>
import AsideComponent from './asideComponent'
import { mapGetters } from 'vuex';
export default {
    name: "Aside",
    components: {
      AsideComponent
    },
    data() {
      return {
        active : "",
        isCollapse: false
      };
    },
    computed:{
      ...mapGetters("router",["asyncRouters"]),
      ...mapGetters("user",["userPageSet"])
    },
    watch:{
        $route(){
            this.active = this.$route.name
        }
    },
    created() {
      this.active = this.$route.name
      this.$bus.on("collapse",item => {
        this.isCollapse = item
      })
    },
    beforeDestroy() {
      this.$bus.off("collapse")
    },
    methods:{
      selectMenuItem(index, b, ele){
        //路由跳转及参数配置
        //参数配置
        const query = {}
        const params = {}
        ele.route.parameters && 
        ele.route.parameters.map(item => {
          if (item.type === 'query') {
            query[item.key] = item.value
          } else {
            params[item.key] = item.value
          }
        })
        if (index === this.$route.name) return
        if (index.indexOf('http://') > -1 || index.indexOf('https://') > -1) {
          console.log("0000,window.open",index)
          window.open(index)
        } else {
          this.$router.push({name: index, query, params})
        }
      }
    }
}
</script>

<style>
  .el-menu-vertical:not(.el-menu--collapse) {
    width: 200px;
    min-height: 100vh;
  }
</style>