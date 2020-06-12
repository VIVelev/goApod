<template>
    <transition-group tag="div" class="container">
        <Slides
            v-for="(article, index) of $store.getters.articles"
            :key="index"
            :article="article"
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
    created() {
        this.$store.dispatch('getArticles', 2)
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
