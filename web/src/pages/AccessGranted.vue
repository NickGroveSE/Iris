<script lang="ts" setup>

import axios from "axios";


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
  <button @click="onClick">Click Me</button>
</template>

<style scoped>

</style>