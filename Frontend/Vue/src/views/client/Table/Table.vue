<template>
    <h1>Client TableList</h1>
    <filters @filters="ApplyFilters"></filters>
    <!-- <li v-for="tableitem in state.tablelist" :key="tableitem.id">{{ tableitem }}</li> -->
    <TableItem_Client v-for="tableitem in state.tablelist" :key="tableitem.id" :tableitem="tableitem"
        @click="reserva(tableitem)" />

    <h1 v-if="(state.reserve.length != 0)">Reservadas</h1>
    <li v-for="table in state.reserve" :key="table.id">{{ table.id }} <button type="button"
            @click="deleteReserve(table.id)">delete</button></li>

</template>

<script>
import Constant from '../../../Constant';
import { reactive, computed } from 'vue'
import { useStore } from 'vuex'
import filters from "../../../components/client/Filters.vue"
import TableItem_Client from '../../../components/client/TableItem.vue';
import { useTableFilters } from '../../../composables/table/useTable';

export default {
    setup() {
        const store = useStore();

        store.dispatch("table/" + Constant.INITIALIZE_TABLE);
        if (store.state.thematic.thematiclist.length == 0) {
            store.dispatch("thematic/" + Constant.INITIALIZE_THEMATIC)
        }

        const state = reactive({
            tablelist: computed(() => store.getters["table/getTable"]),
            reserve: []
        });

        const ApplyFilters = (filter) => {
            state.tablelist = useTableFilters(filter)
        }

        const reserva = (table) => {
            let check = null
            if (state.reserve.length != 0) {
                check = state.reserve.find(item => {

                    return item.id == table.id;
                })
            }
            console.log(check);
            if (!check) {
                state.reserve.push(table)
            }
        }

        const deleteReserve = (table) => {
            let check = null
            state.reserve.find(item => {
                check++
                if (item.id == table) {
                    state.reserve.splice(check - 1, 1)
                }
                return item.id == table
            })
        }
        return { state, ApplyFilters, reserva, deleteReserve };
    },
    components: { TableItem_Client, filters }
}
</script>

<style>

</style>