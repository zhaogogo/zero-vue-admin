const state = {
    esconn: 0
}

const actions = {}

const mutations = {
    SETESCONN(state, esconn) {
        state.esconn = esconn
    }
}

const getters = {
    getESConn(state) {
        return state.esconn
    }
}

export const esconn = {
    namespaced: true,
    state,
    actions,
    mutations,
    getters
}