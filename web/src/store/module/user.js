import {login} from '@/api/user/login'
import {setUserPageSet} from '@/api/user/user'
import router from '@/router/index'
import { Message } from 'element-ui'

const getDefaultState = () =>{
    return {
        token: "",
        userPageSet: {
            name: "",
            nick_name: "",
            avatar: "",
            defaultRouter: "dashboard",
            sideMode: '#191a23',
            activeTextColor: '#1890ff',
            textColor: "#fff"
        }
    }
}

const state = getDefaultState()

const actions = {
    async loginin(ctx, userInfo) {
        const res = await login(userInfo)  //响应拦截器错误也会传递到这里
        if (res.code === 200) {
            ctx.commit("SETTOKEN",res.token)
            await ctx.dispatch("router/set_async_router",{},{root:true})
            const asyncRouters = ctx.rootGetters['router/asyncRouters']
            router.addRoutes(asyncRouters)
            ctx.commit("SETUSERPAGESET", res.userPageSet)
            router.push({name: ctx.getters["userPageSet"].defaultRouter })
            return true
        }
    },
    async changeSideMode(ctx, data) {
        ctx.commit("CHANGESIDEMODE", data)
        const userpageset = await setUserPageSet(ctx.state.userPageSet)
        if (userpageset.code === 200) {
            Message({
                type: "success",
                message: userpageset.msg
            })
        }else {
            Message({
                type: "error",
                message: userpageset.msg
            })
        }
    },
    async changeTextColor(ctx, data) {
        ctx.commit("CHANGETEXTCOLOR", data)
        const userpageset = await setUserPageSet(ctx.state.userPageSet)
        if (userpageset.code === 200) {
            Message({
                type: "success",
                message: userpageset.msg
            })
        }else {
            Message({
                type: "error",
                message: userpageset.msg
            })
        }
    },
    async changeActiveColor(ctx, data) {
        ctx.commit("CHANGEACTIVECOLOR", data)
        const userpageset = await setUserPageSet(ctx.state.userPageSet)
        if (userpageset.code === 200) {
            Message({
                type: "success",
                message: userpageset.msg
            })
        }else {
            Message({
                type: "error",
                message: userpageset.msg
            })
        }
    }
}

const mutations = {
    SETTOKEN(state,token){
        state.token = token
    },
    LOGOUT(state){
        Object.assign(state,getDefaultState())
        sessionStorage.clear()
        router.push({name:"login", replace: true}).catch(() => {})
        window.location.reload()
    },
    SETUSERPAGESET(state,user) {
        state.userPageSet.avatar = user.avatar
        state.userPageSet.defaultRouter = user.defaultRouter
        state.userPageSet.sideMode = user.sideMode
        state.userPageSet.activeTextColor = user.activeTextColor
        state.userPageSet.textColor = user.textColor
        state.userPageSet.name = user.name
        state.userPageSet.nick_name = user.nick_name
    },
    CHANGESIDEMODE(state,sidemode){
        state.userPageSet.sideMode = sidemode
    },
    CHANGETEXTCOLOR(state,textColor) {
        state.userPageSet.textColor = textColor
    },
    CHANGEACTIVECOLOR(state,activeTextColor) {
        state.userPageSet.activeTextColor = activeTextColor
    }
}

const getters = {
    token: state => state.token,
    // userInfo(state) {
    userPageSet(state) {
        return state.userPageSet
    },
    sideMode(state) {
        return state.userPageSet.sideMode
    },
    textColor(state) {
        return state.userPageSet.textColor
    },
    activeTextColor(state) {
        return state.userPageSet.activeTextColor
    }
}

export const user = {
    namespaced: true,
    state,
    actions,
    mutations,
    getters
}