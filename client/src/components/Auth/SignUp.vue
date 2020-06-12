<template>
    <form @submit.prevent="signUp">
        <h1>Sign Up</h1>
        <input type="text" placeholder="Name" v-model="name" />
        <button type="submit">Sign Up</button>
        <span>
            <p>Already have an account?</p>
            <button type="button" @click="$emit('changeComponent', 'SignIn')">
                Sign In.
            </button>
        </span>
    </form>
</template>

<script>
export default {
    data() {
        return {
            name: null,
        }
    },
    methods: {
        async signUp() {
            // this.$store.commit('changeUser', this.name)
            if (!this.name) {
                return
            }
            const res = await fetch(`${process.env.VUE_APP_API_URL}/authors`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ name: this.name }),
            })
            const json = await res.json()
            if (json.message) {
                console.log(json)
                alert(json.message)
            } else {
                this.$store.commit('changeUser', json)
            }
            this.$store.commit('modifyPopUp', false)
        },
    },
}
</script>

<style scoped>
form {
    display: flex;
    flex-direction: column;
    padding: 4rem 0 2rem;
    color: #1b1b1b;
    width: 100%;
}
h1 {
    margin: 0 auto;
    color: #1b1b1b;
    width: 80%;
}
input {
    border: none;
    border-bottom: 2px solid #777;
    padding: 0.5rem;
    margin: 2rem auto;
    font-size: 1rem;
    width: 80%;
}
::placeholder {
    color: #1b1b1b;
}
input:focus {
    border-color: #1b1b1b;
}

button {
    border: none;
    background-color: transparent;
    font-size: 1.3rem;
    cursor: pointer;
    color: #1b1b1b;
}
span {
    display: flex;
    margin: auto;
}
span p {
    font-size: 0.8rem;
}
span button {
    font-size: 0.8rem;
    font-weight: bold;
}
</style>
