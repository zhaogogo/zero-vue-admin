import {service} from "@/utils/request"


export const getAllRole = () => {
    return service({
        url: "/api/v1/system/role/all",
        method:"GET"
    })
}

export const refreshPermission = () => {
    return service({
        url: "/api/v1/system/role/refreshpermission",
        method: "GET"
    })
}

export const detailRole = (data) => {
    return service({
        url: "/api/v1/system/role/" + data.id,
        method: "GET"
    })
}

export const createRole = (data) => {
    return service({
        url: "/api/v1/system/role/create",
        method: "POST",
        data
    })
}

export const deleteRole = (data) =>{
    return service({
        url: "/api/v1/system/role/" + data.id,
        method: "DELETE"
    })
}

export const deleteSoftRole = (data) => {
    return service({
        url: "/api/v1/system/role/" + data.id + "/soft",
        method: "DELETE",
        data
    })
}

export const updateRole = (data) => {
    return service({
        url:"/api/v1/system/role/" + data.id,
        method: "PUT",
        data
    })
}

export const getroleMenu = (data) => {
    return service({
        url:"/api/v1/system/role/menupermission/" + data.id,
        method:"GET"
    })
}