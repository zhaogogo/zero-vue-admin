import {getToken} from '@/api/user/auth'

export const user = {
    namespaced: true,
    state,
    actions,
    mutations,
    getters
}

const state = {
    token: getToken(),
    userinfo: {
        id: 0,
        name: "",
        defaultrouter: "首页",
        sideMode: 'dark',
        activeColor: '#1890ff',
        baseColor: "#fff"
    },
    routers: []
}

const actions = {
    login(ctx, userInfo) {

    }
}

const mutations = {}

const getters = {
    token: state => state.token,
    userinfo: state => state.userinfo
}