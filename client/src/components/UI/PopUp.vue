<template>
    <div>
        <form @submit.prevent="signIn">
            <h1>Sign in</h1>
            <input type="text" placeholder="Name" v-model="name" />
            <button type="submit">Sign in</button>
            <!-- {{ names }} -->
        </form>
    </div>
</template>

<script>
export default {
    data() {
        return {
            name: null,
            names: null,
        }
    },
    methods: {
        signIn() {
            this.$store.commit('changeUser', this.name)
            this.$store.commit('modifyPopUp', false)
        },
    },
    async created() {
        const res = await fetch('http://localhost:5000/authors')
        const json = await res.json()
        console.log(json)
        this.names = json
    },
}
</script>

<style scoped>
div {
    position: fixed;
    top: 35%;
    left: 0;
    right: 0;
    width: 90%;
    background: white;
    border-radius: 4px;
    max-width: 500px;
    z-index: 3;
    margin: auto;
}
form {
    display: flex;
    flex-direction: column;
    max-width: 500px;
    padding: 4rem 4rem 2rem;
}
h1 {
    margin: 0;
}
input {
    border: none;
    border-bottom: 2px solid #777;
    padding: 0.5rem;
    margin: 2rem 0;
    font-size: 1rem;
}
input:focus {
    border-color: black;
}

button {
    border: none;
    background-color: transparent;
    font-size: 1.3rem;
    cursor: pointer;
}
</style>
