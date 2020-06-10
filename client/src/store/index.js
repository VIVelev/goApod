import Vue from 'vue'
import Vuex from 'vuex'
import slider from './modules/slider'
import user from './modules/user'
import popUp from './modules/popUp'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {},
    mutations: {},
    actions: {},
    modules: {
        slider,
        user,
        popUp,
    },
})
