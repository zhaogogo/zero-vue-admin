import router from '@/router'
import store  from '@/store'
import Nprogress from 'nprogress'
let asyncRouterFlag = 0

const whiteList = ['login']

router.beforeEach(async(to, from, next) => {
    Nprogress.start()
    const token = store.getters["user/token"]
    // 在白名单中的判断情况
    // 修改网页标签名称
    document.title = to.meta.title || "go-zero"
    if (whiteList.indexOf(to.name) > -1) {
        if (token) { //如果已经登陆
            next({name: store.getters["user/userPageSet"].defaultRouter})
        }else {
            next()
        }
    }else {
        // 不在白名单中并且已经登陆的时候
        if(token) {
            // 添加asyncRouterFlag防止多次获取动态路由和栈溢出
            if (!asyncRouterFlag&&store.getters["router/asyncRouters"].length == 0) {
                asyncRouterFlag++
                await store.dispatch('router/set_async_router')
                await store.dispatch("user/getUserInfo")
                const asyncRouter = store.getters['router/asyncRouters']
                router.addRoutes(asyncRouter)
                next({...to,replace: true})
            }else {
                if(to.matched.length) {
                    next()
                }else {
                    next({path:'/layout/404'})
                }
            }
        }
            // 不在白名单中并且未登陆的时候
        if(!token) {
            next({
                name: "login",
                query: {
                    redirect: document.location.hash
                }
            })
        }
    }
})

router.afterEach(()=>{
    // 路由加载完成后关闭进度条
    Nprogress.done()
})

router.onError(()=>{
    // 路由发生错误后销毁进度条
    Nprogress.remove()
})