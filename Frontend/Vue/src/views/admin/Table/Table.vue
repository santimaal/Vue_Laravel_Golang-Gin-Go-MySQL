<template>
    <div class="template">
        <h1>Admin Tablelist</h1>
        <router-link to="addtable"><button class="btn btn-primary m-1">Add</button></router-link>
        <TableItem_admin v-for="tableitem in state.tablelist" :key="tableitem.id" :tableitem="tableitem" />
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

<style>
.template {
    padding: 3%;
}
</style>