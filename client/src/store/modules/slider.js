const state = {
    articles: null,
    activeElement: 1,
}
const getters = {
    activeElement({ activeElement }) {
        return activeElement
    },
    articles({ articles }) {
        return articles
    },
}
const mutations = {
    next(state) {
        if (state.articles.length - 2 === state.activeElement) {
            state.activeElement = 1
        } else {
            state.activeElement++
        }
    },
    prev(state) {
        if (1 === state.activeElement) {
            state.activeElement = state.articles.length - 2
        } else {
            state.activeElement--
        }
    },
    setArticles(state, payload) {
        const articles = [payload[payload.length - 1], ...payload, payload[0]]
        state.articles = articles
    },
}
const actions = {
    async getArticles({ commit }, payload) {
        const res = await fetch(
            `${process.env.VUE_APP_API_URL}/top-articles?limit=${payload}`
        )
        const json = await res.json()
        commit('setArticles', json)
        console.log(json)
    },
}
export default {
    state,
    getters,
    mutations,
    actions,
}
