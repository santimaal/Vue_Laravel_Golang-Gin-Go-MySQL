<template>
    <search @search="ApplyFilters" />
    <h1>Client TableList</h1>
    <filters @filters="ApplyFilters"></filters>
    <!-- <li v-for="tableitem in state.tablelist" :key="tableitem.id">{{ tableitem }}</li> -->
    <TableItem_Client v-for="tableitem in state.show_tablelist" :key="tableitem.id" :tableitem="tableitem"
        @click="reserva(tableitem)" />

    <nav aria-label="">
        <ul class="pagination">
            <li class="page-item">
                <a class="page-link" href="#" aria-label="Previous">
                    <span aria-hidden="true">&laquo;</span>
                    <span class="sr-only">Previous</span>
                </a>
            </li>
            <li class="page-item" v-for="row, id in state.total_pages" :key="id" :class="isActive(id)"
                @click="changePage(id)"><a class="page-link" href="#">{{ row }}</a>
            </li>
            <!-- <li class="page-item" @click="(state.page = 0)"><a class="page-link" href="#">1</a></li>
            <li class="page-item"><a class="page-link" href="#">2</a></li>
            <li class="page-item"><a class="page-link" href="#">3</a></li> -->
            <li class="page-item">
                <a class="page-link" href="#" aria-label="Next">
                    <span aria-hidden="true">&raquo;</span>
                    <span class="sr-only">Next</span>
                </a>
            </li>
        </ul>
    </nav>

    <h1 v-if="(state.reserve.length != 0)">Reservadas</h1>
    <li v-for="table in state.reserve" :key="table.id">
        {{ table.id }}
        <button type="button" @click="deleteReserve(table.id)">delete</button>
    </li>
</template>

<script>
import Constant from '../../Constant';
import { reactive, computed } from 'vue'
import { useStore } from 'vuex'
import filters from "../../components/client/Filters.vue"
import TableItem_Client from '../../components/client/TableItem.vue';
import search from '../../components/client/Search.vue';
import { useTableFilters } from '../../composables/table/useFilters';
import { useRoute, useRouter } from 'vue-router';

export default {
    setup() {
        const store = useStore();
        const router = useRouter();
        const currentRoute = useRoute();



        store.dispatch("table/" + Constant.INITIALIZE_TABLE);
        if (store.state.thematic.thematiclist.length == 0) {
            store.dispatch("thematic/" + Constant.INITIALIZE_THEMATIC)
        }

        const state = reactive({
            show_tablelist: computed(() => store.getters["table/getTable"].slice(0, 6)),
            save_tablelist: computed(() => store.getters["table/getTable"]),
            reserve: [],
            page: 0,
            total_pages: computed(() => Math.ceil(store.getters["table/getTable"].length / 6))
        });

        const ApplyFilters = (filter) => {
            if (currentRoute.params.filter != "all" && filter.search == undefined && JSON.parse(atob(currentRoute.params.filter)).search != undefined) {
                filter = { ...filter, search: JSON.parse(atob(currentRoute.params.filter)).search }
            }

            filter = { ...filter, offset: state.page * 6 }
            console.log(filter);
            const urlfilter = btoa(JSON.stringify(filter));
            router.push({ name: "client_table", params: { filter: urlfilter } });
            state.save_tablelist = useTableFilters(filter)
            setTimeout(() => {
                state.show_tablelist = computed(() => state.save_tablelist.slice(0, 6))
                state.total_pages = computed(() => Math.ceil(state.save_tablelist.length / 6))
            }, 100);
            state.page = 0
        }

        if (currentRoute.params.filter && currentRoute.params.filter != "all") {
            // state.show_tablelist = useTableFilters(JSON.parse(atob(currentRoute.params.filter)))
            ApplyFilters(JSON.parse(atob(currentRoute.params.filter)))
        }

        const isActive = (id)=> {
            return {active: id == state.page}
        }

        const changePage = (page) => {
            // console.log(test);
            state.page = page
            console.log(state.page);
            console.log(state.save_tablelist.slice(page * 6, page * 6 + 6));
            state.show_tablelist = computed(() => state.save_tablelist.slice(page * 6, page * 6 + 6))
            // ApplyFilters({ ...JSON.parse(atob(currentRoute.params.filter)), offset: page * 6 })
            // const urlfilter = JSON.parse(atob(currentRoute.params.filter));
            // state.tablelist = useTableFilters({ ...urlfilter, offset: page * 6 })
        }

        const reserva = (table) => {
            let check = null
            if (state.reserve.length != 0) {
                check = state.reserve.find(item => {

                    return item.id == table.id;
                })
            }
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
        return { state, ApplyFilters, reserva, deleteReserve, changePage, isActive };
    },
    components: { TableItem_Client, filters, search }
}
</script>

<style>

</style>