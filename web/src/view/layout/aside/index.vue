<template>
  <div style="height:100vh" :style="{backgroundColor:userInfo.sideMode}">
    <el-menu
        :collapse="isCollapse"
        :default-active="active" 
        class="el-menu-vertical"
        :background-color="userInfo.sideMode"
        :active-text-color="userInfo.activeTextColor"
        :text-color="userInfo.textColor"
        unique-opened
        @select="selectMenuItem"
    >
      <template v-for="item in asyncRouters[0].children">
        <aside-component v-if="!item.hidden" :key="item.name" :router-info="item" />
      </template>
    </el-menu>
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
    methods: {
      handleOpen(key, keyPath) {
        console.log(key, keyPath);
      },
      handleClose(key, keyPath) {
        console.log(key, keyPath);
      }
    },
    watch:{
        $route(){
            this.active = $route.name
        }
    },
    created() {
        this.active = this.$route.name
        this.$bus.on("collapse",item => {
          this.isCollapse = item
        })
    },
    destroyed() {
      this.$bus.off("collapse")
    },
    computed:{
      ...mapGetters("router",["asyncRouters"]),
      ...mapGetters("user",["userInfo"])
    },
    methods:{
      selectMenuItem(indexe, _, ele){
        //路由参数配置
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