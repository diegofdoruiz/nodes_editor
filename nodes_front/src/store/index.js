import Vue from 'vue';
import Vuex from 'vuex';
import article from'./modules/article';
import authentication from'./modules/authentication';
import nodes from'./modules/nodes';

Vue.use(Vuex);

const store = new Vuex.Store({
  modules: {
    article,
    authentication,
    nodes
  }
});

export default store;