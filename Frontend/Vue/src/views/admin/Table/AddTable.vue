<template>
    <div class="all_list">
        <div class="d-flex justify-content-center align-items-center">
            <div class="d-flex flex-column mt-5">
                <div class="div_border">
                    <h1 class="text-center mt-2 title">Create Table</h1>
                    <div class="pl-5 pr-5 pb-4 pt-2">
                        <div>
                            <label>Thematic :</label>
                            <select v-model="state.table.id_thematic" class="m-3">
                                <option v-for="(item, id) in state.thematiclist" :key="id" :value="item.id">{{
                                    item.name
                                }}
                                </option>
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
                        <div class="d-flex justify-content-center">
                            <button class="btn btn-outline-success mr-2" type="button" @click="addTable">Add</button>
                            <router-link to="atable"><button class="btn btn-outline-danger ml-2">Cancel</button></router-link>
                        </div>
                    </div>
                </div>
            </div>
        </div>
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
                    state.table.is_active = state.table.is_active == 'false' ? false : true
                }
            } catch {
                console.log("error");
            }
            store.dispatch("table/" + Constant.ADD_TABLE, { table: state.table })
            router.push({ name: "admin_table" });
        }


        return { state, addTable }
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

.title{
    text-decoration: underline;
    font-family: "Lobster", cursive;
}
</style>