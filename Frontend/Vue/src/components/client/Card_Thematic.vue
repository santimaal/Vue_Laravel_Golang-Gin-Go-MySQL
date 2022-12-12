<template>
  <div class="all_cards" v-if="thematicsData.length > 0">
    <div
      class="maincontainer"
      v-for="thematic in thematicsData"
      :key="thematic.id"
    >
      <div class="back">
        <h2>{{ thematic.name }}</h2>
        <p>
          In this {{ thematic.name }} you can find the best gastronomic dishes
          of {{ thematic.location }}. You can also have a unique experience that
          you will remember for the rest of your life
        </p>
        <button @click="redirectReserve(thematic.id)">Reserve table</button>
      </div>
      <div class="front">
        <div class="image">
          <img :src="thematic.img" />
          <h2>{{ thematic.name }}</h2>
        </div>
      </div>
    </div>
    <!-- <InfiniteLoading @infinite="scroll" :distance="1"/> -->
  </div>
  <div v-else>
    <span>Don't have thematics</span>
  </div>
</template>

<script>
import { useThematicInfinite } from "../../composables/thematics/useThematics";
import { reactive } from "vue";
import { useRouter } from "vue-router";

// import InfiniteLoading from "v3-infinite-loading";
// import "v3-infinite-loading/lib/style.css";

export default {
  // components: { InfiniteLoading },
  setup() {
    const router = useRouter();
    const Thematics = useThematicInfinite();
    var thematicsData = reactive(Thematics);

    const redirectReserve = (id) => {
      console.log("Click en la Theamtica " + id);
      let search = { id_thematic: id };

      router.push({
        name: "client_table",
        params: { filter: btoa(JSON.stringify(search)) },
      });

      // localStorage.setItem('id', id);
      // window.location.href="/#/reservation";
    };

    const scroll = {
      prueba: "hola",
    };
    console.log(scroll.prueba);
    return { thematicsData, redirectReserve, scroll };
  },
};
</script>

<style>
.all_cards {
  display: grid;
  grid-template-columns: repeat(3, 20%);
  justify-content: center;
  justify-items: center;
  grid-gap: 0.8rem 9rem;
  margin-bottom: 2%;
  margin-top: 2%;
}

.maincontainer {
  width: 302px;
  height: 299px;
}

.back {
  border: 2px solid #3ec9a5;
}

.back h2 {
  position: absolute;
  font-style: italic;
  color: #3ec9a5;
}

.back p {
  position: absolute;
  top: 20%;
  font-size: medium;
}

.back button {
  position: absolute;
  width: 80%;
  margin-top: 70%;
  margin-left: -40%;
  border-radius: 10px;
}

.front img {
  object-fit: cover;
  width: 100%;
  height: 299px;
  opacity: 0.7;
  border: 3px black solid;
  border-radius: 10px;
}

.front h2 {
  position: absolute;
  padding: 10px;
  top: 200px;
  color: #000;
}
.maincontainer > .front {
  position: absolute;
  transform: perspective(600px) rotateY(0deg);
  width: 302px;
  height: 290px;
  backface-visibility: hidden;
  transition: transform 0.5s linear 0s;
}

.maincontainer > .back {
  position: absolute;
  transform: perspective(600px) rotateY(180deg);
  background: #262626;
  color: #fff;
  width: 302px;
  height: 290px;
  border-radius: 10px;
  padding: 5px;
  backface-visibility: hidden;
  transition: transform 0.5s linear 0s;
}

.maincontainer:hover > .front {
  transform: perspective(600px) rotateY(-180deg);
}

.maincontainer:hover > .back {
  transform: perspective(600px) rotateY(0deg);
}
</style>
