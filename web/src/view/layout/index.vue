<template>
  <el-container class="layout-cont">
    <el-container>
    <el-aside width="auto" class="main-cont main-left">
        <Aside></Aside>
    </el-aside>
      <el-main class="main-cont main-right"
      :style="{marginLeft:`${isCollapse?'54px':'220px'}`}"
      >
        <div
          :style="{width:`calc(100% - ${isCollapse?'54px':'220px'})`}"
          class="topfix"
        >
          <el-row>
            <el-header class="header-cont">
              <el-col :xs="2" :gl="1" :md="1" :sm="1" :xl="1">
                <div class="menu-total" @click="totalCollapse">
                  <i v-if="isCollapse" class="el-icon-s-unfold" />
                  <i v-else class="el-icon-s-fold" />
                </div>
              </el-col>
              <el-col :xs="10" :lg="14" :md="14" :sm="9" :xl="14">
                <el-breadcrumb class="breadcrumb" separator-class="el-icon-arrow-right">
                  <el-breadcrumb-item
                    v-for="item in matched.slice(1,matched.length)"
                    :key="item.path"
                  >{{ item.meta.title }}</el-breadcrumb-item>
                </el-breadcrumb>
              </el-col>
              <el-col :xs="12" :lg="9" :md="9" :sm="14" :xl="9">
                <div class="fl-right right-box">
                  <Search />
                  <el-dropdown>
                    <span class="header-avatar" style="cursor: pointer">
                      <span style="margin-left: 5px">{{ userPageSet.nick_name }}</span>
                      <i class="el-icon-arrow-down" />
                    </span>
                    <el-dropdown-menu slot="dropdown" class="dropdown-group">
                      <el-dropdown-item icon="el-icon-s-custom" @click.native="toPerson">个人信息</el-dropdown-item>
                      <el-dropdown-item icon="el-icon-table-lamp" @click.native="LOGOUT">登 出</el-dropdown-item>
                    </el-dropdown-menu>
                  </el-dropdown>
                </div>
              </el-col>
            </el-header>
          </el-row>
          <HistoryComponent></HistoryComponent>
        </div>
        <router-view v-loading="loadingFlag" element-loading-text="正在加载中" class="admin-box"></router-view>
        <BottomInfo></BottomInfo>
        <Setting></Setting>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'
import Aside from './aside/index.vue'
import HistoryComponent from './aside/historyComponent/history.vue'
import Search from './search/search.vue'
import BottomInfo from  './bottomInfo/index.vue'
import Setting from './setting/index.vue'

export default {
    name:"Layout",
    components: {
        Aside,
        Search,
        BottomInfo,
        Setting,
        HistoryComponent
    },
    data() {
      return {
        isCollapse: false,
        loadingFlag: false
      }
    },
    computed: {
      ...mapGetters("user",["userPageSet"]),
      matched() {
        return this.$route.matched
      }
    },
    methods: {
      ...mapMutations("user",["LOGOUT"]),
      totalCollapse(){
        this.isCollapse = !this.isCollapse
        this.$bus.emit("collapse", this.isCollapse)
      },
      toPerson(){
        this.$router.push({name:"persionhome"})
      }
    }
}
</script>

<style lang="scss">
@import '@/styles/basics.scss';
  // .el-header, .el-footer {
  //   background-color: #ffffff;
  //   color: #333;
  //   text-align: center;
  //   line-height: 60px;
  // }
  
  // .el-aside {
  //   background-color: #D3DCE6;
  //   color: rgb(255, 255, 255);
  //   text-align: center;
  //   line-height: 200px;
  //   height: 100vh;
  //   padding: 0;
  //   margin: 0;
  // }
  
  // .el-main {
  //   background-color: #E9EEF3;
  //   color: #333;
  //   text-align: center;
  //   line-height: 160px;
  // }
  
  // body > .el-container {
  //   margin-bottom: 40px;
  // }
  
  // .el-container:nth-child(5) .el-aside,
  // .el-container:nth-child(6) .el-aside {
  //   line-height: 260px;
  // }
  
  // .el-container:nth-child(7) .el-aside {
  //   line-height: 320px;
  // }
  // .el-breadcrumb {
  //     padding: 0 5px;
  //     line-height: inherit;
  // }

</style>