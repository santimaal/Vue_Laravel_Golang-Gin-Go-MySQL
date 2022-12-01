<template>
    <div class="m-3">
        <h1>Create Table</h1>
        <div>
            <label>Thematic :</label>
            <select v-model="state.table.id_thematic" class="m-3">
                <option v-for="(item, id) in state.thematiclist" :key="id" :value="item.id">{{ item.name }}</option>
            </select>
        </div><br>
        <div>
            <label>Status Table :</label>
            <select v-model="state.table.is_active" class="m-3 mt-1">
                <option value=true selected>True</option>
                <option value=false>False</option>
            </select>
        </div><br>
        <div>
            <label>Capacity :</label>
            <input class="m-3" type="number" min="1" v-model="state.table.capacity" />
        </div><br>
        <div>
            <label class="mt-2">Location:</label>
            <select class="m-3" v-model="state.table.location">
                <option value="outside" selected>Outside</option>
                <option value="inside">Inside</option>
            </select>
        </div><br>
        <button class="btn btn-primary m-1" type="button" @click="addTable">Add</button>
        <router-link to="atable"><button class="btn btn-primary m-1">Cancel</button></router-link>
    </div>
</template>

<script>
import Constant from '../../../Constant';
import { reactive } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';

export default {
    setup() {
        const store = useStore();
        const router = useRouter()
        const state = reactive({
            table: { id_thematic: store.state.thematic.thematiclist[0].id, is_active: false, capacity: 1, location: "outside" },
            thematiclist: store.state.thematic.thematiclist
        });


        const addTable = () => {
            try {
                if (state.table.is_active.indexOf("'")) {
                    console.log("sii");
                    state.table.is_active = state.table.is_active == 'false' ? false : true
                }
            } catch {
                console.log("all good");
            }
            store.dispatch("table/" + Constant.ADD_TABLE, { table: state.table })
            router.push({ name: "admin_table" });
        }


        return { state, addTable }
    }
}

</script>

<style>

</style>