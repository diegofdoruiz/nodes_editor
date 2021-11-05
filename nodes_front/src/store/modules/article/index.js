import axios from 'axios';
const state = {
    articleItems: []
}

const mutations = {
    UPDATE_ARTICLE_ITEMS(state, payload) {
        state.articleItems = payload;
    }
}

const actions = {
    getArticleItems({ commit }) {
        axios.get(`http://localhost:3333/programs`).then((response) => {
            commit('UPDATE_ARTICLE_ITEMS', response.data)
        });
    }
}

const getters = {
    articleItems: state => state.articleItems,
    articleItemById: (state) => (id) => {
        return state.articleItems.find(articleItem => articleItem.id === id)
    }
}

const articleModule = {
    state,
    mutations,
    actions,
    getters
}


export default articleModule;