import {service} from "@/utils/request"


export const getAllRole = () => {
    return service({
        url: "/api/v1/system/role/all",
        method:"GET"
    })
}