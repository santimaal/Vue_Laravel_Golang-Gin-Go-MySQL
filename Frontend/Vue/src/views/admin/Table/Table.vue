<template>
    <h1>Admin Tablelist</h1>
    <router-link to="addtable"><button>Add</button></router-link>
    <TableItem_admin v-for="tableitem in state.tablelist" :key="tableitem.id" :tableitem="tableitem" />

</template>

<script>
import Constant from '../../../Constant';
import { reactive, computed } from 'vue'
import { useStore } from 'vuex'
import TableItem_admin from '../../../components/admin/TableItem.vue';

export default {
    setup() {
        const store = useStore();
        const state = reactive({
            tablelist: computed(() => store.getters["table/getTable"]),
        });
        store.dispatch("table/" + Constant.INITIALIZE_TABLE);
        return { state };
    },
    components: { TableItem_admin }
}
</script>

<style>

</style>