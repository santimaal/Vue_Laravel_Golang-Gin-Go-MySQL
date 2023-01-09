import { createApp } from 'vue';
import 'bootstrap/dist/css/bootstrap.css';
import './main.css';
import App from './App.vue'
import Toaster from "@meforma/vue-toaster";
import router from "./router";
import store from "./store";

//font-awesome-icons
import { library } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

import {faMagnifyingGlass,faPerson,faLocationDot,faUtensils,faPersonShelter,faRightFromBracket, faArrowsRotate} from '@fortawesome/free-solid-svg-icons';
library.add(faMagnifyingGlass,faPerson,faLocationDot,faUtensils,faPersonShelter,faRightFromBracket,faArrowsRotate);

 const app= createApp(App)
  app.use(router);
  app.use(store);
  app.use(Toaster);
  app.component('font-awesome-icon', FontAwesomeIcon);
  app.mount("#app");

