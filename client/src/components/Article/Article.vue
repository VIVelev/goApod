<template>
    <span>
        <div class="body-cmp">
            <div class="image">
                <transition>
                    <div id="map" :class="hideMap ? 'show' : 'hide'"></div>
                </transition>
                <img
                    src="https://cdn.spacetelescope.org/archives/images/wallpaper2/heic2007a.jpg"
                    alt=""
                />
                <button class="like" @click="like()">
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
    </span>
</template>

<script>
import 'ol/ol.css'
import { Map, View } from 'ol'
import TileLayer from 'ol/layer/Tile'
import OSM from 'ol/source/OSM'
export default {
    data() {
        return {
            liked: false,
            map: null,
            hideMap: false,
            article: null,
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
                            center: [
                                this.article?.lat ?? 0,
                                this.article?.long ?? 0,
                            ],
                            zoom: 5,
                        }),
                    })
                }, 1000)
                console.log(this.article?.lat ?? 0, this.article?.long ?? 0)
            }
            this.hideMap = !this.hideMap
        },
    },
    async created() {
        const res = await fetch(
            `${process.env.VUE_APP_API_URL}/articles/${this.$route.params.id}`
        )
        const json = await res.json()
        if (json.message) {
            //handle error
            return
        }
        this.article = json
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
    color: #1b1b1b;
}
</style>
