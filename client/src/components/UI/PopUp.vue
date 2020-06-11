<template>
    <div>
        <component
            @changeComponent="component = $event"
            :is="component"
        ></component>
    </div>
</template>

<script>
import SignUp from '../Auth/SignUp'
import SignIn from '../Auth/SignIn'
export default {
    components: {
        SignUp,
        SignIn,
    },
    data() {
        return {
            component: 'SignUp',
        }
    },
}
</script>

<style scoped>
div {
    position: fixed;
    top: 35%;
    left: 0;
    right: 0;
    width: 90%;
    background: white;
    border-radius: 4px;
    max-width: 500px;
    z-index: 3;
    margin: auto;
}
</style>
