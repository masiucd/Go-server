<template>
  <div class="addTodo">
    <form @submit="handleSubmit">
      <input type="text" v-model="title" placeholder="add a task" />
      <button type="submit">Submit</button>
    </form>
  </div>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  name: "AddTodo",
  data() {
    return {
      title: ""
    };
  },
  // computed: mapGetters(["addTodo"]),
  created() {},
  methods: {
    ...mapActions(["addTodo"]),
    handleSubmit(e) {
      e.preventDefault();
      this.addTodo(this.title);
      this.title = "";
    }
  }
};
</script>

<style scoped>
.addTodo {
  padding: 2rem 1rem;
  display: flex;
  flex-direction: column;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  width: 100%;
}
form {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 100%;
}
input {
  width: 80%;
  border: 2px solid #333;
  padding: 0.4rem 1rem;
  border-radius: 1rem;
  font-size: 1.2rem;
  transition: all 300ms ease-in-out;
  margin: 0.5rem 0;
}

input:focus {
  transform: scale(1.02);
  border: 2px solid rgba(41, 121, 255, 1);
}

button {
  padding: 0.4rem 0.8rem;
  border: 2px solid #fff;
  background: transparent;
  color: #fff;
  width: 14rem;
  margin: 0.5rem 0;
  border-radius: 1rem;
  transition: all 300ms ease-in-out;
  cursor: pointer;
  font-size: 1rem;
  position: relative;
}

button:hover {
  background: rgba(41, 121, 255, 1);
  color: #333;
}
button:active {
  top: 6px;
}
</style>