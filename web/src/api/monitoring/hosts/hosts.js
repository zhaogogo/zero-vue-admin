import { service } from "@/utils/request.js"

export const hosts = (data) => {
    return service({
        url: "/api/v1/monitoring/hosts/paging",
        method:"POST",
        data
    })
}

export const createHost = (data) => {
    return service({
        url: "/api/v1/monitoring/hosts",
        method:"POST",
        data
    })
}

export const updateHost = (data) => {
    return service({
        url: "/api/v1/monitoring/hosts",
        method:"PUT",
        data
    })
}

export const hostSlienceRule = (data) => {
    return service({
        url: "/api/v1/monitoring/hosts/slience/" + data,
        method:"POST"
    })
}

export const putHostSlienceRule = (data) => {
    return service({
        url: "/api/v1/monitoring/hosts/slience/" + data.host,
        method:"PUT",
        data
    })
}

export const getAllSlience = () => {
    return service({
        url: "/api/v1/monitoring/hosts/slience/all",
        method: "GET"
    })
}

export const handlerHostsSlience = (data) => {
    return service({
        url: "/api/v1/monitoring/hosts/slience",
        method: "POST",
        data
    })
}