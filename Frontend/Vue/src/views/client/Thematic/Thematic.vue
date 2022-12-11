<template>
    <h1>Client ThematicList</h1>
    <ThematicItem_client v-for="thematicitem in state.thematiclist" :key="thematicitem.id" :thematicitem="thematicitem" />

</template>

<script>
import Constant from '../../../Constant';
import { reactive, computed } from 'vue'
import { useStore } from 'vuex'
import ThematicItem_client from '../../../components/client/Card_Thematic.vue';

export default {
    setup() {
        const store = useStore();

        store.dispatch("thematic/" + Constant.INITIALIZE_THEMATIC);
        
        const state = reactive({
            thematiclist: computed(() => store.getters["thematic/getThematic"]),
        });
        return { state };
    },
    components: { ThematicItem_client }
}
</script>

<style>

</style>