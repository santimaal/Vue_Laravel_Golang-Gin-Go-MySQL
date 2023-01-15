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
      noti:0,
      chat_id:"",
    },
  },
  namespaced: true,
  mutations: {
    [Constant.SET_USER]: (state, payload) => {
      state.user = {
        is_active: payload.is_active,
        name: payload.name,
        email: payload.email,
        img: payload.img,
        type: payload.type,
        noti: payload.noti,
        chat_id: payload.chat_id
      };
      if (!payload.redirect){
        router.push({ name: 'home' });
      }
    },
    [Constant.USER_REGISTER]: (state, payload) => {
      toaster.success("User " + payload.name.toUpperCase() + " has successfully registered", { position: "top-right", duration: 5000, dismissible: true });
      router.push({ name: 'login' });
    },
    [Constant.LOGOUT]: (state) => {
      toaster.success("User " + state.user.name.toUpperCase() + " has log out successfully", { position: "top-right", duration: 5000, dismissible: true });
      state.user = {
        is_active: "",
        name: "",
        email: "",
        img: "",
        type: "",
        noti: 0
      };
      router.push({ name: 'home' });
    },
  },
  actions: {
    [Constant.GET_PROFILE]: (store) => {
      
      UserService.getProfile()
        .then(function (res) {
          store.commit(Constant.SET_USER, res.data);
        })
        .catch(function (error) {
          console.log(error);
        });
     
    },
    [Constant.USER_REGISTER]: (store, payload) => {
      UserService.userRegister(payload)
        .then(function (res) {
          if (res.status == 200) {
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
          localStorage.setItem("token", res.data.token);
          if (res.data.type == "admin") {
            toaster.info("Comprobando Credenciales Admin...", { position: "top-right", duration: 3000, dismissible: true });
            UserService.userLoginAdmin(payload)
              .then(function (res) {
                setTimeout(function () {
                  toaster.success("Admin " + res.data.user.name.toUpperCase() + " loged successfully loged", { position: "top-right", duration: 5000, dismissible: true });
                  localStorage.setItem("token_admin", res.data.authorisation.token);
                  store.commit(Constant.SET_USER, res.data.user);
                }, 2000);
              })
              .catch(function (error) {
                console.log(error);
                toaster.error("Error login Admin", { position: "bottom", duration: 5000, dismissible: true });
              });
          } else {
            store.commit(Constant.SET_USER, res.data);
            toaster.success(res.data.name.toUpperCase() + " loged successfully", { position: "top-right", duration: 5000, dismissible: true });
          }
        })
        .catch(function (error) {
          if (error["response"]["data"] == "Email or password is not correct") {
            toaster.error("Email or password is not correct", { position: "bottom", duration: 5000, dismissible: true });
          } else {
            console.log(error);
            toaster.error("Error login", { position: "top-right", duration: 5000, dismissible: true });
          }
        });
    },
    [Constant.USER_REGISTER_ADMIN]: (store, payload) => {
      UserService.userRegisterAdmin(payload)
        .then(function (res) {
          if (res.status == 200) {
            store.commit(Constant.USER_REGISTER, payload);
          }
        })
        .catch(function (error) {
          if (error["response"]["data"]["errors"]["email"] || error["response"]["data"]["errors"]["name"]) {
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
      console.log("SOY EL LOGOUT");
      localStorage.removeItem("token");
      localStorage.removeItem("token_admin");
      store.commit(Constant.LOGOUT);
    },
    [Constant.USER_UPDATE]: (store, payload) => {
      let old_form = payload;
      let newForm = {};

      Object.keys(old_form).forEach(key => {
        if(old_form[key]) {
          newForm[key] = old_form[key];
        }
      });
      
      UserService.userUpdate(newForm)
      .then(function (res) {
        res.data.redirect = "no";
        store.commit(Constant.SET_USER, res.data);
        toaster.success("Updated correctly!", { position: "bottom", duration: 5000, dismissible: true });

      })
      .catch(function () {
          if (newForm.email){
            toaster.error("The email is already in use", { position: "bottom", duration: 5000, dismissible: true });
          }else {
            toaster.error("Error update User", { position: "top-right", duration: 5000, dismissible: true });
          }
         
      });
    },
  },
  getters: {
    getUser(state) {
      return state.user;
    },
    getAuth(state) {
      return state.user.type;
    },
    getNotis(state) {
      return state.user.noti;
    },
    setNotis(state) {
      state.user.noti = 0;
    }
  }
};
