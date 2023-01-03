import Constant from "../../Constant";
import UserService from "../../services/UserService";

export const user = {
    state: {
        user: {
            "is_active": false,
            "name": "",
            "email": "",
            "img": "",
            "type": "client"
        }
    },
    namespaced: true,
    mutations: {
        [Constant.SET_USER]: (state, payload) => {
            console.log(payload);
            state.user = {
                "is_active": payload.is_active,
                "name": payload.name,
                "email": payload.email,
                "img": payload.img,
                "type": payload.type
            }
        },
    },
    actions: {
        [Constant.USER_REGISTER]: (store, payload) => {
            UserService.userRegister(payload)
                .then(function (res) {
                    console.log(res);
                    // store.commit(Constant.ADD_THEMATIC, res.data.data);
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        [Constant.USER_LOGIN]: (store, payload) => {
            UserService.userLogin(payload)
                .then(function (res) {
                    console.log(res.data);
                    store.commit(Constant.SET_USER, res.data);
                    localStorage.setItem("token", JSON.stringify(res.data.token))
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        [Constant.USER_REGISTER_ADMIN]: (store, payload) => {
            UserService.userRegisterAdmin(payload)
                .then(function (res) {
                    console.log(res);
                    // store.commit(Constant.ADD_THEMATIC, res.data.data);
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
    },
    getters: {
        getUser(state) {
            return state.user;
        }
    },
}