import {getPermission} from "@/api/user/user"
import {asyncRouterHandle} from '@/utils/asyncRouter'

const routerList = []
const formatRouter = (routes) => {
  routes && routes.map(item => {
    if ((!item.children || item.children.every(ch => ch.hidden)) && item.name !== '404') {
      routerList.push({ label: item.meta.title, value: item.name })
    }
    if (item.children && item.children.length > 0) {
      formatRouter(item.children)
    }
  })
}

const state = {
    asyncRouters: [],
    routerList: routerList
}

const actions = {
    async set_async_router(ctx, value) {
        const baseRouter = [
            {
                path: '/layout',
                name: 'layout',
                component: 'view/layout/index.vue',
                meta: {
                    title: '底层layout'
                },
                children: []
            }
        ]
        const userMenus = await getPermission()
        if (userMenus.code != 200) {
            return false
        }
        var asyncRouter = userMenus && userMenus.menus
        //返回的是一个数组 因为第二个操作数是 数组 且第一个操作数是对象为true
        if (asyncRouter === null) {
            var asyncRouter = new Array()
        }
        asyncRouter.push({
            path: '404',
            name: '404',
            hidden: true,
            meta: {
                title: '迷路了*.*'
            },
            component: 'view/error/index.vue'
        })
        asyncRouter.push({
            path: 'persionhome',
            name: 'persionhome',
            hidden: true,
            meta: {
                title: 'home'
            },
            component: 'view/home/index.vue'
        })
        formatRouter(asyncRouter)
        baseRouter[0].children = asyncRouter
        asyncRouterHandle(baseRouter)
        ctx.commit("SET_ASYNC_ROUTER", baseRouter)
        return true
    }
}

const mutations = {
    SET_ASYNC_ROUTER(state, value) {
        state.asyncRouters = value
    }
}

const getters = {
    asyncRouters(state) {
        return state.asyncRouters
    },
    routerList(state) {
        return state.routerList
    }
}

export const router = {
    namespaced: true,
    state,
    actions,
    mutations,
    getters
}