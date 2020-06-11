<template>
    <div>
        <h1>Posts</h1>
        {{ posts }}
    </div>
</template>

<script>
export default {
    data() {
        return {
            posts: [],
        }
    },
    async created() {
        const res = await fetch(`${process.env.VUE_APP_API_URL}/articles`)
        const json = await res.json()
        this.posts = json
    },
}
</script>

<style scoped>
div {
    margin-top: 5rem;
}
</style>
