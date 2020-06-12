<template>
    <div>
        <img :src="article.imageUrl" :alt="article.title" />
        <p>{{ article.title }}</p>
    </div>
</template>

<script>
export default {
    props: ['article'],
}
</script>

<style scoped>
div {
    background-color: #1b1b1b;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.4), 0 6px 20px 0 rgba(0, 0, 0, 0.4);
    margin: 0.5rem;
    cursor: pointer;
    border-radius: 4px;
    display: flex;
    flex-direction: column;
}
img {
    width: 100%;
    display: block;
    border-top-left-radius: 4px;
    border-top-right-radius: 4px;
}
p {
    padding: 1rem 1rem;
    font-size: 1.2rem;
    text-align: center;
    align-content: center;
    margin: auto;
    color: white;
    font-weight: bold;
}
</style>
