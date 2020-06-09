<template>
    <div id="app">
        <router-view />
        <Nav />
    </div>
</template>

<script>
import Nav from './components/Layout/Nav'
export default {
    components: {
        Nav,
    },
    data() {
        return {}
    },
}
</script>

<style>
* {
    box-sizing: border-box;
    outline: none;
}
html {
    font-family: 'Quicksand';
}

body {
    margin: 0;
    padding: 0;
}
</style>
