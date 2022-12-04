<template>
    <nav class="navbar navbar-expand-md navbar-dark">
        <span><img src="../../assets/sanvic.png" alt=""></span>
        <button class="navbar-toggler" type="button" @click="changeIsNavShow">
            <span class="navbar-toggler-icon"></span>
        </button>

        <div :class="navClass">
            <ul class="navbar-nav">
                <li class="nav-item">
                    <router-link class="nav-link" to="/home">Home</router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link" to="/atable">TablelistAdmin</router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link" to="/athematic">ThematicAdmin</router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link" to="/reserve">Reserva</router-link>
                </li>
            </ul>
        </div>
    </nav>
</template>

<script>
import { reactive, computed } from 'vue';

export default {
    setup() {
        const state = reactive({
            isNavShow: false
        })
        const navClass = computed(() => state.isNavShow ? "collapse navbar-collapse show" : "collapse navbar-collapse")
        const changeIsNavShow = () => {
            state.isNavShow = !state.isNavShow;
        }

        return { state, changeIsNavShow, navClass };
    }
}
</script>

<style lang="scss">
nav {
    width: 100% !important;
    background-color: black;

    img {
        width: 20%;
    }
}
</style>