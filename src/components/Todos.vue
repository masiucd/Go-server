<template>
  <div>
    <!--  -->
    <div class="options-todos">
      <div class="option opt-one">Doble click to Complete</div>
      <div class="option opt-two">
        <span>uncomplete</span>
      </div>
      <div class="option opt-three">
        <span>complete</span>
      </div>
    </div>
    <!--  -->
    <div class="todos">
      <div
        @dblclick="handleDblClick(todo)"
        class="todo"
        v-bind:class="{'is-complete':todo.completed}"
        v-for="todo in allTodos"
        v-bind:key="todo.id"
      >
        <h3>{{todo.title}}</h3>
        <span id="delete" @click="handleDelete(todo.id)">❌</span>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "Todos",
  methods: {
    ...mapActions(["getTodos", "deleteTodo", "updateTodo"]),
    handleDelete(id) {
      this.deleteTodo(id);
    },
    handleDblClick(todo) {
      let updateTodo = {
        id: todo.id,
        title: todo.title,
        completed: !todo.completed
      };
      this.updateTodo(updateTodo);
    }
  },
  computed: mapGetters(["allTodos"]),
  created() {
    this.getTodos();
  }
};
</script>

<style scoped>
.options-todos {
  display: flex;
  margin: 1rem auto;
  justify-content: space-between;

  width: 80%;
  padding: 1rem 0;
  align-items: center;
  flex-wrap: wrap;
}

.opt-one {
  font-size: 18px;
  font-weight: 900;
}
.opt-two {
  background: tomato;
  padding: 0.5rem;
  width: 5rem;
  height: 1rem;
  display: flex;
  justify-content: center;
}
.opt-three {
  background: rgba(41, 121, 255, 1);
  padding: 0.5rem;
  width: 5rem;
  height: 1rem;
  display: flex;
  justify-content: center;
}
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
  background: tomato;
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

.is-complete {
  background: rgba(41, 121, 255, 1);
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