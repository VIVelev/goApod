<template>
    <div>
        <h1>Discover the cosmos!</h1>
        <p>
            Each day a different image or photograph of our fascinating universe
            is featured, along with a brief explanation written by a
            professional astronomer.
        </p>
    </div>
</template>

<script>
export default {}
</script>

<style scoped>
div {
    display: grid;
    grid-template-columns: 1fr 1500px 1fr;
    grid-template-rows: auto auto;
    align-content: center;
    height: 75vh;
    background-color: black;
}
@media (max-width: 1500px) {
    div {
        grid-template-columns: 1rem 1fr 1rem;
    }
}
h1 {
    grid-column: 2/3;
    grid-row: 1;
    font-size: 5rem;
    margin: 0;
    width: fit-content;
    color: white;
}
p {
    grid-column: 2/3;
    grid-row: 2;
    margin: 1rem 0;
    padding: 1.5rem 0;
    border-top: 3px solid rgba(255, 255, 255, 0.75);
    font-size: 1.2rem;
    font-weight: 600;
    color: rgba(255, 255, 255, 0.75);
    text-align: justify;
    max-width: 450px;
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
