import {login} from '@/api/user/login'
import {setUserPageSet,currentUserInfo} from '@/api/user/user'
import router from '@/router/index'
import { Message } from 'element-ui'

const getDefaultState = () =>{
    return {
        token: "",
        userPageSet: {
            name: "",
            nick_name: "",
            avatar: "",
            defaultRouter: "",
            sideMode: '#191a23',
            activeTextColor: '#1890ff',
            textColor: "#fff"
        },
        roles: [],
        current_role: {}
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
            ctx.commit("SETUSERROLE",res.roles)
            router.push({name: ctx.getters["userPageSet"].defaultRouter })
            return true
        }else {
            return false
        }
    },
    async changeSideMode(ctx, data) {
        ctx.commit("CHANGESIDEMODE", data)
        const userpageset = await setUserPageSet(ctx.state.userPageSet)
        if (userpageset.code !== 200) {
            Message({
                type: "error",
                message: userpageset.msg
            })
        }
    },
    async changeTextColor(ctx, data) {
        ctx.commit("CHANGETEXTCOLOR", data)
        const userpageset = await setUserPageSet(ctx.state.userPageSet)
        if (userpageset.code !== 200) {
            Message({
                type: "error",
                message: userpageset.msg
            })
        }
    },
    async changeActiveColor(ctx, data) {
        ctx.commit("CHANGEACTIVECOLOR", data)
        const userpageset = await setUserPageSet(ctx.state.userPageSet)
        if (userpageset.code !== 200) {
            Message({
                type: "error",
                message: userpageset.msg
            })
        }
    },
    async getUserInfo(ctx, data) {
        const res = await currentUserInfo()
        ctx.commit("SETUSERPAGESET", res.userPageSet)
        ctx.commit("SETUSERROLE", res.roles)
        ctx.commit("SETCURRENTROLE", res.currentRole)
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
    SETUSERROLE(state,rolelist) {
        // for (var index in rolelist) {
        //     console.log("===>", index, rolelist[index])
        //     this.$set(state.roles, index, rolelist[index])
        // }
        state.roles = rolelist
    },
    SETCURRENTROLE(state,role) {
        state.current_role = role
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
    },
    roles(state) {
        return state.roles
    },
    currentRole(state) {
        return state.current_role
    }
}

export const user = {
    namespaced: true,
    state,
    actions,
    mutations,
    getters
}