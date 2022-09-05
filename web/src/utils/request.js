import axios from 'axios'
import store from '@/store'

import { Message, MessageBox } from "element-ui"
export const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_API,
    timeout: 5000
})

//http request拦截器
service.interceptors.request.use(
    config => {
        const token = store.getters["user/token"]
        const userinfo = store.getters["user/userInfo"]
        config.headers = {
            'Content-Type': "application/json",
            'Authorization': "Bearer " + token,
            'x-user-id': userinfo.id
        }
        return config
    },
    error => { // eslint-disable-line no-unused-vars
        Message({
            type: 'error',
            message: '请求拦截器错误',
            showClose: true
        })
        console.log('请求拦截器错误 ==> ',error)
    }
)

// http response 拦截器
service.interceptors.response.use(
    response => {
        // console.log("@@@", response)
        //刷新token
        // if (response.headers["new-token"]) {
        //     store.commit("user/SETTOKEN",response.headers["new-token"])
        // }
        if (response.status === 200) {
            return response
        } else {
            MessageBox.confirm('You have been logged out, you can cancel to stay on this page, or log in again', 'Confirm logout', {
                confirmButtonText: 'Re-Login',
                cancelButtonText: 'Cancel',
                type: 'warning'
            })
            .then(() => {
                store.commit('user/LOGOUT')
            })
            return Promise.reject(new Error("错误XXXXX ==> ",statusText))
        }
    },
    error => { // eslint-disable-line no-unused-vars
        console.log("响应拦截器错误 ==> ",error)
        Message({
            showClose: true,
            message: error.response.data,
            type: 'error'
        })
        return error
    }
)