import Constant from "../../Constant";
import UserService from "../../services/UserService";
import { createToaster } from "@meforma/vue-toaster";
import router from '../../router/index';
const toaster = createToaster();

export const user = {
  state: {
    user: {
      is_active: false,
      name: "",
      email: "",
      img: "",
      type: "",
    },
  },
  namespaced: true,
  mutations: {
    [Constant.SET_USER]: (state, payload) => {
      console.log(payload);
      state.user = {
        is_active: payload.is_active,
        name: payload.name,
        email: payload.email,
        img: payload.img,
        type: payload.type,
      };
    },
    [Constant.USER_REGISTER]: (state, payload) => {
      toaster.success("User " + payload.name.toUpperCase() + " has successfully registered",{ position: "top-right", duration: 5000, dismissible: true });
      router.push({ name: 'login' });
    },
    [Constant.USER_LOGIN]: (state, payload) => {
        state.user = {
          is_active: payload.is_active,
          name: payload.name,
          email: payload.email,
          img: payload.img,
          type: payload.type,
        };
        router.push({ name: 'home' });
    },
    [Constant.LOGOUT]: (state) => {
      state.user = {
        is_active: "",
        name: "",
        email: "",
        img: "",
        type: "",
      };
      router.push({ name: 'home' });
    },
  },
  actions: {
    [Constant.GET_PROFILE]: (/*store*/) => {
        UserService.getProfile()
          .then(function (res) {
            console.log(res);
          })
          .catch(function (error) {
            console.log(error);
          });
    },
    [Constant.USER_REGISTER]: (store, payload) => {
      UserService.userRegister(payload)
        .then(function (res) {
            if(res.status == 200){
                store.commit(Constant.USER_REGISTER, payload);
            } 
        })
        .catch(function (error) {
          if (error["response"]["data"] == "Email is registered") {
            toaster.error("Email exist", { position: "bottom", duration: 5000, dismissible: true });
          } else {
            toaster.error("Error register", { position: "top-right", duration: 5000, dismissible: true });
          }
        });
    },
    [Constant.USER_LOGIN]: (store, payload) => {
      UserService.userLogin(payload)
        .then(function (res) {
            if (res.data.type == "admin") {
                UserService.userLoginAdmin(payload)
                .then(function (res) {
                  console.log(res);
                })
                .catch(function (error) {
                    console.log(error);
                    toaster.error("Error login Admin", { position: "bottom", duration: 5000, dismissible: true });
                });
            }else {
                store.commit(Constant.USER_LOGIN, res.data);
            }
          
          localStorage.setItem("token", JSON.stringify(res.data.token));
        })
        .catch(function (error) {
            if (error["response"]["data"] == "Email or password is not correct") {
                toaster.error("Email or password is not correct", { position: "bottom", duration: 5000, dismissible: true });
            } else {
                toaster.error("Error login", { position: "top-right", duration: 5000, dismissible: true });
            }
        });
    },
    [Constant.USER_REGISTER_ADMIN]: (store, payload) => {
      UserService.userRegisterAdmin(payload)
        .then(function (res) {
            if(res.status == 200){
                store.commit(Constant.USER_REGISTER, payload);
            }
        })
        .catch(function (error) {
          if ( error["response"]["data"]["errors"]["email"] || error["response"]["data"]["errors"]["name"]) {
            if (error["response"]["data"]["errors"]["email"] == "validation.unique") {
              toaster.error("Email exist", { position: "bottom", duration: 5000, dismissible: true });
            }
            if (error["response"]["data"]["errors"]["name"] == "validation.unique") {
              toaster.error("Name exist", { position: "bottom", duration: 5000, dismissible: true });
            }
          } else {
            toaster.error("Error register", { position: "top-right", duration: 5000, dismissible: true });
          }
        });
    },
    [Constant.LOGOUT]: (store) => {
      localStorage.removeItem("token");
      store.commit(Constant.LOGOUT);
    },
  },
  getters: {
    getUser(state) {
      return state.user;
    },
    getAuth(state) {
      return state.user.type;
    },
  },
};
