<template>
  <nav class="navbar navbar-expand-md navbar-dark">
    <router-link class="nav-link" to="/home"><img class="logo" src="../assets/logo/logo_vert.png" alt=""/></router-link>
    <button class="navbar-toggler" type="button" @click="changeIsNavShow">
      <span class="navbar-toggler-icon"></span>
    </button>

    <div :class="navClass">
      <ul class="navbar-nav navRight">
        <li class="nav-item">
          <router-link class="nav-link  button" to="/home">Home</router-link>
        </li>
        <li class="nav-item">
          <router-link class="nav-link button" to="/table/all">Reserves</router-link>
        </li>
        <li class="nav-item">
          <router-link class="nav-link button" to="/atable">DashboardTable</router-link>
        </li>
        <!-- <li class="nav-item">
                    <router-link class="nav-link liNav" to="/thematic">ThematicList</router-link>
                </li> -->
        <li class="nav-item">
          <router-link class="nav-link button" to="/athematic">DashboardThematic</router-link>
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
      justify-content: end;
      .nav-item{
        margin-right: 2%;
      }
    }
  }
}

.button:focus {
  outline: none;
}

.button {
  text-transform: uppercase;
  width: 100%;
  background: none;
  position: relative;
  z-index: 1;
  letter-spacing: 2px;
  color: #fff;
  font-weight: bold;
  font-size: medium;
  border: 4px solid;
  cursor: pointer;
}

.button::after {
  content: "";
  background-color: transparent;
  position: absolute;
  top: -20px;
  left: -20px;
  right: -20px;
  bottom: -20px;
  opacity: 0;
  z-index: -1;
}

.button:hover::after {
  opacity: 1;
  background-color: white;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}

.button,
.button:after {
  -webkit-transition: all 0.5s ease-out;
  -moz-transition: all 0.5s ease-out;
  -o-transition: all 0.5s ease-out;
  transition: all 0.5s ease-out;
}

.button:hover {
  border: 4px solid white;
  color: black !important;
  border-radius: 5px;
}

//Icon Toggler
.navbar-toggler {
  display: flex;
  margin-left: 90%;
  margin-top: -10%;
}
</style>
