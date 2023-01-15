<template>
  <div class="all height">
    <Header />
    <router-view></router-view>
    <Footer />
  </div>
</template>

<script setup>
import Header from "./components/Header";
import Footer from "./components/Footer";
import 'bootstrap/dist/js/bootstrap.js'
import Constant from "./Constant";
import { useStore } from "vuex";
const store = useStore();

if (localStorage.getItem('token')) {
  store.dispatch("user/" + Constant.GET_PROFILE);
}


</script>

<style>
.all {
  background-color: rgb(201, 243, 234);
}
</style>
