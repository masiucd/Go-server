<template>
  <div class="filter">
    <h4>Filter your todos:</h4>
    <select @change="filterTodos($event)">
      <option value="200">200</option>
      <option value="100">100</option>
      <option value="50">50</option>
      <option value="20">20</option>
      <option value="10">10</option>
      <option value="5">5</option>
    </select>
  </div>
</template>

<script>
import { mapActions } from "vuex";

export default {
  name: "Filter",
  methods: mapActions(["filterTodos"])
};
</script>

<style>
.filter {
  width: 80%;
  margin: 0 auto;
  padding: 0.3rem 0 1rem 0;
}
select {
  margin-top: 20px;
  padding: 6px;
  border: 1px solid rgba(41, 121, 255, 1);
}
</style>