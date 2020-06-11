<template>
    <div class="root">
        <div class="body-cmp">
            <div class="image">
                <img
                    src="https://cdn.spacetelescope.org/archives/images/wallpaper2/heic2007a.jpg"
                    alt=""
                />
                <button class="like" @click="like()">
                    <i class="material-icons" v-if="!liked">favorite_border</i>
                    <i class="material-icons fill" v-else>favorite</i>
                </button>
            </div>
            <div class="content">
                <div>
                    <h1 class="title">Title for today is this title</h1>
                    <p class="date">Date: 22 june 2020</p>
                    <p class="author">Author: Jhon Doe</p>
                </div>
                <p class="information">
                    Lorem Ipsum is simply dummy text of the printing and
                    typesetting industry. Lorem Ipsum has been the industry's
                    standard dummy text ever since the 1500s, when an unknown
                    printer took a galley of type and scrambled it to make a
                    type specimen book. It has survived not only five centuries,
                    but also the leap into electronic typesetting, remaining
                    essentially unchanged. It was popularised in the 1960s with
                    the release of Letraset sheets containing Lorem Ipsum
                    passages, and more recently with desktop publishing software
                    like Aldus PageMaker including versions of Lorem Ipsum.
                </p>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            liked: false,
        }
    },
    methods: {
        like() {
            if (this.$store.getters.user) {
                this.liked = !this.liked
            } else {
                this.$store.commit('modifyPopUp', true)
            }
        },
    },
    async created() {
        const res = await fetch(`${process.env.VUE_APP_API_URL}/apod`)
        const json = await res.json()
        console.log(json)
    },
}
</script>

<style scoped>
.root {
    padding: 0.5rem;
}
.body-cmp {
    margin: 6rem auto 2rem;
    display: flex;
    flex-direction: column;
    width: fit-content;
    border-radius: 4px;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.4), 0 6px 20px 0 rgba(0, 0, 0, 0.4);
}
.image {
    position: relative;
    margin: 0;
}
img {
    display: block;
    max-width: 800px;
    border-top-left-radius: 4px;
    border-top-right-radius: 4px;
    width: 100%;
}
.like {
    color: rgba(255, 255, 255, 0.377);
    border: none;
    background-color: transparent;
    position: absolute;
    transition: 0.3s;
    bottom: 0.75rem;
    right: 0.75rem;
}
.like i {
    font-size: 2rem;
}
.fill {
    color: tomato;
}
.like:hover {
    transform: scale(1.1);
    color: white;
}
p {
    width: 100%;
}
button {
    cursor: pointer;
}
.content {
    display: flex;
    max-width: 800px;
    justify-content: space-between;
}
.content div {
    background-color: #5e5e5e;
    padding: 1rem;
    color: white;
    font-weight: bold;
    width: fit-content;
    min-width: 30%;
}
.date {
    margin-bottom: 0;
}
.information {
    text-align: justify;
    padding: 1rem;
    color: #383838;
}
</style>
