import axios from 'axios';
import router from '@/router'

const state = {
    token: '',
    errors: [],
    username_error: false,
    password_error: false
}

const mutations = {
    UPDATE_TOKEN(state, payload) {
        state.token = payload;
    },
    UPDATE_ERRORS(state, payload) {
        state.errors = payload;
    },
    UPDATE_USERNAME_ERROR(state, payload) {
        state.username_error = payload;
    },
    UPDATE_PASSWORD_ERROR(state, payload) {
        state.password_error = payload;
    }
}

const actions = {
    ResetField({ commit }, payload){
        if(payload == "username"){
            commit('UPDATE_USERNAME_ERROR', false);
        }else{
            commit('UPDATE_PASSWORD_ERROR', false);
        }
    },
    signIn({ commit }, payload){
        commit('UPDATE_TOKEN', "");
        commit('UPDATE_USERNAME_ERROR', false);
        commit('UPDATE_PASSWORD_ERROR', false);
        axios({
            method: 'POST',
            url: 'http://localhost:3333/user/login',
            data: {
                username: payload.username,
                password: payload.password,
            }
        }).then(response => {
            if(response.data.token){
                let user = {}
                user["username"] = response.data.username;
                user["token"] = response.data.token;
                localStorage.setItem('user', JSON.stringify(user));
                commit('UPDATE_TOKEN', response.data.token);
                router.push('dashboard');
            }
            if(response.data.errors){
                response.data.errors.map(function(key) {
                    if(key.field == "username"){
                        commit('UPDATE_USERNAME_ERROR', true);
                    }
                    if(key.field == "password"){
                        commit('UPDATE_PASSWORD_ERROR', true);
                    }
                });
                commit('UPDATE_ERRORS', response.data.errors);
            }
        })        
    }
}

const getters = {
    token: state => state.token,
    errors: state => state.errors,
    username_error: state => state.username_error,
    password_error: state => state.password_error
}

const authenticationModule = {
    state,
    mutations,
    actions,
    getters
}

export default authenticationModule;