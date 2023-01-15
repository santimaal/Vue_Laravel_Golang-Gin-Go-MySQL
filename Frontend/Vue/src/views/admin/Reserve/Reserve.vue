<template>
    <table class="table table-striped">
        <thead class="thead-dark">
            <tr>
                <th scope="col">#</th>
                <th scope="col">Id_Table</th>
                <th scope="col">Id_User</th>
                <th scope="col">Is_Confirmed</th>
                <th scope="col">Dateini</th>
                <th scope="col">Datefin</th>
                <th scope="col" class="text-center">Options</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="data in state.thematic" :key="data.id">
                <td scope="col">{{ data.id }}</td>
                <td scope="col">{{ data.id_table }}</td>
                <td scope="col">{{ data.id_user }}</td>
                <td scope="col">{{ data.is_confirmed }}</td>
                <td scope="col">{{ data.dateini }}</td>
                <td scope="col">{{ data.datefin }}</td>
                <td scope="col" class="text-center d-flex justify-content-around" v-if="data.is_confirmed == 'pending'">
                    <button type="button" class="btn btn-outline-success border-radius"
                        @click="setReserve('accepted', data.id)">V</button>
                    <button type="button" class="btn btn-outline-danger border-radius"
                        @click="setReserve('denied', data.id)">X</button>
                </td>

            </tr>
        </tbody>
    </table>
</template>
<script>
import { reactive } from 'vue';
import { useStore } from 'vuex';
import { useGetReservesAdmin, useChangeStatReserve } from "../../../composables/reserve/useReserve"

export default {
    setup() {
        const store = useStore()
        const state = reactive({
            thematic: useGetReservesAdmin()
        })
        store.getters["user/setNotis"];

        const setReserve = async (status, id) => {
            await useChangeStatReserve(status, id)
            state.thematic = state.thematic.map(item => {
                if (item.id === id) {
                    item.is_confirmed = status;
                }
                return item;
            });
        }

        // useGetReservesAdmin())
        return { state, setReserve }
    }
}
</script>
<style scoped>

</style>