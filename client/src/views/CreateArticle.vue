<template>
    <div>
        <h1>Create Article</h1>
        <form @submit.prevent="submit()">
            <input type="text" placeholder="Title" v-model="title" />
            <input type="text" placeholder="Image URL" v-model="image" />
            <textarea
                placeholder="Article's content"
                rows="6"
                v-model="content"
            />
            <input type="text" placeholder="Event Name" v-model="eventName" />
            <input type="date" v-model="date" />
            <input type="number" placeholder="Lat" step="0.01" v-model="lat" />
            <input
                type="number"
                placeholder="Long"
                step="0.01"
                v-model="long"
            />
            <button type="submit">Submit Article</button>
        </form>
    </div>
</template>

<script>
export default {
    data() {
        return {
            title: null,
            image: null,
            content: null,
            eventName: null,
            date: null,
            lat: null,
            long: null,
        }
    },
    methods: {
        async submit() {
            alert('hello')
            let res = await fetch(`${process.env.VUE_APP_API_URL}/locations`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    lat: this.lat,
                    long: this.long,
                }),
            })
            const location = await res.json()
            console.log(location)
            res = await fetch(`${process.env.VUE_APP_API_URL}/events`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    name: this.eventName,
                    date: this.date,
                    locationId: location.id,
                }),
            })
            const event = await res.json()
            console.log(event)
            console.log(
                this.title,
                +event.id,
                this.image,
                this.content,
                this.$store.getters.user.id
            )
            res = await fetch(`${process.env.VUE_APP_API_URL}/articles`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    title: this.title,
                    imageUrl: this.image,
                    text: this.content,
                    eventId: +event.id,
                    authorId: this.$store.getters.user.id,
                }),
            })
            const article = await res.json()
            console.log(article)
        },
    },
}
</script>

<style scoped>
div {
    display: flex;
    max-width: 700px;
    margin: 7rem auto 0;
    flex-direction: column;
}
form {
    display: flex;
    flex-direction: column;
    /* margin: auto; */
}
input {
    padding: 0.5rem;
    border-radius: 4px;
    border: 1px solid black;
    margin: 0.5rem 0;
    font-size: 1rem;
}
textarea {
    margin: 0.5rem 0;
    border-color: black;
    color: black;
    padding: 0.5rem;
    font-size: 1.1rem;
    border-radius: 4px;
}
button {
    background-color: transparent;
    border: 2px solid black;
    width: fit-content;
    font-size: 1rem;
    border-radius: 4px;
    padding: 0.5rem 2rem;
    font-weight: bold;
}
</style>
