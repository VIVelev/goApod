const state = {
    popUpShown: false,
}
const getters = {
    isPopUpShown({ popUpShown }) {
        return popUpShown
    },
}
const mutations = {
    modifyPopUp(state, payload) {
        state.popUpShown = payload
    },
}
const actions = {}

export default {
    state,
    getters,
    mutations,
    actions,
}
