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