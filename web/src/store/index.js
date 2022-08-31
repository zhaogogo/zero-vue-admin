import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersistence from "vuex-persist";
import { user } from './module/user'
import { router } from './module/router'
Vue.use(Vuex)

const vueLocal = new VuexPersistence({
    storage: window.localStorage,
    modules: ["userInfo"]
})

const store = new Vuex.Store({
    modules:{
        user,
        router
    },
    plugins: [vueLocal.plugin]
})

export default store
