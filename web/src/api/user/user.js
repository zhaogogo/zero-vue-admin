import {service} from '@/utils/request.js'

export const getPermission = () => {
    return service({
        url: "/api/v1/system/menu/usermenus",
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

export const userInfo = (id) => {
    return service({
        url:"/api/v1/system/user/detail/" + id,
        method: "GET"
    })
}



export const changeUserPassword = (data) => {
    return service({
        url: "/api/v1/system/user/password",
        method: "POST",
        data
    })
}

export const updateUserRole = (data) => {
    return service({
        url: "/api/v1/system/user/updateUserRole",
        method: "POST",
        data
    })
}

export const softDeleteUser = (data) => {
    return service({
        url: "/api/v1/system/user/deletesoft/" + data.id,
        method: "DELETE",
        data
    })
}

export const deleteUser = (data) => {
    return service({
        url: "/api/v1/system/user/delete/" + data.id,
        method:"DELETE"
    })
}

export const setUserPageSet = (data) => {
    return service({
        url: "/api/v1/system/user/setuserpageset",
        method: "PUT",
        data: data
    })
}

export const addUser = (data) => {
    return service({
        url: "/api/v1/system/user/add",
        method:"POST",
        data
    })
}

export const editUser = (data) => {
    return service({
        url:"/api/v1/system/user/edit/" + data.id,
        method: "POST",
        data
    })
}

