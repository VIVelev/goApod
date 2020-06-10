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
* {
    box-sizing: border-box;
}
.page-body {
    display: flex;
    flex-direction: column;
}
.top-rated-photos {
    display: flex;
    flex-direction: column;
    height: 100vh;
    justify-content: center;
    align-items: center;
}
h1 {
    font-size: 4rem;
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
