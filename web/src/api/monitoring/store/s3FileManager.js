import { service } from '@/utils/request.js'

export const fileEnvSelect = () => {
    return service({
        url: "/api/v1/monitoring/store/file/envselect",
        method: "GET"
    })
}

export const fileList = (data) => {
    return service({
        url: "/api/v1/monitoring/store/file/",
        method: "POST",
        data
    })
}