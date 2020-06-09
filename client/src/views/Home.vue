<template>
    <div class="page-body">
        <Intoduction />
        <div class="cp">
            <h1>
                Top Rated Photos
            </h1>
            <div>
                <img
                    src="https://apod.nasa.gov/apod/image/2006/catseye2_not_2048.jpg"
                    alt="ddz"
                />
                <img
                    class="active"
                    src="https://apod.nasa.gov/apod/image/2006/catseye2_not_2048.jpg"
                    alt="ddz"
                />
                <img
                    src="https://apod.nasa.gov/apod/image/2006/catseye2_not_2048.jpg"
                    alt="ddz"
                />
            </div>
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
}
</script>

<style scoped>
.page-body {
    display: flex;
    flex-direction: column;
}
.cp {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    max-width: 1500px;
    width: 100%;
    margin: auto;
    height: 75vh;
    background-color: rgb(255, 255, 255);
}
.cp div {
    display: flex;
    flex-direction: row;
    max-height: 60%;
    overflow: hidden;
    width: 90%;
    align-items: center;
    padding: 0 1rem;
    justify-content: center;
}
h1 {
    color: black;
    font-size: 4rem;
}
img {
    height: 100%;
    display: block;
    margin: 2rem;
    filter: blur(1rem);
    transform: scale(0.8);
    opacity: 0.7;
}
.active {
    filter: brightness();
    align-self: center;
    transform: scale(1.1);
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
