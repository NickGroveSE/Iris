<script setup>

    import { onMounted, ref, reactive } from 'vue';

    defineProps({
        tracks: Object
    })

    // onMounted(() => {
    //     console.log(tracks[0])
    // })

</script>

<template>

    <div id="iris" :style="`background-color: rgba(${tracks[0].glimpses[0][0]},${tracks[0].glimpses[0][1]},${tracks[0].glimpses[0][2]},1.0)`">
        <div v-for="(track, index) in tracks" class="song">
            <div 
                v-for="(color, rgbIndex) in track.glimpses"
                class="glimpse"
                :style="`background: radial-gradient(circle at center, rgba(${color[0]},${color[1]},${color[2]}, 1) 0, rgba(${color[0]},${color[1]},${color[2]}, 0) 70%) no-repeat; top: calc(${Math.random() * 100}% - 300px / 2); left: calc(${Math.random() * 100}% - 300px / 2);`"
            ></div>
        </div>
        <!-- <svg id="filterSVG">
            <filter id="grainy">
                <feTurbulence
                    type="turbulence"
                    base-frequency="0.65"
                />
            </filter>
            <rect x="0" y="0" width="450" height="450" /> 
        </svg> -->
    </div>

</template>

<style scoped>

    #iris{
        width: 500px;
        height: 500px;
        margin: 100px auto 0 auto;
        background-color: white;
        filter: drop-shadow(5px 5px var(--primary-accent)) drop-shadow(5px 5px var(--secondary-accent));
        overflow: hidden;
        /* filter: url(#grainy); */
        z-index: 0;
    }

    #iris::after {
        content: "";
        position: absolute;
        width: 100%;
        height: 100%;
        backdrop-filter: blur(10px); /* apply the blur */
        pointer-events: none; /* make the overlay click-through */
    }


    .glimpse {
        position: absolute;
        mix-blend-mode: hard-light;
        
        width: 300px;
        height: 300px;

        opacity: 1;
    }

</style>
