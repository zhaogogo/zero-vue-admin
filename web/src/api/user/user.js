import {service} from '@/utils/request.js'

export const getPermission = () => {
    return service({
        url: "/api/v1/system/menu/usermenus",
        method: "GET"
    })
}

export const currentUserInfo = () => {
    return service({
        url:"/api/v1/system/user/currentset",
        method: "GET"
    })
}

export const allUser = () => {
    return service({
        url:"/api/v1/system/user/all",
        method: "GET"
    })
}

export const pagingUser = (data) => {
    return service({
        url: "/api/v1/system/user/paging",
        method: "POST",
        data
    })
}

export const userDetail = (id) => {
    return service({
        url:"/api/v1/system/user/" + id,
        method: "GET"
    })
}

export const updatepassword = (data) => {
    return service({
        url: "/api/v1/system/user/" + data.id + "/password",
        method: "PUT",
        data
    })
}

export const changeLoginPassword = (data) => {
    return service({
        url: "api/v1/system/user/password",
        method: "PUT",
        data
    })
}

export const updateRole = (data) => {
    return service({
        url: "/api/v1/system/user/" + data.id + "/role",
        method: "PUT",
        data
    })
}

export const deleteSoft = (data) => {
    return service({
        url: "/api/v1/system/user/" + data.id + "/soft",
        method: "DELETE",
        data
    })
}

export const deleteUser = (data) => {
    return service({
        url: "/api/v1/system/user/" + data.id,
        method:"DELETE"
    })
}

export const setUserPageSet = (data) => {
    return service({
        url: "/api/v1/system/user/page",
        method: "PUT",
        data: data
    })
}

export const addUser = (data) => {
    return service({
        url: "/api/v1/system/user/create",
        method:"POST",
        data
    })
}

export const editUser = (data) => {
    return service({
        url:"/api/v1/system/user/" + data.id,
        method: "PUT",
        data
    })
}

export const changeRole = (data) => {
    return service({
        url: "/api/v1/system/user/changerole",
        method: "POST",
        data
    })
}