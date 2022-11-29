<template>
    <h1>Client TableList</h1>
    <TableItem_Client v-for="tableitem in state.tablelist" :key="tableitem.id" :tableitem="tableitem" />

</template>

<script>
import Constant from '../../../Constant';
import { reactive, computed } from 'vue'
import { useStore } from 'vuex'
import TableItem_Client from '../../../components/client/TableItem.vue';

export default {
    setup() {
        const store = useStore();

        store.dispatch("table/" + Constant.INITIALIZE_TABLE);

        const state = reactive({
            tablelist: computed(() => store.getters["table/getTable"]),
        });
        return { state };
    },
    components: { TableItem_Client }
}
</script>

<style>

</style>