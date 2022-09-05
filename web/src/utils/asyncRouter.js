import _import from './_import_production'
export const asyncRouterHandle = (asyncRouter) => {
    asyncRouter.map(item => {
        if (item.component) {
            item.component = _import(item.component)
        }else {
            delete item["component"]
        }

        if (item.children) {
            asyncRouterHandle(item.children)
        }
    })
}