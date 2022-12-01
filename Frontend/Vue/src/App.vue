<template>
  <Header v-if="head.option == ''" />
  <HeaderC v-if="head.option == 'client'" />
  <HeaderA v-if="head.option == 'admin'" />
  <button @click="cambia('client')">Client</button>
  <button @click="cambia('admin')">Admin</button>
  <router-view></router-view>
</template>

<script>
import Header from './components/Header';
import HeaderA from './components/admin/Header';
import HeaderC from './components/client/Header';

import { reactive } from 'vue'

export default {
  components: { Header,HeaderC,HeaderA },
  setup() {

    const head = reactive({
      option:''
    });

    const cambia = (result) => {
      head.option=result
    }
    return { head, cambia }
  }
}
</script>

<style>

</style>