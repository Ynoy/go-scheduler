import Vue from 'vue'

export default {
    fetchNodes() {
        return Vue.axios.get('/apis/dashboard/nodes')
    },
    fetchPipelines() {
        return Vue.axios.get('/apis/dashboard/pipelines')
    },
    fetchFailures() {
        return Vue.axios.get('/apis/dashboard/failures')
    }
}