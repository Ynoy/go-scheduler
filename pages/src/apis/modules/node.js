import {Vue} from "vue/types/vue";

export default {
    fetch(params) {
        return Vue.axios.get('/apis/node', {
            params: params
        })
    },
    create(params) {
        return Vue.axios.post('/apis/node', params)
    },
    update(id, params) {
        return Vue.axios.put(`/apis/node/${id}`, params)
    },
    delete(id) {
        return Vue.axios.delete(`/apis/node/${id}`)
    },
    fetchPipelines(params) {
        return Vue.axios.get('/api/node/pipelines', {
            params: params
        })
    },
    bindPipeline(params) {
        return Vue.axios.post('/api/node/pipeline', params);
    },
    unbindPipeline(id) {
        return Vue.axios.delete(`/api/node/pipeline/${id}`);
    }
}
