<template>
    <div id="app" :style="$store.getters.isPopUpShown || 'overflow: hidden'">
        <router-view />
        <Nav />

        <transition
            enter-active-class="animated slideInUp"
            leave-active-class="animated slideOutDown"
        >
            <Backdrop v-if="$store.getters.isPopUpShown" />
        </transition>
        <transition
            enter-active-class="animated fadeInDownBig"
            leave-active-class="animated fadeOutUpBig"
        >
            <PopUp v-if="$store.getters.isPopUpShown" />
        </transition>
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

::-webkit-scrollbar {
    width: 3px;
}

::-webkit-scrollbar-track {
    background: #f1f1f1;
}

::-webkit-scrollbar-thumb {
    background: #888;
}
</style>
