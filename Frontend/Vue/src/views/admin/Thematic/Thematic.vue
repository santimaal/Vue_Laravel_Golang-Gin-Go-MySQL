<template>
    <h1>Admin ThematicList</h1>
    <!-- <router-link to="addthematic"><button>Add</button></router-link> -->
    <table class="table">
        <thead class="thead-dark">
            <tr>
                <th scope="col">ID:</th>
                <th scope="col">Name</th>
                <th scope="col">Location</th>
                <th scope="col">Img</th>
                <th scope="col"> <router-link to="addthematic"><button>Add</button></router-link></th>
            </tr>
        </thead>
        <tbody>
            <ThematicItem_admin v-for="thematicitem in state.thematiclist" :key="thematicitem.id"
                :thematicitem="thematicitem" />
        </tbody>
    </table>


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
th {
    text-align: center;
}
</style>