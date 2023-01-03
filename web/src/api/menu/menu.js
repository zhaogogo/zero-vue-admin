import {service} from "@/utils/request"

export const allMenu = (data) => {
    return service({
        url: '/api/v1/system/menu/all',
        method: "GET"
    })
}

export const addMenu = (data) => {
    return service({
        url:"/api/v1/system/menu/create",
        method:"POST",
        data
    })
}

export const menuInfo = (id) => {
    return service({
        url:"/api/v1/system/menu/" + id,
        method: "GET"
    })
}

export const deleteMenu = (id) => {
    return service({
        url: "/api/v1/system/menu/" + id,
        method: "DELETE"
    })
}

export const updateMenu = (data) => {
    return service({
        url: "/api/v1/system/menu/" + data.id,
        method: "PUT",
        data
    })
}

export const updateUsermenuParam = (data) => {
    return service({
        url: "/api/v1/system/menu/" + data.id + "/userparam",
        method: "PUT",
        data
    })
}