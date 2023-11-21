import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
    {
        path: "/",
        redirect: '/login'
    },
    {
        path: "/login",
        name: "Login",
        component: () => import("@/view/login/index.vue")
    },
    {
        path: "/:catchAll(.*)",
        meta: {
            closeTab: true,
        },
        component: () => import("@/view/error/index.vue")
    }
]

const router = createRouter({
    history: createWebHashHistory(), //使用hash路由模式 /#/
    routes
})

export default router