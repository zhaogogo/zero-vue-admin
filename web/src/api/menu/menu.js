import {service} from "@/utils/request"

export const allMenus = (data) => {
    console.log("data",data)
    return service({
        url: '/api/v1/usercenter/allmenus',
        method: "GET"
    })
}