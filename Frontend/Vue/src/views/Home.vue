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
        <CardThematic :propThematic="state.thematicsData" />
      </div>
    </div>
  </div>
</template>

<script>
import { reactive } from "vue";
import Carousel from "../components/client/Carousel_welcome.vue";
import CardThematic from "../components/client/Card_Thematic.vue";
import { useThematicInfinite } from "../composables/thematics/useThematics";
import { useCountThematics } from "../composables/thematics/useThematics";

export default {
  setup() {
    const state = reactive({
      thematicsData: useThematicInfinite(1, 6),
      scroll: {
        offset: 0,
        prod_tot: useCountThematics(),
        active: "no",
      },
    });

    const makeScroll = () => {
      let scrollTop = document.documentElement.scrollTop;
      let scrollHeight = document.documentElement.scrollHeight;
      let clientHeight = document.documentElement.clientHeight;

      if (state.scroll.offset >= state.scroll.prod_tot) {
        state.scroll.active = "yes";
      }

      if (state.scroll.active == "no") {
        if (scrollTop + clientHeight >= scrollHeight) {
          state.scroll.offset += 6;
          state.listTmp = useThematicInfinite(state.scroll.offset, 6);
          setTimeout(() => {
            state.listTmp.map((item) => {
              state.thematicsData.push(item);
            });
          }, 100);
        }
      }
    };

    return { state, scroll, makeScroll };
  },

  mounted() {
    window.addEventListener("scroll", this.makeScroll);
  },
  components: { Carousel, CardThematic },
};
</script>

<style lang="scss">
@import url("https://fonts.googleapis.com/css2?family=Lobster&display=swap");

.homeHtml {
  background-color: rgb(201, 243, 234);
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
