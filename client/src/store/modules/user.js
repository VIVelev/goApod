const state = {
    user: null,
}
const getters = {
    user({ user }) {
        return user
    },
}
const mutations = {
    changeUser(state, payload) {
        console.log(payload)
        state.user = payload
    },
}
const actions = {
    register() {
        //TO DO: sent the name of the user to the backend
    },
    login() {
        //send the user information to the backend
    },
}

export default {
    state,
    getters,
    mutations,
    actions,
}
