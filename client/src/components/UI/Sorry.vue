<template>
    <div>
        <h1 class="error"><slot></slot></h1>
        <h1 class="face">:(</h1>
    </div>
</template>

<script>
export default {}
</script>

<style scoped>
div {
    position: absolute;
    top: 15%;
    left: 50%;
    transform: translate(-50%);
    color: #1b1b1b;
}
.error {
    font-size: 3rem;
    text-align: center;
}
.face {
    transform: rotate(90deg);
    text-align: center;
    font-size: 5rem;
}
</style>
