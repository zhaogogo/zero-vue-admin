import { service } from '@/utils/request.js'

export const alertRules = (data) => {
    return service({
        url: "/api/v1/prometheus/alertrule/paging",
        method: "POST",
        data
    })
}