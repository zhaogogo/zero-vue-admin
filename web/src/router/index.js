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

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
    const newRouter = createRouter()
    router.matcher = newRouter.matcher // reset router
  }

export default router