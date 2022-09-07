import Vue from "vue";
import Router from 'vue-router'

Vue.use(Router)

// // 获取原型对象上的push函数
// const originalPush = Router.prototype.push
// // 修改原型对象中的push方法
// Router.prototype.push = function push(location) {
//     return originalPush.call(this,location).catch(err => err)
// }


const constantRoutes = [
    {
        path: '/',
        redirect: "/login"
    },
    {
        path: '/login',
        name: 'login',
        component: () => import('@/view/login/index.vue')
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