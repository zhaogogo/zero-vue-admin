import {service} from '@/utils/request'

export const login = (data) => {
    return service({
        url: "/api/user/login",
        method: 'post',
        data
    })
}