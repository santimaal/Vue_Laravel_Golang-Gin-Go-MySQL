<template>
    <div class="m-4">
        <div class="row">
            <div class="col p-3">
                <h2>Update Table</h2>
            </div>
        </div>
        <div class="row">
            <div class="col">
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
                    <label>Thematic : </label>
                    <select v-model="state.tableitemlocal.id_thematic" class="m-3">
                        <option v-for="(item,id) in state.thematiclist" :key="id" :value="item.id">{{item.name}}</option>
                    </select>
                </div>
                <div class="form-group">
                    <button type="button" class="btn btn-primary m-1" @click="updateTable">Update</button>
                    <router-link to="/atable"><button class="btn btn-primary m-1">Cancel</button></router-link>
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

export default {
    setup() {
        const store = useStore();
        const router = useRouter();
        const currentRoute = useRoute();
        const tableitem = store.state.table.tablelist.find(item => item.id == currentRoute.params.id);
        const state = reactive({
            tableitemlocal: { ...tableitem },
            thematiclist: store.state.thematic.thematiclist
        });

        console.log(state.tableitemlocal.id_thematic);

        const updateTable = () => {
            router.push({ name: "todoList" });
            store.dispatch(Constant.UPDATE_TODO, { todoitem: state.tablelist });
        }

        return { state,updateTable };
    }
}
</script>
  
<style>

</style>