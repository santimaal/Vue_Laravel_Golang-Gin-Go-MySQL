<template>
  <div class="all_cards" v-if="propThematic.length > 0">
    <div class="maincontainer" v-for="thematic in propThematic" :key="thematic.id">
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
  </div>
  <div v-else>
    <span>Don't have thematics</span>
  </div>
</template>

<script>
import { useRouter } from "vue-router";

export default {
  props: {
    propThematic: Object
  },

  setup() {
    const router = useRouter();

    const redirectReserve = (id) => {
      console.log("Click en la Theamtica " + id);
      let search = { id_thematic: id };

      router.push({
        name: "client_table",
        params: { filter: btoa(JSON.stringify(search)) },
      });
    };

    return { redirectReserve };
  },
};
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Lobster&display=swap");

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
  border: 2px solid lightslategray;
}

.back h2 {
  position: absolute;
  font-style: italic;
  color: lightslategray;
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
  border: 2px solid lightslategray;
  font-family: "Lobster", cursive;
  font-size: large;
}

.back button:hover {
  background-color: black;
  color: aliceblue;
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

.maincontainer>.front {
  position: absolute;
  transform: perspective(600px) rotateY(0deg);
  width: 302px;
  height: 290px;
  backface-visibility: hidden;
  transition: transform 0.5s linear 0s;
}

.maincontainer>.back {
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

.maincontainer:hover>.front {
  transform: perspective(600px) rotateY(-180deg);
}

.maincontainer:hover>.back {
  transform: perspective(600px) rotateY(0deg);
}
</style>
