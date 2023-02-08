import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersistence from "vuex-persist";

import { user } from './module/user'
import { router } from './module/router'
import { esconn } from './module/esconn.js'

Vue.use(Vuex)

const vuexLocal = new VuexPersistence({
    storage: window.localStorage,
    modules: ["user"]
})

const store = new Vuex.Store({
    modules:{
        user,
        router,
        esconn
    },
    plugins: [vuexLocal.plugin]
})

export default store
