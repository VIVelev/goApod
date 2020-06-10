import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Today from '../views/Today'

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home,
    },
    {
        path: '/today',
        name: 'Today',
        component: Today,
    },
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes,
})

export default router
