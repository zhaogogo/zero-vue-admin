import { service } from '@/utils/request.js'

export const connectList = (data) => {
    return service({
        url: "/api/v1/monitoring/store/connect/paging",
        method: "POST",
        data
    })
}

export const connectCreate = (data) => {
    return service({
        url: "/api/v1/monitoring/store/connect/create",
        method: "POST",
        data
    })
}

export const connectDel = (data) => {
    return service({
        url: "/api/v1/monitoring/store/connect/" + data.id,
        method: "DELETE"
    })
}

export const connectDetail = (data) => {
    return service({
        url: "/api/v1/monitoring/store/connect/" + data.id,
        method: "GET"
    })
}

export const connectUpdate = (data) => {
    return service({
        url: "/api/v1/monitoring/store/connect/" + data.id,
        method: "PUT",
        data
    })
}