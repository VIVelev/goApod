const state = {
    images: [
        'https://images.unsplash.com/photo-1494548162494-384bba4ab999?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1000&q=80',
        'https://www.gettyimages.co.uk/gi-resources/images/500px/983801190.jpg',
        'https://www.gettyimages.com/gi-resources/images/500px/983794168.jpg',
        'https://images.unsplash.com/photo-1535332371349-a5d229f49cb5?ixlib=rb-1.2.1&w=1000&q=80',
        'https://cdn.spacetelescope.org/archives/images/wallpaper2/heic2007a.jpg',
        'https://images.unsplash.com/photo-1538370965046-79c0d6907d47?ixlib=rb-1.2.1&w=1000&q=80',
        'https://s23527.pcdn.co/wp-content/uploads/2020/01/100k-moon.jpg.optimal.jpg',
        'https://images.unsplash.com/photo-1494548162494-384bba4ab999?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&w=1000&q=80',
        'https://www.gettyimages.co.uk/gi-resources/images/500px/983801190.jpg',
    ],
    activeElement: 1,
}
const getters = {
    activeElement({ activeElement }) {
        return activeElement
    },
    images({ images }) {
        return images
    },
}
const mutations = {
    next(state) {
        if (state.images.length - 2 === state.activeElement) {
            state.activeElement = 1
        } else {
            state.activeElement++
        }
    },
    prev(state) {
        if (1 === state.activeElement) {
            state.activeElement = state.images.length - 2
        } else {
            state.activeElement--
        }
    },
}
const actions = {}
export default {
    state,
    getters,
    mutations,
    actions,
}
