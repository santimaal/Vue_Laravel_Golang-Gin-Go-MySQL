<template>
    <table class="table">
        <thead class="thead-dark">
            <tr>
                <th class="title" colspan="10">Dashboard Tables</th>
            </tr>
            <tr>
                <th scope="col">Id:</th>
                <th scope="col">Status Table</th>
                <th scope="col">Capacity</th>
                <th scope="col">Location</th>
                <th scope="col">Thematic</th>
                <th scope="col"><router-link to="addtable"><button
                            class="btn btn-primary m-1">Add</button></router-link></th>
            </tr>
        </thead>
        <tbody v-if="(state.tablelist.length > 0)">
            <TableItem_admin v-for="tableitem in state.tablelist" :key="tableitem.id" :tableitem="tableitem" />
        </tbody>
        <div v-if="(state.tablelist.length == 0)"> Don't have tables at this moment</div>
    </table>
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

<style>
/* thead tr {
    background-color: black !important;
    color: white;
} */
th {
    text-align: center;
}

.table {
    background-color: antiquewhite;
    width: 80%;
    float: right;
    margin-right: 2%;
    margin-top: 1%;
}

.title {
    background-color: black !important;
    font-size: xx-large;
    border-bottom: 2px solid lightblue !important;
}
</style>