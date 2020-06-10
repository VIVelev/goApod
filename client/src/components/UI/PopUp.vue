<template>
    <div>
        <form>
            <h1>Sign in</h1>
            <input type="text" placeholder="Name" v-model="name" />
            <button @click="signIn">Sign in</button>
        </form>
    </div>
</template>

<script>
export default {
    data() {
        return {
            name: null,
        }
    },
    methods: {
        signIn() {
            this.$store.commit('changeUser', this.name)
            this.$store.commit('modifyPopUp', false)
        },
    },
}
</script>

<style scoped>
div {
    position: fixed;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 90%;
    background: white;
    border-radius: 4px;
    max-width: 500px;
    z-index: 3;
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
