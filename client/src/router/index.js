import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Today from '../views/Today'
import Posts from '../views/Posts'
import Article from '../components/Article/Article'
import ArticlesList from '../components/Article/ArticlesList'

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
    {
        path: '/posts',
        component: Posts,
        children: [
            {
                path: '',
                name: 'Posts',
                component: ArticlesList,
            },
            {
                path: ':id',
                name: 'Post',
                component: Article,
            },
        ],
    },
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes,
})

export default router
