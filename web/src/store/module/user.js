import {login} from '@/api/user/login'
import {getToken,setToken,removeToken} from '@/api/user/auth'
import router from '@/router/index'

const getDefaultState = () =>{
    return {
        token: getToken(),
        userinfo: {
            id: 0,
            name: "",
            sideMode: 'dark',
            activeColor: '#1890ff',
            baseColor: "#fff"
        }
    }
}

const state = getDefaultState()

const actions = {
    async loginin(ctx, userInfo) {
        const res = await login(userInfo)  //响应拦截器错误也会传递到这里
        if (res.status === 200) {
            ctx.commit("SETTOKEN",res.data.token)
            setToken(res.data.token)
            await ctx.dispatch("router/set_async_router",{},{root:true})
            const asyncRouters = ctx.rootGetters['router/asyncRouters']
            router.addRoutes(asyncRouters)
            // console.log("router", router.getRoutes())
            router.push({name: ctx.getters["userInfo"].defaultRouter })
            return true
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
        removeToken()
        router.push({name:"login", replace: true})
        window.location.reload()
    },
    SETUSERINFO(state,user) {
        state.userinfo = user
    }
}

const getters = {
    token: state => state.token,
    userInfo(state) {
        return state.userinfo
    } 
}

export const user = {
    namespaced: true,
    state,
    actions,
    mutations,
    getters
}