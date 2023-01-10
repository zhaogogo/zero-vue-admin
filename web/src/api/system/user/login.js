import {service} from '@/utils/request'

export const login = (data) => {
    return service({
        url: "/api/v1/system/user/login",
        method: 'post',
        data
    })
}