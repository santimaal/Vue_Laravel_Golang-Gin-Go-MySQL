<template>
    <h1>Client TableList</h1>
    <filters @filters="ApplyFilters"></filters>
    <!-- <li v-for="tableitem in state.tablelist" :key="tableitem.id">{{ tableitem }}</li> -->
    <TableItem_Client v-for="tableitem in state.tablelist" :key="tableitem.id" :tableitem="tableitem" />

</template>

<script>
import Constant from '../../../Constant';
import { reactive, computed } from 'vue'
import { useStore } from 'vuex'
import filters from "../../../components/client/Filters.vue"
import TableItem_Client from '../../../components/client/TableItem.vue';
import { useTableFilters } from '../../../composables/table/useTable';

export default {
    setup() {
        const store = useStore();

        store.dispatch("table/" + Constant.INITIALIZE_TABLE);
        if (store.state.thematic.thematiclist.length == 0) {
            store.dispatch("thematic/" + Constant.INITIALIZE_THEMATIC)
        }

        const state = reactive({
            tablelist: computed(() => store.getters["table/getTable"]),
        });

        const ApplyFilters = (filter) => {
            state.tablelist = useTableFilters(filter)
        }
        return { state, ApplyFilters };
    },
    components: { TableItem_Client, filters }
}
</script>

<style>

</style>