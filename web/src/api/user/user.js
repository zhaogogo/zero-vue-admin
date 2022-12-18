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

export const getUserInfo = (id) => {
    return service({
        url:"/api/v1/system/user/" + id,
        method: "GET"
    })
}

export const getAllUser = () => {
    return service({
        url:"/api/v1/system/user/alluser",
        method: "GET"
    })
}

export const changeUserPassword = (data) => {
    return service({
        url: "/api/v1/system/user/changeUserPassword",
        method: "post",
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
    console.log(data)
    return service({
        url: "/api/v1/system/user/softdelete",
        method: "DELETE",
        data
    })
}

export const deleteUser = (data) => {
    return service({
        url: "/api/v1/system/user/delete",
        method:"DELETE",
        data
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
        url: "/api/v1/system/user/addUser",
        method:"POST",
        data
    })
}

export const editUser = (data) => {
    return service({
        url:"/api/v1/system/user/editUser",
        method: "POST",
        data
    })
}

