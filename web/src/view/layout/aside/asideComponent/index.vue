<template>
  <component :is="menuSelect" v-if="!routerInfo.hidden" :router-info="routerInfo">
    <template v-if="routerInfo.children&&routerInfo.children.length">
        <AsideComponent v-for="item in routerInfo.children" :key="item.name" :router-info="item"></AsideComponent>
    </template>
  </component>
</template>

<script>
import MenuItem from './menuItem'
import AsyncSubmenu from './asyncSubmenu'
export default {
    name: "AsideComponent",
    components: {
        MenuItem,
        AsyncSubmenu
    },
    props: {
        routerInfo: {
            default: function() {
                return null
            },
            type: Object
        }
    },
    computed:{
        menuSelect() {
            if (this.routerInfo.children && this.routerInfo.children.filter(item => !item.hidden).length) {
                return 'AsyncSubmenu'
            }else {
                return "MenuItem"
            }
        }
    }
}
</script>

<style>

</style>