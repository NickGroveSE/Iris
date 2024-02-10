<script setup>

    // defineProps({
    //     tracks: Object
    // })

</script>

<template>

    <div id="iris" :style="`background-color: #${tracks[0].colors[0]}`">
        <!-- <div v-for="(track, index) in tracks" class="song">
            <div 
                v-for="color in track.colors"
                class="glimpse"
                :style="`background: radial-gradient(circle at center, #${color}B3 0, #${color}00 70%) no-repeat; top: calc(${Math.random() * 100}% - 150px / 2); left: calc(${Math.random() * 100}% - 150px / 2); z-index: ${Math.random() * (10/(index + 1))};`"
            ></div>
        </div> -->
    </div>

</template>

<style scoped>

    #iris{
        width: 300px;
        height: 300px;
        margin: 50px auto 0 auto;
        background-color: white;
        filter: drop-shadow(5px 5px var(--primary-accent)) drop-shadow(5px 5px var(--secondary-accent));
        overflow: hidden;
    }

    /* .glimpse {
        position: absolute;
        mix-blend-mode: hard-light;
        
        width: 150px;
        height: 150px;

        opacity: 1;
    } */

</style>
