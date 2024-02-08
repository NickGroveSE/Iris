<script lang="ts" setup>

import axios from "axios";
// import { onMounted } from "vue";
import { useRoute } from "vue-router"

const { username }=useRoute().params


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
  <h1>{{ username }}</h1>
  <button @click="onClick">Click Me</button>
</template>

<style scoped>

</style>