<template>
    <div class="page-body">
        <Intoduction />
        <div class="top-rated-photos">
            <h1>Top rated photos</h1>
            <Slider></Slider>
        </div>
    </div>
</template>

<script>
import Intoduction from '../components/UI/Indroduction'
import Slider from '../components/Slider/Slider'
export default {
    components: {
        Intoduction,
        Slider,
    },
}
</script>

<style scoped>
.page-body {
    display: flex;
    flex-direction: column;
}
.top-rated-photos {
    display: flex;
    flex-direction: column;
    height: 100vh;
    justify-content: flex-end;
    align-items: center;
    /* position: relative; */
}
h1 {
    font-size: 4rem;
    margin-top: 10rem;
    text-align: center;
    color: #1b1b1b;
    /* top: 1rem; */
    /* top: 10%; */
}
@media (max-width: 800px) {
    h1 {
        font-size: 3rem;
    }
    p {
        font-size: 1rem;
        padding: 0.75rem 0;
    }
}
@media (max-width: 500px) {
    h1 {
        font-size: 2rem;
    }
}
</style>
