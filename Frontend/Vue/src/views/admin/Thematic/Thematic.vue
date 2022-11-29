<template>
    <h1>Admin ThematicList</h1>
    <router-link to="addthematic"><button>Add</button></router-link>
    <ThematicItem_admin v-for="thematicitem in state.thematiclist" :key="thematicitem.id" :thematicitem="thematicitem" />

</template>

<script>
import Constant from '../../../Constant';
import { reactive, computed } from 'vue'
import { useStore } from 'vuex'
import ThematicItem_admin from '../../../components/admin/ThematicItem.vue';

export default {
    setup() {
        const store = useStore();

        store.dispatch("thematic/" + Constant.INITIALIZE_THEMATIC);
        
        const state = reactive({
            thematiclist: computed(() => store.getters["thematic/getTable"]),
        });
        return { state };
    },
    components: { ThematicItem_admin }
}
</script>

<style>

</style>