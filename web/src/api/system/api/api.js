import { service } from '@/utils/request.js'

export const allAPI = () => {
    return service({
        url: "/api/v1/system/api/all",
        method: "GET"
    })
}

export const pagingAPI = (data) => {
    return service({
        url: "/api/v1/system/api/paging",
        method: "POST",
        data
    })
}

export const detailAPI = (data) => {
    return service({
        url: "/api/v1/system/api/" + data.id,
        method: "GET"
    })
}

export const deleteAPI = (data) => {
    return service({
        url: "/api/v1/system/api/" + data.id,
        method: "DELETE",
        data
    })
}

export const deleteMulitAPI = (data) => {
    return service({
        url: "/api/v1/system/api/multiple",
        method: "DELETE",
        data
    })
}

export const createApi = (data) => {
    return service({
        url:"/api/v1/system/api/create",
        method: "POST",
        data
    })
}

export const editApi = (data) => {
    return service({
        url:"/api/v1/system/api/" + data.id,
        method: "PUT",
        data
    })
}

