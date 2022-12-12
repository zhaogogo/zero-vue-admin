import {service} from '@/utils/request.js'

export const getPermission = () => {
    return service({
        url: "/api/v1/system/menu/usermenus",
        method: "GET"
    })
}

export const getPagingUser = (data) => {
    return service({
        url: "/api/v1/system/user/paginguser",
        method: "POST",
        data
    })
}

export const setUserPageSet = (data) => {
    return service({
        url: "/api/v1/usercenter/setuserpageset",
        method: "PUT",
        data: data
    })
}

export const getAllUser = (data) => {
    return service({
        url: "/api/v1/usercenter/getalluser",
        method: "GET",
    })
}

export const softDeleteUser = (data) => {
    console.log(data)
    return service({
        url: "/api/v1/usercenter/softdelete/",
        method: "DELETE",
        data
    })
}