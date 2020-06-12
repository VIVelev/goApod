<template>
    <div class="root">
        <Loader v-if="!article" />
        <Sorry v-if="error">No article for today</Sorry>
        <Article v-if="article && !error" :article="article" />
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
        const res = await fetch(
            `${process.env.VUE_APP_API_URL}/articles/${this.$route.params.id}`
        )
        const json = await res.json()
        if (json.message) {
            console.log(json.message)
            //handle error
            return
        }
        this.article = json
        this.article.date = new Date(json.date).toString().slice(3, 15)
    },
}
</script>

<style scoped>
.root {
    padding: 1rem;
}
</style>
