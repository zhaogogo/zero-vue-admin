<template>
    <div>
        <el-button type="primary" class="drawer-container" icon="el-icon-setting" @click="showSettingDrawer"></el-button>
        <el-drawer
            title="用户页面配置"
            :visible.sync="drawer"
            :direction="direction"
            :beforce-close="handleClose"
        >
            <div class="setting_body">
                <div class="setting_card">
                    <div class="setting_title">侧边栏主题 (注：自定义请先配置背景色)</div>
                    <div class="setting_content">
                        <div class="them-box"></div>
                        <div class="color-box">
                            <div>
                                <div class="setting_title">自定义背景色</div>
                                <el-color-picker :value="sideMode" @change="changeMode"></el-color-picker>
                            </div>
                            <div>
                                <div class="setting_title">侧边栏文本颜色</div>
                                <el-color-picker :value="textColor" @change="changeTextColor"></el-color-picker>
                            </div>
                                <div class="setting_title">活跃色</div>
                                <el-color-picker :value="activeTextColor" @change="changeActiveColor"></el-color-picker>
                        </div>
                    </div>
                </div>
            </div>
        </el-drawer>
    </div>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
    name: "Setting",
    data(){
        return {
            drawer: false,
            direction:"rtl"
        }
    },
    computed: {
        ...mapGetters("user",["sideMode","textColor","activeTextColor"])
    },
    methods: {
        handleClose(){
            this.drawer = false
        },
        showSettingDrawer(){
            this.drawer = true
        },
        changeMode(e) {
            if (e === null) {
                return 
            }
            this.$store.dispatch("user/changeSideMode",e)
        },
        changeTextColor(e) {
            if (e === null) {
                return 
            }
            this.$store.dispatch("user/changeTextColor",e)
        },
        changeActiveColor(e) {
            if (e === null) {
                return 
            }
            this.$store.dispatch("user/changeActiveColor",e)
        }
    }
}
</script>

<style lang="scss" scoped>
#el-drawer__title {
    line-height: 54px !important;
}
.drawer-container {
    position: fixed;
    right: 0;
    bottom: 15%;
    height: 40px;
    width: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 999;
    color: #fff;
    border-radius: 4px 0 0 4px;
    cursor: pointer;
    -webkit-box-shadow: inset 0 0 6 rgba(0,0,0,10%);
}
.setting_body{
    padding: 20px;
    .setting_car{
        margin-bottom: 20px;
    }
}
</style>