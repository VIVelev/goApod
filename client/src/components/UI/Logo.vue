<template>
    <span>
        <p>go</p>
        <img src="../../assets/logo.png" alt="logo" />
        <p>APOD</p>
    </span>
</template>

<style scoped>
span {
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 2rem;
    margin: 0;
}
img {
    max-width: 2rem;
}
p {
    margin: 0;
    color: black;
    text-shadow: -1.2px -1.2px 0 white, 1.2px -1.2px 0 white,
        -1.2px 1.2px 0 white, 1.2px 1.2px 0 white;
}
</style>
