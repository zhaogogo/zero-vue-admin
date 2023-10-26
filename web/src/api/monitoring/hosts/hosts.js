import { service } from "@/utils/request.js"

export const hosts = (data) => {
    return service({
        url: "/api/v1/monitoring/hosts/paging",
        method:"POST",
        data
    })
}

export const hostSlienceRule = (data) => {
    return service({
        url: "/api/v1/monitoring/hosts/slience/" + data,
        method:"POST"
    })
}