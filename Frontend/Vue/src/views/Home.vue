<template>
  <div class="homeHtml">
    <div class="wl">
      <h2>
        Welcome to Sanvic, the restaurant with the most varied gastronomy in
        Spain
      </h2>
    </div>
    <div>
      <h2>Gastronomic dishes</h2>
      <Carousel></Carousel>
    </div>
    <div>
      <h2>Types of gastronomy</h2>
      <div class="d-flex justify-content-center">
        <!-- <CardThematic v-for="thematicitem in state.thematiclist" :key="thematicitem.id" :thematicitem="thematicitem"/> -->
        <CardThematic/>

      </div>
    </div>
  </div>
</template>

<script>
import Constant from "../Constant";
import { reactive, computed } from "vue";
import { useStore } from "vuex";

import Carousel from "../components/client/Carousel_welcome.vue";
import CardThematic from "../components/client/Card_Thematic.vue";

export default {
  setup() {
    const store = useStore();

    store.dispatch("thematic/" + Constant.INITIALIZE_THEMATIC);

    const state = reactive({
      thematiclist: computed(() => store.getters["thematic/getThematic"]),
    });
    return { state };
  },
  components: { Carousel, CardThematic },
};
</script>

<style lang="scss">
@import url("https://fonts.googleapis.com/css2?family=Lobster&display=swap");

.homeHtml {
  text-align: center;
  .wl {
    text-align: center;
    h2 {
      margin-top: 0.5%;
      color: rgb(161, 158, 158);
      font-size: large;
    }
  }
  h2 {
    font-family: "Lobster", cursive;
    font-size: xx-large;
  }
}
</style>
