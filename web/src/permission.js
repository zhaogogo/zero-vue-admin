import router from '@/router'
import store  from '@/store'

let asyncRouterFlag = 0

const whiteList = ['login']

router.beforeEach(async(to, from, next) => {
    const token = store.getters["user/token"]
    // 在白名单中的判断情况
    // 修改网页标签名称
    document.title = to.meta.title || "go-zero"
    if (whiteList.indexOf(to.name) > -1) {
        if (token) {
            console.log("denglu1")
            next({name: store.getters["user/userPageSet"].defaultRouter})
        }else {
            console.log("denglu2")
            next()
        }
    }else {
        // 不在白名单中并且已经登陆的时候
        if(token) {
            // 添加flag防止多次获取动态路由和栈溢出
            if (!asyncRouterFlag&&store.getters["router/asyncRouters"].length == 0) {
                asyncRouterFlag++
                await store.dispatch('router/set_async_router')
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
            console.log("denglu1111")
            next({
                name: "login",
                query: {
                    redirect: document.location.hash
                }
            })
        }
    }
})