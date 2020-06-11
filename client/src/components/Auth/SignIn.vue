<template>
    <form>
        <h1>Sign In</h1>
        <div class="search">
            <i class="material-icons">search</i>
            <input type="text" placeholder="Name?" v-model="name" />
        </div>
        <section>
            <p
                v-for="user in filteredUsers"
                :key="user.id"
                @click="selected = user"
                :class="user !== selected || 'selected'"
            >
                {{ user.name }}
            </p>
        </section>
        <span>
            <p>Don't have an account?</p>
            <button type="button" @click="$emit('changeComponent', 'SignUp')">
                Sign Up.
            </button>
        </span>
    </form>
</template>

<script>
export default {
    data() {
        return {
            users: [],
            name: '',
            selected: null,
        }
    },
    computed: {
        filteredUsers() {
            return this.users.filter(u => u.name.includes(this.name))
        },
    },
    async created() {
        const res = await fetch(`${process.env.VUE_APP_API_URL}/authors`)
        const json = await res.json()
        console.log(json)
        this.users = json
    },
}
</script>

<style scoped>
h1 {
    margin: 0;
}
.search {
    border: 1px solid #555;
    border-radius: 20px;
    margin: 1rem 0;
    display: flex;
    align-items: center;
}
.search input {
    padding: 0.7rem 0.5rem 0.55rem;
    border: none;
    border-radius: 20px;
    margin: 0;
    color: #555;
}
.search i {
    margin-left: 0.75rem;
    color: #555;
}
input {
    padding: 0.5rem;
    font-size: 1rem;
}
section {
    max-height: 100px;
    overflow-y: scroll;
    margin: 0 1rem;
}
form {
    display: flex;
    flex-direction: column;
    max-width: 500px;
    padding: 4rem 4rem 2rem;
}
form {
    display: flex;
    flex-direction: column;
}
button {
    border: none;
    background-color: transparent;
    font-size: 1.3rem;
    cursor: pointer;
}
p {
    cursor: pointer;
    border-radius: 4px;
    padding: 0.5rem;
    margin: 0 0.25rem;
    border: 1px solid transparent;
}
span {
    display: flex;
    margin: auto;
    padding: 0.5rem;
}
span p {
    font-size: 0.8rem;
    margin: 0;
    padding: 0;
}
span button {
    font-size: 0.8rem;
    font-weight: bold;
    padding: 0;
}

.selected {
    color: teal;
    background-color: rgb(228, 228, 228);
    border-color: teal;
}
</style>
