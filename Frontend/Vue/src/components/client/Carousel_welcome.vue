<template>
  <Carousel
    class="carousel"
    :items-to-show="2.5"
    :wrap-around="true"
    :autoplay="3000"
  >
    <Slide v-for="img in imgCarousel" :key="img">
      <div class="carousel__item">
        <router-link to="/table/all">
          <div class="card" style="width: 17rem">
            <img :src="img" class="card-img-top" alt="Not img" />
          </div>
        </router-link>
      </div>
    </Slide>

    <template #addons>
      <Navigation />
    </template>
  </Carousel>
</template>

<script>
import "vue3-carousel/dist/carousel.css";
import { Carousel, Slide, Navigation } from "vue3-carousel";

export default {
  name: "App",
  components: {
    Carousel,
    Slide,
    Navigation,
  },
  setup() {
    let imgCarousel = [
      "https://media.smartbox.com/pim/1000002221294892070931.jpg",
      "https://media.smartbox.com/pim/1000002419655147719912.jpg",
      "https://media.smartbox.com/pim/1000002338669465665301.jpg",
      "https://media.smartbox.com/pim/100000233866611568404.jpg",
      "https://media.smartbox.com/pim/1000002456861956491381.jpg",
      "https://media.smartbox.com/pim/10000022211751694910013.jpg",
      "https://media.smartbox.com/pim/10000022106851519717083.jpg",
      "https://media.smartbox.com/pim/1000002221294892070931.jpg?",
      "https://media.smartbox.com/pim/10000024520101939582082.jpg",
      "https://media.smartbox.com/pim/1000002211695727769637.jpg",
      "https://media.smartbox.com/pim/10000022173171214476114.jpg",
      "https://media.smartbox.com/pim/10000025856661564828011.jpg",
    ];
    return { imgCarousel };
  },
};
</script>

<style lang="scss">
.carousel {
  margin-bottom: 2%;
  .card {
    border-radius: 30px;
    img {
      border-radius: 30px;
      border: 3px groove black;
    }
  }
}
</style>
