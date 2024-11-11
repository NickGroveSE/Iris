<script setup>

    import Track from './Track.vue'

    defineProps({
        tracks: Object
    })

</script>

<template>

    <div id="track-wrapper">
        <Track
            class="track"
            v-for="track in tracks"
            :name="track.name"
            :artist="track.artist"
            :imgurl="track.imgurl"
            :glimpses="track.glimpses"
        />
    </div>

</template>

<style scoped>

    .track {
        padding: 5px 0 5px 0;
        margin: 0 auto;
        width: 90%;
        border-bottom: 1px solid var(--primary-accent);
    }

    .track:last-child {
        border-bottom: 0px solid white;
    }

</style>