import {service} from '@/utils/request.js'

export const getPermission = () => {
    return service({
        url: "/api/user/permission",
        method: "GET"
    })
}