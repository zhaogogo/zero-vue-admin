import {service} from "@/utils/request"

export const allMenu = (data) => {
    return service({
        url: '/api/v1/system/menu/allmenu',
        method: "GET"
    })
}

export const addMenu = (data) => {
    return service({
        url:"/api/v1/system/menu/addmenu",
        method:"POST",
        data
    })
}