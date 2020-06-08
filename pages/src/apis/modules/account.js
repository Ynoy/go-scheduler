import Vue from "vue";

// 用户相关接口请求封装

export default {
    signIn(params) {
        return Vue.axios.post('/apis/auth/signIn', params)
    },
    profile() {
        return Vue.axios.get('/apis/account/profile')
    }
}