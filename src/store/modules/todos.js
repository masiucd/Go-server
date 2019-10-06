
import axios from 'axios';
// https://jsonplaceholder.typicode.com/todos

const state = {
  todos: [
    {
      id: 1,
      title: 'apa apa apa'
    },
    {
      id: 2,
      title: 'banan banan'
    },
    {
      id: 3,
      title: 'kiwi kiwi kiwi'
    },
  ]
};

const getters = {
  allTodos: state => state.todos
};


const actions = {};

const mutations = {};


export default {
  state,
  getters,
  actions,
  mutations
};