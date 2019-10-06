/* eslint-disable no-console */

import axios from 'axios';


const api = axios.create({
  baseURL: 'https://jsonplaceholder.typicode.com',
  timeout: 5000
})

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


const actions = {
  async getTodos({commit}){
    const res = await api.get('/todos');

    // sending the action to aut mutation
    commit('setTodos', res.data)
  },
  async addTodo({commit},title){
    const res = await api.post(`/todos`, {title, completed:false })

    commit('newTodo', res.data);
  },

  async deleteTodo({commit},id) {
    await api.delete(`/todos/${id}`)
    commit('deleteTodo',id);
  }
};

const mutations = {
  setTodos: (state,todos) => (state.todos = todos),
  newTodo: (state,todo) => state.todos.unshift(todo),
  deleteTodo: (state,id) => state.todos = state.todos.filter(todo => todo.id !== id),
};


export default {
  state,
  getters,
  actions,
  mutations
};