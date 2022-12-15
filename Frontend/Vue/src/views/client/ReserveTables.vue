<template>
  <search @search="ApplyFilters" />
  <filters @filters="ApplyFilters"></filters>

  <div v-if="state.show_tablelist.length > 0">
    <div class="all_cards_table">
      <Card_Table v-for="tableitem in state.show_tablelist" :key="tableitem.id" :tableitem="tableitem"
        @click="reserva(tableitem)" />
    </div>

    <Pagination :totalpages="state.total_pages" @changepage="loadnewtables" />

    <h1 v-if="state.reserve.length != 0">Reservadas</h1>
    <li v-for="table in state.reserve" :key="table.id">
      {{ table.id }}
      <button type="button" class="btn btn-primary" @click="deleteReserve(table.id)">delete</button>
    </li>
  </div>
  <div v-else>
    <!-- <v-icon name="la-spinner-solid" animation="spin" fill="black" scale="4" /> -->
    <h1>Don't have tables with filters applied</h1>
  </div>
</template>

<script>
import Constant from "../../Constant";
import { reactive, computed } from "vue";
import { useStore } from "vuex";
import filters from "../../components/client/Filters.vue";
import Card_Table from "../../components/client/Card_Table.vue";
import Pagination from "../../components/client/Pagination.vue";
import search from "../../components/client/Search.vue";
import { useTableFilters } from "../../composables/table/useFilters";
import { useRoute, useRouter } from "vue-router";

export default {
  setup() {
    const store = useStore();
    const router = useRouter();
    const currentRoute = useRoute();

    store.dispatch("table/" + Constant.INITIALIZE_TABLE);
    if (store.state.thematic.thematiclist.length == 0) {
      store.dispatch("thematic/" + Constant.INITIALIZE_THEMATIC);
    }

    const state = reactive({
      show_tablelist: computed(() =>
        store.getters["table/getTable"].slice(0, 6)
      ),
      save_tablelist: computed(() => store.getters["table/getTable"]),
      reserve: [],
      page: 0,
      total_pages: computed(() =>
        Math.ceil(store.getters["table/getTable"].length / 6)
      )
    });

    const ApplyFilters = (filter) => {
      if (
        currentRoute.params.filter != "all" &&
        filter.search == undefined &&
        JSON.parse(atob(currentRoute.params.filter)).search != undefined
      ) {
        filter = {
          ...filter,
          search: JSON.parse(atob(currentRoute.params.filter)).search,
        };
      }

      filter = { ...filter, offset: state.page * 6 };
      const urlfilter = btoa(JSON.stringify(filter));
      router.push({ name: "client_table", params: { filter: urlfilter } });
      state.save_tablelist = useTableFilters(filter);
      setTimeout(() => {
        state.show_tablelist = computed(() => state.save_tablelist.slice(0, 6));
        state.total_pages = computed(() =>
          Math.ceil(state.save_tablelist.length / 6)
        );
      }, 100);
      state.page = 0;
    };

    if (currentRoute.params.filter && currentRoute.params.filter != "all") {
      ApplyFilters(JSON.parse(atob(currentRoute.params.filter)));
    }

    const loadnewtables = (page) => {
      state.show_tablelist = computed(() =>
        state.save_tablelist.slice(page * 6, page * 6 + 6)
      );
    }

    const reserva = (table) => {
      let check = null;
      if (state.reserve.length != 0) {
        check = state.reserve.find((item) => {
          return item.id == table.id;
        });
      }
      if (!check) {
        state.reserve.push(table);
      }
    };

    const deleteReserve = (table) => {
      let check = null;
      state.reserve.find((item) => {
        check++;
        if (item.id == table) {
          state.reserve.splice(check - 1, 1);
        }
        return item.id == table;
      });
    };
    return { state, ApplyFilters, reserva, deleteReserve, loadnewtables };
  },
  components: { Card_Table, filters, search, Pagination },
};
</script>

<style>
.pagination {
  background-color: rgb(201, 243, 234);
  border-radius: 0px;
}

.all_cards_table {
    display: grid;
    grid-template-columns: repeat(3, 22%);
    justify-content: center;
    justify-items: center;
    grid-gap: 2rem 2rem;
    margin-bottom: 2%;
    margin-top: 2%;
    background-color: rgb(201, 243, 234);
}

.change_color {
  background-color: rgb(201, 243, 234);
}
</style>
