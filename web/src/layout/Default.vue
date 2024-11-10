<script setup>

</script>

<template>

    <header>
        <div class="header-element" id="app-name">
            Iris
        </div>
        <a class="header-element" id="app-logo"></a>
    </header>

    <main id="main">
        <slot/>
    </main>

    <!-- <footer>
        <a class="footer-element" href="https://nickgrove.netlify.app/">My Portfolio Site</a>
        <a class="footer-element" href="https://github.com/NickGroveSE">Github</a>
        <a class="footer-element" >Buy Me a Coffee :)</a>
    </footer> -->

</template>

<style scoped>

    /* Header Styling */
    header {
        height: 60px;
    }

    .header-element {
        vertical-align: top;
        margin: 10px 10px 10px 0;
        line-height: 40px;
        display: inline-block;
    }

    #app-name {
        font-size: 40px;
        font-style: italic;
        font-weight: 700;
    }

    #app-logo {
        width: 1000px;
        height: 1000px;
        background: url('../../public/iris.svg') no-repeat;
        background-size: 72px 48px;
    }

    /* Main Styling */
    #main {
        width: 100%;
        height: calc(100vh - 120px);
    }

    /* Footer Styling */
    /* footer {
        height: 60px;
        width: 100%;
        margin: auto 0;
        text-align: center;
    }

    .footer-element {
        vertical-align: top;
        margin: 10px 10px 10px 10px;
        line-height: 40px;
        display: inline-block;
        color: var(--text);
        
    } */

</style>