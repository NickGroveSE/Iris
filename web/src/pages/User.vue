<script lang="ts" setup>

import axios from "axios";
import { ref } from "vue";
import { useRoute } from "vue-router"

const user = ref(useRoute().params)
const code = ref(useRoute().query.code)
const state = ref(useRoute().query.state)

console.log(code)
console.log(state)


const onClick = async () => {
  const response = await axios
    .get(
      "http://localhost:5000/music"
    )
    .then((res) => res.data)
    .catch((err) => console.error(err));

    console.log(response);
};

</script>

<template>
  <h1>{{ user.username }}</h1>
  <button @click="onClick">Click Me</button>

</template>

<style scoped>

</style>