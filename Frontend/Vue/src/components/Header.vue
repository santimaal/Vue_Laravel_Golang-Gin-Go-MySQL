<template>
  <nav class="navbar navbar-expand-md navbar-dark">
    <router-link class="nav-link" to="/home"><img class="logo" src="../assets/logo/sanvic.png" alt="" /></router-link>
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
        <li class="nav-item dropdown" v-if="state.auth == 'admin'">
          <a class="nav-link dropdown-toggle button" href="#" role="button" data-toggle="dropdown" aria-expanded="false">Dashboard</a>
          <div class="dropdown-menu">
            <a class="dropdown-item" href="#"><router-link class="nav-link button" to="/atable">Table</router-link></a>
            <a class="dropdown-item" href="#"><router-link class="nav-link button" to="/athematic">Thematic</router-link></a>
            <a class="dropdown-item" href="#"><router-link class="nav-link button" to="/athematic">User</router-link></a>
            <a class="dropdown-item" href="#"><router-link class="nav-link button" to="/athematic">Reserve</router-link></a>
          </div>
        </li>
        <li class="nav-item" v-if="state.auth == ''">
          <router-link class="nav-link button" to="/login">Login</router-link>
        </li>
        <li class="nav-item" v-if="state.auth == ''">
          <router-link class="nav-link button" to="/register">Register</router-link>
        </li>
        <li class="nav-item" v-if="state.auth">
          <router-link class="nav-link button" to="/profile">{{ state.user.name }}</router-link>
        </li>
        <li class="nav-item" v-if="state.auth != ''">
          <a class="nav-link button" @click="logout"><font-awesome-icon class="logout_icon"
              icon="fa-solid fa-right-from-bracket" /></a>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script>
import { reactive, computed } from "vue";
import { useStore } from "vuex";
import Constant from '../Constant';



export default {
  setup() {
    const store = useStore()
    const state = reactive({
      isNavShow: false,
      user: computed(() => store.getters["user/getUser"]),
      auth: computed(() => store.getters["user/getAuth"])
    });
    const navClass = computed(() =>
      state.isNavShow
        ? "collapse navbar-collapse show"
        : "collapse navbar-collapse"
    );
    const changeIsNavShow = () => {
      state.isNavShow = !state.isNavShow;
    };

    const logout = () => {
      store.dispatch("user/" + Constant.LOGOUT)
    }

    return { state, changeIsNavShow, navClass, logout };
  },
};
</script>

<style lang="scss">
nav {
  width: 100% !important;
  background-color: black !important;
  height: 12vh;

  .logo {
    width: auto;
    margin-left: 5%;
    height: 10vh;
  }

  div {
    .navRight {
      width: 100%;
      display: flex;
      justify-content: end;

      .nav-item {
        margin-right: 2%;
        background-color: black;
        z-index: 100;

        .logout_icon {
          height: 20px;
        }

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
  margin-top: -15%;
}

.collapse ul li .button {
  padding-left: 2%;
}
.dropdown-menu{
  background-color: transparent;
  border: 0px;
  margin-right: 10%;
}
.dropdown-item .button{
  background-color: black;
  text-align: center;
}

.dropdown-item:hover {
  background-color: transparent;
}

.dropdown-menu {
  margin-right: 20% !important;
}


</style>
