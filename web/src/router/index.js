import Vue from "vue";
import Router from 'vue-router'

Vue.use(Router)

export const constantRoutes = [
    {
        path: '/',
        redirect: "/login"
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/login/index.vue')
    }
]

const createRouter = () => new Router({
    scrollBehavior: () => ({y:0}), 
    routes: constantRoutes
})

const router = createRouter()

export default router