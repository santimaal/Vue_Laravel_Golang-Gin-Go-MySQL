import Constant from "../../Constant";
import UserService from "../../services/UserService";

export const user = {
    state: {
        user: {
            "name": "",
            "email": "",
            "img": "",
        }
    },
    namespaced: true,
    mutations: {
        // [Constant.ADD_THEMATIC]: (state, payload) => {
        //     state.thematiclist.push({ ...payload });
        // },
        // [Constant.DELETE_THEMATIC]: (state, payload) => {
        //     let index = state.thematiclist.findIndex(
        //         (item) => item.id === payload
        //     );
        //     state.thematiclist.splice(index, 1);
        // },
        // [Constant.UPDATE_THEMATIC]: (state, payload) => {
        //     let index = state.thematiclist.findIndex(
        //         item => item.id == payload.id
        //     );
        //     state.thematiclist[index] = payload;
        // },
        // [Constant.INITIALIZE_THEMATIC]: (state, payload) => {
        //     state.thematiclist = payload;
        // }
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
    },
    getters: {
        getUser(state) {
            return state.user;
        }
    },
}