import axios from 'axios'
import store from '@/store'

import { Message, MessageBox } from "element-ui"
const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_API,
    timeout: 5000
})

//http request拦截器
service.interceptors.request.use(
    config => {
        const token = store.getters["user/token"]
        const userinfo = store.getters["user/userinfo"]
        config.headers = {
            'Content-Type': "application/json",
            'x-token': token,
            'x-user-id': userinfo.id
        }
        return config
    },
    error => {
        Message({
            type: 'error',
            message: '请求拦截器错误',
            showClose: true
        })
    }
)

// http response 拦截器
service.interceptors.response.use(
    response => {
        const res = response.data
        //刷新token
        if (res.code === 0) {
            return res
        }else {
            MessageBox.confirm('You have been logged out, you can cancel to stay on this page, or log in again', 'Confirm logout', {
                confirmButtonText: 'Re-Login',
                cancelButtonText: 'Cancel',
                type: 'warning'
            })
            .then(() => {
                store.dispatch('user/resetToken').then(() => {
                    location.reload()
                })
            })
        }
    }
)