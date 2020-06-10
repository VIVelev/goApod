<template>
    <div @click="$store.commit('modifyPopUp', false)"></div>
</template>

<script>
export default {}
</script>

<style scoped>
div {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background-color: rgba(0, 0, 0, 0.8);
    z-index: 2;
}
</style>
