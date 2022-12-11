import Constant from "../../Constant";
import TableService from "@/services/TableService";


export const table = {
    state: {
        tablelist: [],
    },
    namespaced: true,
    mutations: {
        [Constant.ADD_TABLE]: (state, payload) => {
            state.tablelist.push({ ...payload });
        },
        [Constant.DELETE_TABLE]: (state, payload) => {
            let index = state.tablelist.findIndex(
                (item) => item.id === payload
            );
            state.tablelist.splice(index, 1);
        },
        [Constant.UPDATE_TABLE]: (state, payload) => {
            let index = state.tablelist.findIndex(
                item => item.id == payload.id
            );
            state.tablelist[index] = payload;
        },
        [Constant.INITIALIZE_TABLE]: (state, payload) => {
            state.tablelist = payload;
        }
    },
    actions: {
        [Constant.ADD_TABLE]: async (store, payload) => {
            await TableService.createTable(payload.table)
                .then(function (res) {
                    console.log(res.data.data);
                    store.commit(Constant.ADD_TABLE, res.data.data);
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        [Constant.DELETE_TABLE]: async (store, payload) => {
            await TableService.deleteTableById(payload.id)
                .then(function (res) {
                    if (res.statusText !== "OK") {
                        throw Error("Ha habido algun problema al eliminar la mesa");
                    } else {
                        store.commit(Constant.DELETE_TABLE, payload.id);
                    }

                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        [Constant.UPDATE_TABLE]: async (store, payload) => {
            console.log(payload.tableitem.id);
            await TableService.updateTable(payload.tableitem, payload.tableitem.id)
                .then(function (res) {
                    console.log(res);
                    store.commit(Constant.UPDATE_TABLE, payload.tableitem);
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        [Constant.INITIALIZE_ONETABLE]: async (store, payload) => {
            console.log(payload);
            await TableService.getTableById(payload.id)
                .then(function (res) {
                    console.log(res.data);
                    return res.data
                   
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
        [Constant.INITIALIZE_TABLE]: async (store) => {
            await TableService.getAllTable()
                .then(function (res) {
                    store.commit(Constant.INITIALIZE_TABLE, res.data);
                })
                .catch(function (error) {
                    console.log(error);
                });
        },
    },
    getters: {
        getTable(state) {
            return state.tablelist;
        }
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