<template>
  <div class="todos">
    <div class="todo" v-for="todo in allTodos" v-bind:key="todo.id">
      <h3>{{todo.title}}</h3>
      <span id="delete" @click="handleDelete(todo.id)">❌</span>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "Todos",
  methods: {
    ...mapActions(["getTodos", "deleteTodo"]),
    handleDelete(id) {
      this.deleteTodo(id);
    }
  },
  computed: mapGetters(["allTodos"]),
  created() {
    this.getTodos();
  }
};
</script>

<style scoped>
.todos {
  width: 80%;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(17rem, 1fr));
  grid-gap: 25px;
  grid-template-rows: auto;
  margin: 0 auto;
  justify-items: center;
}
.todo {
  color: #fff;
  text-align: center;
  background: rgba(41, 121, 255, 1);
  border: 2px solid rgba(41, 141, 255, 3);
  border-radius: 1rem;
  box-shadow: 1px 2px 1px #333;
  padding: 1rem;
  display: flex;
  width: 12em;
  justify-content: space-between;
  min-height: 3rem;
  display: flex;
  align-items: center;
}
#delete {
  float: right;
  cursor: pointer;
}

@media (max-width: 1555px) {
  .todos {
    /* display: grid; */
    grid-template-columns: repeat(auto-fit, minmax(13rem, 1fr));
    grid-gap: 25px;
    grid-template-rows: auto;
    margin: 0 auto;
    justify-items: center;
  }
  .todo {
    width: 80%;
    margin: 0.5rem;
  }
}

@media (max-width: 780px) {
  .todos {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
  .todo {
    width: 80%;
    margin: 0.5rem;
  }
}
</style>