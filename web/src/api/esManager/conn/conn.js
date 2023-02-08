import { service } from '@/utils/request.js'

export const esConnList = (data) => {
    return service({
        url: "/api/v1/esmanager/conn/paging",
        method: "POST",
        data
    })
}

export const esPing = (data) => {
    return service({
        url: "/api/v1/esmanager/conn/ping/" + data.id,
        method:"GET"
    })
}

export const deleteESConn = (data) => {
    return service({
        url:"/api/v1/esmanager/conn/" + data.id,
        method: "DELETE"
    })
}

export const esConnDetail = (data) => {
    return service({
        url:"/api/v1/esmanager/conn/" + data.id,
        method: "GET"
    })
}

export const esConnAll = () => {
    return service({
        url: "/api/v1/esmanager/conn/all",
        method: "GET"
    })
} 

export const createESConn = (data) => {
    return service({
        url:"/api/v1/esmanager/conn/create",
        method: "POST",
        data
    })
}

export const updateESConn = (data) => {
    return service({
        url:"/api/v1/esmanager/conn/" + data.id,
        method: "PUT",
        data
    })
}