<template>
    <div class="all_list">
        <div class="d-flex justify-content-center align-items-center">
            <div class="d-flex flex-column mt-5">
                <div class="div_border">
                    <h1 class="text-center mt-2 title">Update Table</h1>
                    <div class="pl-5 pr-5 pb-4 pt-2">
                        <div class="form-group">
                            <label>Status Table:</label>
                            <select v-model="state.tableitemlocal.is_active" class="m-3">
                                <option value="true">Reserved</option>
                                <option value="false">No Reserved</option>
                            </select>
                        </div>
                        <div class="form-group">
                            <label>Capacity :</label>
                            <input type="number" class="m-3 .col-md-1" v-model="state.tableitemlocal.capacity" />
                        </div>
                        <div class="form-group">
                            <label>Location :</label>
                            <select v-model="state.tableitemlocal.location" class="m-3">
                                <option value="outside">Outside</option>
                                <option value="inside">Inside</option>
                            </select>
                        </div>
                        <div class="form-group">
                            <label>Thematic: </label>
                            <select v-model="state.tableitemlocal.id_thematic" class="m-3">
                                <option v-for="(item, id) in state.thematiclist" :key="id" :value="item.id">{{
                                    item.name
                                }}
                                </option>
                            </select>
                        </div>
                        <div class="d-flex justify-content-center">
                            <button type="button" class="btn btn-outline-success mr-2" @click="updateTable">Update</button>
                            <router-link to="/atable"><button class="btn btn-outline-danger ml-2">Cancel</button></router-link>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script>
import Constant from '../../../Constant';
import { reactive } from 'vue'
import { useStore } from 'vuex'
import { useRouter, useRoute } from 'vue-router';
import TableService from "@/services/TableService";
import ThematicService from "@/services/ThematicService";

export default {
    setup() {
        const store = useStore();
        const router = useRouter();
        const currentRoute = useRoute();
        const state = reactive({
            tableitemlocal: {},
            thematiclist: store.state.thematic.thematiclist
        });

        if (store.state.table.tablelist.length != 0) {
            state.tableitemlocal = store.state.table.tablelist.find(item => item.id == currentRoute.params.id);
        } else {
            TableService.getTableById(currentRoute.params.id).then(data => {
                state.tableitemlocal = data.data;
            })
            ThematicService.getAllThematic().then(data => {
                state.thematiclist = data.data;
            })
        }

        const updateTable = () => {
            state.tableitemlocal.is_active = state.tableitemlocal.is_active == 'false' ? false : true;
            router.push({ name: "admin_table" });
            store.dispatch("table/" + Constant.UPDATE_TABLE, { tableitem: state.tableitemlocal });
        }

        return { state, updateTable };
    }
}
</script>
  
<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Lobster&display=swap");

.all_list {
    height: 81vh !important;
}

.div_border {
    border: 2px solid black;
    background-color: white;
    border-radius: 5px;
}

.title {
    text-decoration: underline;
    font-family: "Lobster", cursive;
}
</style>