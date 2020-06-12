<template>
    <div class="root">
        <Loader v-if="!article" />
        <Sorry v-if="error">No article for today</Sorry>
        <Article v-if="article && !error" />
    </div>
</template>

<script>
import Article from '../components/Article/Article'
import Sorry from '../components/UI/Sorry'
import Loader from '../components/UI/Loader'
export default {
    components: {
        Article,
        Loader,
        Sorry,
    },
    data() {
        return {
            article: null,
            error: false,
        }
    },
    async created() {
        const res = await fetch(`${process.env.VUE_APP_API_URL}/apod`)
        const json = await res.json()
        if (json.message) {
            setTimeout(() => {
                this.error = true
                this.article = json
            }, 500)
        } else {
            this.article = json
        }

        console.log(json)
    },
}
</script>

<style scoped>
.root {
    padding: 0.5rem;
}
</style>
