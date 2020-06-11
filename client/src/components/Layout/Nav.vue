<template>
    <nav>
        <ul>
            <router-link class="logo" to="/" tag="li">
                <Logo />
            </router-link>
            <div class="routes">
                <li>
                    <router-link to="/posts" tag="p">
                        All Posts
                    </router-link>
                </li>
                <li>
                    <router-link to="/today" tag="p">
                        Today's Picture
                    </router-link>
                </li>
                <li v-if="$store.getters.user">
                    <p @click="$store.commit('changeUser', null)">Logout</p>
                </li>
                <li v-else>
                    <p @click="$store.commit('modifyPopUp', true)">Sign in</p>
                </li>
            </div>
            <li class="burger">
                <div></div>
                <div></div>
                <div></div>
            </li>
        </ul>
    </nav>
</template>

<script>
import Logo from '../UI/Logo'
export default {
    components: {
        Logo,
    },
}
</script>

<style scoped>
* {
    outline: none;
    box-sizing: border-box;
}
nav {
    position: fixed;
    left: 0;
    right: 0;
    top: 0;
    background-color: black;
}
ul {
    margin: 0 auto;
    padding: 0;
    list-style: none;
    max-width: 1500px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
}
.logo {
    cursor: pointer;
}
.routes {
    display: flex;
}
.routes li {
    margin-left: 1rem;
}
p {
    color: white;
    font-weight: bold;
    font-size: 1.1rem;
    border: 2px solid transparent;
    padding: 0.75rem 0;
    margin: 1rem 0;
    transition: 0.7s;
    border-radius: 4px;
    cursor: pointer;
}
p:hover {
    border-bottom-color: white;
}
.router-link-active {
    border-bottom-color: white;
}
.burger {
    width: 30px;
    height: 30px;
    background-color: transparent;
    display: flex;
    flex-direction: column;
    justify-content: space-evenly;
    margin: 1rem 0;
    display: none;
}
.burger div {
    height: 2px;
    width: 100%;
    background-color: white;
    border-radius: 4px;
}
@media (max-width: 1500px) {
    ul {
        padding: 0 1rem;
    }
}
@media (max-width: 800px) {
    .routes {
        display: none;
    }
    .burger {
        display: flex;
    }
}
</style>
