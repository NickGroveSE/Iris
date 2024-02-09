<script lang="ts" setup>

import axios from "axios";
import { onMounted } from "vue";

onMounted(async () => {
  const response = await axios
    .get(
      "http://localhost:5000/redirect"
    )
    .then((res) => res.data)
    .catch((err) => console.error(err))

    window.location.href = response.response;
});

</script>

<template>
</template>

<style scoped>

</style>