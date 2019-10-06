<template>
  <div class="todos">
    <div class="todo" v-for="todo in allTodos" v-bind:key="todo.id">
      <h3>{{todo.title}}</h3>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
export default {
  name: "Todos",
  computed: mapGetters(["allTodos"])
  // computed: mapGetters(["allTodos"])
};
</script>

<style scoped>
.todos {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(2rem, 1fr));
  grid-gap: 10px;
  grid-template-rows: auto;
}
.todo {
  color: #fff;
  text-align: center;
  background: rgba(41, 121, 255, 1);
  border: 2px solid rgba(41, 141, 255, 3);
  border-radius: 1rem;
  box-shadow: 1px 2px 1px #333;
}
</style>