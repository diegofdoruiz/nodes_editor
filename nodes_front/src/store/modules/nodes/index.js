import axios from 'axios';

const state = {
    test: false,

}

const mutations = {
    UPDATE_TEST(state, payload) {
        state.test = payload;
    },
}

const actions = {
    async createNode({ commit }, payload){
        commit('UPDATE_TEST', false);
        return axios({
            method: 'POST',
            url: 'http://localhost:3333/nodes/create',
            data: {
                token:"12345",
                node: payload,
            }
        }).then(response => {
            return response;
        });       
    },
    async deleteNode({ commit }, payload){
        commit('UPDATE_TEST', false);
        return axios({
            method: 'DELETE',
            url: 'http://localhost:3333/nodes/'+payload.uid,
            data: {
                token:"12345",
                node: payload,
            }
        }).then(response => {
            return response;
        }) 
    },
    async nodeUpdatePosition({ commit }, payload){
        commit('UPDATE_TEST', false);
        return axios({
            method: 'PUT',
            url: 'http://localhost:3333/nodes/'+payload.uid,
            data: {
                token:"12345",
                pos_x: payload.pos_x,
                pos_y: payload.pos_y
            }
        }).then(response => {
            return response;
        }) 
    },
    async nodeUpdateData({ commit }, payload){
        commit('UPDATE_TEST', false);
        return axios({
            method: 'POST',
            url: 'http://localhost:3333/nodes/data',
            data: {
                token:"12345",
                node_data: payload
            }
        }).then(response => {
            return response;
        }) 
    },
    async createNodeConnections({ commit }, payload){
        commit('UPDATE_TEST', false);
        return axios({
            method: 'POST',
            url: 'http://localhost:3333/nodes/connections/create',
            data: {
                token:"12345",
                output_connection: payload.output_connection,
                input_connection: payload.input_connection
            }
        }).then(response => {
            return response;
        });       
    },
    async deleteConnections({ commit }, payload){
        commit('UPDATE_TEST', false);
        return axios({
            method: 'POST',
            url: 'http://localhost:3333/nodes/connections/delete',
            data: {
                token:"12345",
                output_connection: payload.output_connection, 
                input_connection: payload.input_connection
            }
        }).then(response => {
            return response;
        });
    },
    async searchModule({ commit },payload){
        commit('UPDATE_TEST', false);
        return axios({
            method: 'POST',
            url: 'http://localhost:3333/modules/search',
            data: {
                token:"12345",
                module: payload,
            }
        }).then(response => {
            if(response.data.success){
                return response.data;
            }
        })
    },
    async createModule({ commit }, payload){
        commit('UPDATE_TEST', false);
        return axios({
            method: 'POST',
            url: 'http://localhost:3333/modules/create',
            data: {
                token:"12345",
                module: payload,
            }
        }).then(response => {
            return response;
        })       
    },
    async clearModule({ commit }, payload){
        commit('UPDATE_TEST', false);
        return axios({
            method: 'PUT',
            url: 'http://localhost:3333/modules/'+payload,
            data: {
                token:"12345",
            }
        }).then(response => {
            return response;
        }) 
    },
    async userGetModules({ commit }, payload){
        commit('UPDATE_TEST', false);
        return axios({
            method: 'POST',
            url: 'http://localhost:3333/modules',
            data: {
                token:"12345",
                username: payload,
            }
        }).then(response => {
            return response;
        })       
    },
    async userDeleteModule({ commit }, payload){
        commit('UPDATE_TEST', false);
        return axios({
            method: 'DELETE',
            url: 'http://localhost:3333/modules/'+payload.uid,
            data: {
                token:"12345",
                module: payload,
            }
        }).then(response => {
            return response;
        })       
    }
}

const getters = {
    isset_module: state => state.isset_module,
}

const nodesModule = {
    state,
    mutations,
    actions,
    getters
}

export default nodesModule;