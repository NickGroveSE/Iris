<script>

import axios from "axios";

const onClick = async () => {
  const response = await axios
    .post(
      "http://localhost:5000/music"
    )
    .then((res) => res.data)
    .catch((err) => console.error(err));

  
};

</script>

<template>
<button @click="onClick">Click Me</button>
</template>

<style scoped>

</style>
