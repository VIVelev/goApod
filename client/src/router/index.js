import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Today from '../views/Today'
import Posts from '../views/Posts'
import ArticlePage from '../views/ArticlePage'
import ArticlesList from '../components/Article/ArticlesList'
import CreateArticle from '../views/CreateArticle'

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
                component: ArticlePage,
            },
        ],
    },
    {
        path: '/create-article',
        component: Today,
        component: CreateArticle,
    },
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes,
})

export default router
