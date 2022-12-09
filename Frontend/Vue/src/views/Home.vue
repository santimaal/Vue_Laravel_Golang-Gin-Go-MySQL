<template>
  <div class="homeHtml">
    <div class="wl">
      <h2>Welcome to Sanvic, the restaurant with the most varied gastronomy in Spain</h2>
    </div>
    <div>
        <h2>Gastronomic dishes</h2>
        <Carousel></Carousel>
    </div>
  </div>
</template>

<script>
import Carousel from "../components/client/Carousel_welcome.vue";

export default {
  setup() {
    // const store = useStore();
    // const router = useRouter();
    // const currentRoute = useRoute();
  },
  components: { Carousel },
};
</script>

<style lang="scss">
@import url('https://fonts.googleapis.com/css2?family=Lobster&display=swap');

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
  h2{
    font-family: 'Lobster', cursive;
    font-size:xx-large;
  }
}
</style>
