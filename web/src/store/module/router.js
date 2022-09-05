import {getPermission} from "@/api/user/user"
import {asyncRouterHandle} from '@/utils/asyncRouter'


const state = {
    asyncRouters: []
}

const actions = {
    async set_async_router(ctx, value) {
        const baseRouter = [{
            path: '/layout',
            name: 'layout',
            component: 'view/layout/index.vue',
            meta: {
                title: '底层layout'
            },
            children: []
        }]
        const userInfoMenus = await getPermission()
        if (userInfoMenus.status != 200) {
            return 
        }

        ctx.commit("user/SETUSERINFO", userInfoMenus.data.userInfo,{root:true})
        const asyncRouter = userInfoMenus.data && userInfoMenus.data.menuList
        //返回的是一个数组 因为第二个操作数是 数组 且第一个操作数是对象为true
        // console.log(asyncRouter)
        asyncRouter.push({
            path: '404',
            name: '404',
            hidden: true,
            meta: {
                title: '迷路了*.*'
            },
            component: 'view/error/index.vue'
        })

        asyncRouterHandle(asyncRouter)
        baseRouter[0].children = asyncRouter
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
    }
}

export const router = {
    namespaced: true,
    state,
    actions,
    mutations,
    getters
}