import Constant from "../../Constant";
import ThematicService from "@/services/ThematicService";


export const thematic = {
    state: {
        thematiclist: []
    },
    namespaced: true,
    mutations: {
        [Constant.ADD_THEMATIC]: (state, payload) => {
            state.thematiclist.push({ ...payload });
        },
        [Constant.DELETE_THEMATIC]: (state, payload) => {
            let index = state.thematiclist.findIndex(
                (item) => item.id === payload
            );
            state.thematiclist.splice(index, 1);
        },
        [Constant.UPDATE_THEMATIC]: (state, payload) => {
            let index = state.thematiclist.findIndex(
                item => item.id == payload.id
            );
            state.thematiclist[index] = payload;
        },
        [Constant.INITIALIZE_THEMATIC]: (state, payload) => {
            state.thematiclist = payload;
        }
    },
    actions: {
        [Constant.ADD_THEMATIC]: (store, payload) => {
            ThematicService.createThematic(payload.thematic)
                .then(function (res) {
                    store.commit(Constant.ADD_THEMATIC, res.data.data);
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        [Constant.DELETE_THEMATIC]: (store, payload) => {
            ThematicService.deleteThematicById(payload.id)
                .then(function (res) {
                    console.log(res);
                    store.commit(Constant.DELETE_THEMATIC, payload.id);
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        [Constant.UPDATE_THEMATIC]: (store, payload) => {
            console.log(payload.thematicitem.id);
            ThematicService.updateThematic(payload.thematicitem, payload.thematicitem.id)
                .then(function (res) {
                    console.log(res);
                    store.commit(Constant.UPDATE_THEMATIC, payload.thematicitem);
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        [Constant.INITIALIZE_THEMATIC]: (store /* payload */) => {
            ThematicService.getAllThematic()
                .then(function (res) {
                    // console.log(res.data);
                    store.commit(Constant.INITIALIZE_THEMATIC, res.data);
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
    },
    getters: {
        getThematic(state) {
            return state.thematiclist;
        },
        // getOrder(state) {
        //     if (state.tablelist) {
        //         var orders = state.tablelist.filter(function (element) {
        //             return element.order != null;
        //         }).map(function (element) {
        //             return element.order;

        //         })
        //         console.log(orders);
        //         return orders;
        //     }
        //     return;
        // },
    },
}