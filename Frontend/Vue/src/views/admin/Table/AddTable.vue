<template>
    <h1>Create Table</h1>
    Id Thematic <input type="number" min="1" v-model="state.table.id_thematic" />
    <br><br>
    Active <select name="" id="" v-model="state.table.is_active">
        <option value=true selected>True</option>
        <option value=false>False</option>
    </select>
    <br><br>
    Capacity<input type="number" min="1" v-model="state.table.capacity" />
    <br><br>
    Location <select name="" id="" v-model="state.table.location">
        <option value="exterior" selected>Exterior</option>
        <option value="interior">Interior</option>
    </select>
    <br><br>
    <button type="button" @click="addTable">Add</button>
    <router-link to="atable"><button>Cancel</button></router-link>
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
            table: { id_thematic: 1, is_active: false, capacity: 1, location: "exterior" }
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