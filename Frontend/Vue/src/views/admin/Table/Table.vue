<template>
    <div class="all_list">
        <div class="d-flex justify-content-center mt-3">
            <table class="table table-striped col-10 text-center">
                <thead class="thead-dark">
                    <tr>
                        <th scope="col">Id:</th>
                        <th scope="col">Status Table</th>
                        <th scope="col">Capacity</th>
                        <th scope="col">Location</th>
                        <th scope="col">Thematic</th>
                        <th scope="col"><router-link to="addtable"><button class="btn btn-outline-warning">Add</button></router-link></th>
                    </tr>
                </thead>
                <tbody v-if="(state.tablelist.length > 0)">
                    <TableItem_admin v-for="tableitem in state.tablelist" :key="tableitem.id" :tableitem="tableitem" />
                </tbody>
                <div v-if="(state.tablelist.length == 0)"> Don't have tables at this moment</div>
            </table>
        </div>
    </div>
</template>


<script>
import Constant from '../../../Constant';
import { reactive, computed } from 'vue'
import { useStore } from 'vuex'
import TableItem_admin from '../../../components/admin/TableItem.vue';

export default {
    setup() {
        const store = useStore();

        store.dispatch("table/" + Constant.INITIALIZE_TABLE);
        store.dispatch("thematic/" + Constant.INITIALIZE_THEMATIC);

        const state = reactive({
            tablelist: computed(() => store.getters["table/getTable"]),
        });
        return { state };
    },
    components: { TableItem_admin }
}
</script>

<style scoped>
th {
    text-align: center;
}

.table {
    background-color: white;
}

.title {
    background-color: black !important;
    font-size: xx-large;
    border-bottom: 2px solid lightblue !important;
}
.all_list {
    height: 81vh !important;
}
</style>