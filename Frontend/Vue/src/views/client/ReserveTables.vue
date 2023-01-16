<template>
  <search @search="ApplyFilters" />
  <filters @filters="ApplyFilters"></filters>

  <div class="change_color" v-if="state.show_tablelist.length > 0">
    <div class="all_cards_table">
      <Card_Table v-for="tableitem in state.show_tablelist" :key="tableitem.id" :tableitem="tableitem"
        v-on:click="checkuser" @click="state.clicked = tableitem.id, dateselected = 0, state.hours = []"
        data-toggle="modal" data-target="#reservetable" />
    </div>

    <Pagination :totalpages="state.total_pages" @changepage="loadnewtables" />

    <h1 v-if="state.reserve.length != 0">Reservadas</h1>
    <li v-for="table in state.reserve" :key="table.id">
      {{ table.id }}
      <button type="button" class="btn btn-primary" @click="deleteReserve(table.id)">delete</button>
    </li>
  </div>
  <div v-else>
    <img v-if="state.loadspinner" width="480" height="480" class="giphy-embed d-flex justify-content-center"
      src="https://media.tenor.com/dJLmV08Db0gAAAAi/liga-arroz.gif" alt="">
    <h1 v-if="!state.loadspinner">Don't have tables with filters applied</h1>
  </div>

  <!-- Modal -->
  <div class="modal fade" id="reservetable" tabindex="-1" role="dialog" aria-labelledby="modalLabel" aria-hidden="true"
    v-if="state.modal">
    <div class="modal-dialog" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="modalLabel">Quieres reservar La Mesa {{ state.clicked }}?</h5>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">&times;</span>
          </button>
        </div>
        <div class="modal-body">
          <input type="date" ref="fechaInput" v-model="dateselected" v-on:change="getHours(dateselected)" />
          <select name="" id="" v-model="hour">
            <option v-for="h in state.hours" :key="h" :value="h">{{ h }}</option>
          </select>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
          <button type="button" class="btn btn-primary" data-dismiss="modal"
            @click="addReserva(dateselected, hour, state.clicked)">Save
            changes</button>
        </div>
      </div>
    </div>
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
import { useGHours, useAddReserve } from "../../composables/reserve/useReserve";
import { useRoute, useRouter } from "vue-router";
import { createToaster } from "@meforma/vue-toaster";

export default {
  mounted() {
    const hoy = new Date()
    this.$refs.fechaInput.min = `${hoy.getFullYear()}-${String(hoy.getMonth() + 1).padStart(2, '0')}-${String(hoy.getDate() + 1).padStart(2, '0')}`

    setTimeout(() => {
      if (this.state.save_tablelist.length == 0) {
        this.state.loadspinner = false;
      }
    }, 2000);
  },
  setup() {
    const store = useStore();
    const router = useRouter();
    const currentRoute = useRoute();
    const toaster = createToaster();

    store.dispatch("table/" + Constant.INITIALIZE_TABLE);
    if (store.state.thematic.thematiclist.length == 0) {
      store.dispatch("thematic/" + Constant.INITIALIZE_THEMATIC);
    }

    const state = reactive({
      modal: true,
      hours: [],
      clicked: 0,
      show_tablelist: computed(() =>
        store.getters["table/getTable"].slice(0, 6)
      ),
      save_tablelist: computed(() => store.getters["table/getTable"]),
      reserve: [],
      page: 0,
      total_pages: computed(() =>
        Math.ceil(store.getters["table/getTable"].length / 6)
      ),
      loadspinner: true
    });

    setTimeout(() => {
      state.modal = false
    }, 1);

    state.total_pages = computed(() => Math.ceil(state.save_tablelist.length / 6))

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
      state.show_tablelist = computed(() => state.save_tablelist.slice(0, 6));
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

    const getHours = (date) => {
      let fecha = new Date(date)
      state.hours = useGHours({ dateini: fecha.toISOString(), id: state.clicked })
    }

    const checkuser = async () => {
      if (!localStorage.getItem('token')) {
        router.push({ name: 'login' })
      }
      if (await !store.getters["user/getUser"].is_active) {
        toaster.info("Usuario deshabilitado", { position: "top-right", duration: 5000, dismissible: true });
      } else {
        state.modal = true
      }
    }

    const addReserva = (date, hour, id) => {
      let time = "T"
      let newhour = parseInt(hour) - 1
      if (newhour == -1) {
        newhour = 23
      }
      if (newhour < 10) {
        time += "0"
      }
      useAddReserve({ dateini: date + time + newhour + ":00:00Z", id_table: id })
    }

    return { state, ApplyFilters, reserva, deleteReserve, loadnewtables, getHours, checkuser, addReserva };
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
  grid-template-columns: repeat(3, 19%);
  justify-content: center;
  justify-items: center;
  grid-gap: 2rem 3rem;
  margin-bottom: 2%;
  margin-top: 2%;
  background-color: rgb(201, 243, 234);
}

.change_color {
  background-color: rgb(201, 243, 234);
}
</style>
