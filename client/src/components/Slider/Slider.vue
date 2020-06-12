<template>
    <transition-group tag="div" class="container">
        <Slides
            v-for="(img, index) of $store.getters.images"
            :key="index"
            :img="img"
            :index="index"
        />
    </transition-group>
</template>

<script>
import Slides from './Slides'
export default {
    data() {
        return {}
    },
    components: {
        Slides,
    },
}
</script>

<style scoped>
.container {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    margin: auto;
    max-width: 1500px;
    width: 90%;
    overflow-x: hidden;
    max-height: 600px;
}
</style>
