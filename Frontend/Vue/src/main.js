import { createApp } from 'vue';
import 'bootstrap/dist/css/bootstrap.css';
import './main.css';
import App from './App.vue'

import router from "./router";
import store from "./store";

// vue-icons
import { OhVueIcon, addIcons } from "oh-vue-icons";
import * as VIcons from "oh-vue-icons/icons";

const Vi = Object.values({ ...VIcons })
addIcons(...Vi);

createApp(App)
  .use(router)
  .use(store)
  .component('v-icon', OhVueIcon)
  .mount("#app");

