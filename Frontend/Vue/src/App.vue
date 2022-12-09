<template>
  <!-- <Header v-if="head.option == ''" />
  <HeaderC v-if="head.option == 'client'" />
  <HeaderA v-if="head.option == 'admin'" /> -->
  <!-- <button @click="cambia('client')">Client</button>
  <button @click="cambia('admin')">Admin</button> -->
  <div class="all">
    <Header />
    <router-view></router-view>
    <Footer />
  </div>
</template>

<script>
// import HeaderA from './components/admin/Header';
// import HeaderC from './components/client/Header';
import Header from "./components/Header";
import Footer from "./components/Footer";

import { reactive } from "vue";

export default {
  components: { Header, Footer /*,HeaderC,HeaderA*/ },
  setup() {
    const head = reactive({
      option: "",
    });

    const cambia = (result) => {
      head.option = result;
    };
    return { head, cambia };
  },
};
</script>

<style>
.all {
  background-color: rgb(201, 243, 234);
}
</style>
