import Constant from "../../Constant";
import TableService from "@/services/TableService";


export const table = {
    state: {
        tablelist: []
    },
    namespaced: true,
    // state: {},
    mutations: {
        [Constant.ADD_TABLE]: (state, payload) => {
            console.log(state + " : " + payload);
            // state.tablelist.push({ ...payload });
        },
        [Constant.INITIALIZE_TABLE]: (state, payload) => {
            console.log("paco")
            // console.log(state + " : " + payload);

            // if (payload) {
            state.tablelist = payload;
            // } else {
            //     state.tablelist = {
            //         id_table: "",
            //         id_order: "",
            //         status: "",
            //     };
            // }
        }
    },
    actions: {
        [Constant.ADD_TABLE]: (store, payload) => {
            console.log(store + " : " + payload);
            // TableService.createTable(payload.tableitem)
            //     .then(function (res) {
            //         store.commit(Constant.ADD_TABLE, res.data.data);
            //     })
            //     .catch(function (error) {
            //         console.log(error);
            //     });
        },
        [Constant.INITIALIZE_TABLE]: (store /* payload */) => {
            TableService.getAllTable()
                .then(function (res) {
                    console.log(res.data.data);
                    store.commit(Constant.INITIALIZE_TABLE, res.data.data);
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
    },
    getters: {
        getTable(state) {
            return state.tablelist;
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