<template>
  <Carousel class="carousel" :items-to-show="2.5" :wrap-around="true" :autoplay="3000">
    <Slide v-for="img in imgCarousel" :key="img">
      <div class="carousel__item">
        <router-link to="/table/all">
            <div class="card" style="width: 17rem">
            <img :src="'/img/' + img" class="card-img-top" alt="Not img"/>
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
    let imgCarousel = ["food1.278a0d27.jpg","food2.862afba0.jpg","food3.19e5461a.jpg","food4.d3341c54.jpg","food5.83f0a590.jpg","food6.df6297f7.jpg","food7.c5814fb5.jpg","food8.ab06ff80.jpg","food9.9e0356c7.jpg","food10.b352d1ee.jpg","food11.4dc6f13b.jpg","food12.4a47d89c.jpg", ];
    return { imgCarousel};
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
      border:3px groove black ;
    }
  }
}
</style>