<template>
    <div class="page-body">
        <Intoduction />
        <div class="cp">
            <!-- <h1>Top rated photos</h1> -->
            <div>
                <img
                    v-for="(img, index) of images"
                    :key="index"
                    :src="img"
                    alt="ddz"
                    :class="{ active: index === 1 }"
                />
            </div>
            <button @click="activeIndex++">Increment</button>
            <button @click="activeIndex--">Decrement</button>
        </div>
    </div>
</template>

<script>
import Parallax from '../components/UI/Parallax'
import Intoduction from '../components/UI/Indroduction'
export default {
    components: {
        Parallax,
        Intoduction,
    },
    data() {
        return {
            activeIndex: 0,
            allImages: [
                'https://cdn.spacetelescope.org/archives/images/wallpaper2/heic2007a.jpg',
                'https://apod.nasa.gov/apod/image/2006/catseye2_not_2048.jpg',
                'https://www.gettyimages.co.uk/gi-resources/images/500px/983801190.jpg',
                'https://cdn.spacetelescope.org/archives/images/wallpaper2/heic2007a.jpg',
                'https://apod.nasa.gov/apod/image/2006/catseye2_not_2048.jpg',
                'https://www.gettyimages.co.uk/gi-resources/images/500px/983801190.jpg',
                'https://cdn.spacetelescope.org/archives/images/wallpaper2/heic2007a.jpg',
                'https://apod.nasa.gov/apod/image/2006/catseye2_not_2048.jpg',
                'https://www.gettyimages.co.uk/gi-resources/images/500px/983801190.jpg',
            ],
        }
    },
    computed: {
        images() {
            const images = []
            let first =
                this.activeIndex === 0
                    ? this.allImages[this.allImages.length - 1]
                    : this.allImages[this.activeIndex - 1]
            let last =
                this.activeIndex === this.allImages.length - 1
                    ? this.allImages[0]
                    : this.allImages[this.activeIndex + 1]
            images[0] = first
            images[1] = this.allImages[this.activeIndex]
            images[2] = last
            if (this.activeIndex === this.allImages.length - 1) {
                this.activeIndex = 0
            }

            return images
        },
    },
}
</script>

<style scoped>
* {
    box-sizing: border-box;
}
.page-body {
    display: flex;
    flex-direction: column;
}
.cp {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 100%;
    max-width: 1500px;
    margin: auto;
    height: 100vh;
}
.cp div {
    display: flex;
    flex-direction: row;
    /* max-width: 1500px; */
    width: 100%;
    justify-content: space-between;
}
img {
    display: block;
    max-width: 30%;
    /* filter: blur(); */
    transform: scale(0.7);
    opacity: 0.7;
}
.active {
    /* filter: brightness(); */
    transform: scale(1);
    opacity: 1;
}
@media (max-width: 800px) {
    h1 {
        font-size: 2rem;
    }
    p {
        font-size: 1rem;
        padding: 0.75rem 0;
    }
}
</style>
