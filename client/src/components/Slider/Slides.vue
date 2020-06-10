<template>
    <div
        v-show="index === $store.getters.activeElement || side"
        :class="side ? 'side' : 'center'"
    >
        <img
            :src="img"
            alt="picture"
            @click="$router.push({ path: '/today' })"
        />
        <!-- should go to the article -->
        <template v-if="index === $store.getters.activeElement">
            <button @click="$store.commit('prev')" class="arrow">
                <i class="material-icons">chevron_left</i>
            </button>
            <button class="arrow right" @click="$store.commit('next')">
                <i class="material-icons">chevron_right</i>
            </button>
            <button class="like" @click="like()">
                <i class="material-icons" v-if="!liked">favorite_border</i>
                <i class="material-icons fill" v-else>favorite</i>
            </button>
        </template>
    </div>
</template>

<script>
export default {
    props: ['img', 'index'],
    data() {
        return {
            liked: false,
            moreInfo: false,
        }
    },
    computed: {
        side() {
            return (
                this.index === this.$store.getters.activeElement - 1 ||
                this.index === this.$store.getters.activeElement + 1
            )
        },
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
}
</script>

<style scoped>
div {
    position: relative;
    display: flex;
}
img {
    max-height: 600px;
    display: block;
    border-radius: 4px;
}
.side {
    transform: scale(0.7);
    filter: blur(0.5rem);
    opacity: 0.7;
}
.arrow {
    color: rgba(255, 255, 255, 0.377);
    border: none;
    background-color: transparent;
    position: absolute;
    top: 45%;
    transition: 0.5s;
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
.fill {
    color: tomato;
}
.like:hover {
    transform: scale(1.1);
}
.like i {
    font-size: 3rem;
}
.right {
    right: 0;
}
i {
    font-size: 5rem;
}
button:hover {
    color: white;
    transform: scale(1.5);
}
.center {
    position: absolute;
    z-index: 1;
}
.more-info {
    position: absolute;
    font-size: 1rem;
    color: rgba(255, 255, 255, 0.377);
    bottom: 0.5rem;
    left: 0.75rem;
    padding: 0;
    display: flex;
}
.more-info i {
    font-size: 1.5rem;
    color: rgba(255, 255, 255, 0.795);
}

.more-info i:hover {
    color: white;
}

.more-info button {
    border: none;
    background-color: transparent;
    transform: scale(1);
    padding: 0;
}
p {
    margin: 0;
    margin-bottom: 0.4rem;
    text-align: center;
    color: white;
    cursor: pointer;
}
</style>
