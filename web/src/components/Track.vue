<script setup>

    defineProps({
        name: String,
        artist: String,
        imgurl: String
        // colors: Object
    })

</script>

<template>

    <div>
        <img class="song-element album-art" :src="imgurl">
        <div class="song-element song-details">
            <div class="name">{{ name }}</div>
            <div class="artist">{{ artist }}</div>
        </div>
        <a class="song-element colors-container" href="https://coolors.co/">
            <div 
                class="color"
                v-for="color in colors"
                :style="`background-color: #${color}`"
            ></div>
        </a>
    </div>

</template>

<style scoped>

    .song-element {
        height: 60px;
        display: inline-block;
        vertical-align: top;
    }

    .album-art {
        width: 60px;
    }

    .song-details {
        width: 350px;
        margin-left: 12px;
    }

    .name {
        font-size: 20px;
        font-weight: 500;
        color: var(--text);
    }

    .artist {
        font-size: 15px;
        font-weight: 500;
        color: var(--primary-accent);
    }


    .colors-container {
        width: 100px;
        float: right;
    }

    .color {
        width: 20px;
        height: 30px;
        float: right;
        vertical-align: top;
    }

</style>