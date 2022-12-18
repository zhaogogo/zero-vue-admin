import {service} from '@/utils/request.js'

export const getPermission = () => {
    return service({
        url: "/api/v1/system/menu/usermenus",
        method: "GET"
    })
}

export const all = () => {
    return service({
        url:"/api/v1/system/user/all",
        method: "GET"
    })
}

export const paging = (data) => {
    return service({
        url: "/api/v1/system/user/paging",
        method: "POST",
        data
    })
}

export const detail = (id) => {
    return service({
        url:"/api/v1/system/user/" + id,
        method: "GET"
    })
}

export const updatepassword = (data) => {
    return service({
        url: "/api/v1/system/user/password/" + data.id,
        method: "POST",
        data
    })
}

export const updateRole = (data) => {
    return service({
        url: "/api/v1/system/user/role/" + data.id,
        method: "POST",
        data
    })
}

export const deleteSoft = (data) => {
    return service({
        url: "/api/v1/system/user/soft/" + data.id,
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
        method: "POST",
        data
    })
}

