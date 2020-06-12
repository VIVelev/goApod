<template>
    <span>
        <Loader v-if="!article" />
        <div class="body-cmp" v-else>
            <div class="image">
                <transition>
                    <div id="map" :class="hideMap ? 'show' : 'hide'"></div>
                </transition>
                <img :src="article.imageUrl" :alt="article.title" />
                <button
                    class="like"
                    @click="like()"
                    :style="!liked || 'opacity: 1'"
                >
                    <i class="material-icons" v-if="!liked">favorite_border</i>
                    <i class="material-icons fill" v-else>favorite</i>
                </button>
                <button
                    class="map"
                    :style="!hideMap || 'color:rgba(0, 0, 0, 0.9)'"
                    @click="toggleMap()"
                >
                    <i class="material-icons">my_location</i>
                </button>
            </div>
            <div class="content">
                <div>
                    <h1 class="title">{{ article.title }}</h1>
                    <p class="date">Date: {{ article.date }}</p>
                    <p class="author">Author: {{ article.authorName }}</p>
                    <p class="event">Event: {{ article.eventName }}</p>
                    {{ article.likes }}
                </div>
                <p class="information">
                    {{ article.text }}
                </p>
            </div>
        </div>
    </span>
</template>

<script>
import 'ol/ol.css'
import { Map, View } from 'ol'
import TileLayer from 'ol/layer/Tile'
import OSM from 'ol/source/OSM'
import Loader from '../UI/Loader'
import Sorry from '../UI/Sorry'
export default {
    components: {
        Loader,
        Sorry,
    },
    props: ['article'],
    data() {
        return {
            liked: false,
            map: null,
            hideMap: false,
        }
    },
    methods: {
        async like() {
            if (!this.$store.getters.user) {
                this.$store.commit('modifyPopUp', true)
                return
            }
            if (!this.liked) {
                this.liked = !this.liked
                const res = await fetch(
                    `${process.env.VUE_APP_API_URL}/likes`,
                    {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json',
                        },
                        body: JSON.stringify({
                            authorId: this.$store.getters.user.id,
                            articleId: this.article.id,
                        }),
                    }
                )
                const json = await res.json()
                console.log(json)
            } else {
                this.liked = !this.liked
                const res = await fetch(
                    `${process.env.VUE_APP_API_URL}/likes/${this.$store.getters.user.id}/${this.article.id}`,
                    {
                        method: 'DELETE',
                    }
                )
                const json = await res.json()
                console.log(json)
            }
        },
        toggleMap() {
            if (!this.map) {
                setTimeout(() => {
                    this.map = new Map({
                        target: 'map',
                        layers: [
                            new TileLayer({
                                source: new OSM(),
                            }),
                        ],
                        view: new View({
                            center: [this.article.locLat, this.article.locLong],
                            zoom: 5,
                        }),
                    })
                }, 1000)
                console.log(this.article.locLat, this.article.locLong)
            }
            this.hideMap = !this.hideMap
        },
    },
}
</script>

<style scoped>
#map {
    position: absolute;
    background-color: white;
    z-index: 20;
    transition: 0.5s;
}
.show {
    transition: 1s;
    width: 100%;
    height: 100%;
}
.hide {
    transition: 1s;
    width: 0;
    height: 0;
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
    color: white;
    opacity: 0.5;
    border: none;
    background-color: transparent;
    position: absolute;
    transition: 0.3s;
    bottom: 0.75rem;
    right: 0.75rem;
}
.map {
    color: white;
    opacity: 0.5;
    border: none;
    background-color: transparent;
    position: absolute;
    transition: 0.3s;
    top: 0.75rem;
    right: 0.5rem;
    z-index: 20;
}
.map i {
    font-size: 1.5rem;
}
.like i {
    font-size: 2rem;
}
.fill {
    color: tomato;
}
.map:hover,
.like:hover {
    transform: scale(1.1);
    opacity: 1;
}
p {
    width: 100%;
}
h1 {
    margin: 0.5rem 0;
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
    background-color: #1b1b1b;
    padding: 1rem 2rem 1rem;
    color: white;
    font-weight: bold;
    width: fit-content;
    min-width: 30%;
}
.date {
    margin-bottom: 0;
}
.author {
    margin: 0.25rem 0;
}
.event {
    margin: 0 0 0.25rem;
}
.information {
    text-align: justify;
    padding: 1rem;
    color: #1b1b1b;
}
</style>
