<template>
    <div id="app">
        <router-view />
        <Nav />
        <template v-if="$store.getters.isPopUpShown">
            <Backdrop />
            <PopUp />
        </template>
    </div>
</template>

<script>
import Nav from './components/Layout/Nav'
import Backdrop from './components/UI/Backdrop'
import PopUp from './components/UI/PopUp'
export default {
    components: {
        Nav,
        Backdrop,
        PopUp,
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
