<script lang="ts" setup>

import axios from "axios";
import { onMounted, ref, reactive } from 'vue';
import { useRoute } from "vue-router"
import DefaultLayout from '../layout/Default.vue'
// import Iris from '../components/Iris.vue'
import Tracklist from '../components/Tracklist.vue'

const user = ref(useRoute().params)
console.log(user)

let tracklistState = reactive({
  tracks: []
})


onMounted(async () => {
  const response = await axios
    .get(
      "http://localhost:5000/music", {
        params: { timerange: "short_term"},
      }
    )
    .then((res) => res.data)
    .catch((err) => console.error(err));

    tracklistState.tracks = response.data

});

const onClick = async (event) => {

  const response = await axios
    .get(
      "http://localhost:5000/music", {
        params: { timerange: event.target.id},
      }
    )
    .then((res) => res.data)
    .catch((err) => console.error(err));

    tracklistState.tracks = response.data

    console.log(response.data)

}

</script>

<template>
  <DefaultLayout>
    <div class="main-panel" id="left-panel">
      <!-- <Iris :tracks="track"/> -->
      <div id="description">This is your Iris. An iridescent morphing collection of all the colors that make up the album artwork for your favorite music. Change the timeframe to when you want to pull from on the right.</div>
    </div>
    <div class="main-panel" id="right-panel">
      <ul id="timeframes">
        <li class="timeframe"><a @click="onClick($event)" id="short_term">1 Month</a></li>
        <li class="timeframe"><a @click="onClick($event)" id="medium_term">6 Months</a></li>
        <li class="timeframe"><a @click="onClick($event)" id="long_term">1 Year</a></li>
      </ul>
      <Tracklist id="tracklist" :tracks="tracklistState.tracks"/>
    </div>
  </DefaultLayout>
</template>

<style scoped>

  .main-panel {
    height: 100%;
    position: relative;
    display: inline-block;
    vertical-align: top;
  }

  #left-panel {
    width: calc(50% - 100px);
    margin-left: 100px;
  }

  #right-panel {
    width: calc(50% - 100px);
    margin-right: 100px;
  }

  /* Intra-Panel Styling*/
  #description {
    width: 600px;
    height: 80px;
    position: absolute;
    left: 0;
    right: 0;
    bottom: 50px;
    margin: 0 auto 0 auto;
    padding: 5px;
    font-size: 20px;
    text-align: center;
  }

  #timeframes {
    width: 500px;
    margin: 0 auto 20px auto;
    
    list-style: none;
  }

  .timeframe {
    width: 130px;
    line-height: 30px;
    display: inline-block;
    font-weight: 600;
    font-size: 28px;
    text-align: center;
    padding: 10px 10px;
  }

  .timeframe:hover{
    background-color: var(--primary-accent);
    cursor: pointer;
    transition: 0.5s;
  }

  a {
    text-decoration: none;
  }

  #yearly {
    border-bottom: 2px solid var(--primary-accent);
    color: var(--primary-accent);
    /* box-shadow: 0px 2px var(--text-accent); */
  }

  /* #tracklist {
    height: 400px;
    width: 450px;
    position: absolute;
    left: 0;
    right: 0;
    bottom: 40px;
    margin: 0 auto 0 auto;
  } */

</style>