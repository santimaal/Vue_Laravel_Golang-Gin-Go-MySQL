<template>
  <nav class="navbar navbar-expand-md navbar-dark">
    <router-link class="nav-link" to="/home"><img class="logo" src="../assets/logo_vert.png" alt=""/></router-link>
    <button class="navbar-toggler" type="button" @click="changeIsNavShow">
      <span class="navbar-toggler-icon"></span>
    </button>

    <div :class="navClass">
      <ul class="navbar-nav navRight">
        <li class="nav-item">
          <router-link class="nav-link liNav" to="/home">Home</router-link>
        </li>
        <li class="nav-item">
          <router-link class="nav-link liNav" to="/table/all">Reserves</router-link>
        </li>
        <li class="nav-item">
          <router-link class="nav-link liNav" to="/atable">DashboardTable</router-link>
        </li>
        <!-- <li class="nav-item">
                    <router-link class="nav-link liNav" to="/thematic">ThematicList</router-link>
                </li> -->
        <li class="nav-item">
          <router-link class="nav-link liNav" to="/athematic">DashboardThematic</router-link>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script>
import { reactive, computed } from "vue";

export default {
  setup() {
    const state = reactive({
      isNavShow: false,
    });
    const navClass = computed(() =>
      state.isNavShow
        ? "collapse navbar-collapse show"
        : "collapse navbar-collapse"
    );
    const changeIsNavShow = () => {
      state.isNavShow = !state.isNavShow;
    };

    return { state, changeIsNavShow, navClass };
  },
};
</script>

<style lang="scss">
nav {
  width: 100% !important;
  background-color: black !important;
  .logo {
    width: 40%;
    margin-left: 5%;
  }
  div {
    .navRight {
      width: 100%;
      display: flex;
      justify-content: space-evenly;
      .liNav {
        color: white !important;
        font-weight:bold;
        font-size: large;
        background-color: rgba(61,190,176,255);
        border-radius: 10px;
        border: 1px solid orange;
      }
    }
  }
}
.navbar-toggler {
  display: flex;
  margin-left: 90%;
  margin-top: -10%;
}
</style>
