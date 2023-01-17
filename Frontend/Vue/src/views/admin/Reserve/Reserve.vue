<template>
    <div class="all_list">
        <!-- table-wrapper-scroll-y my-custom-scrollbar -->
        <div class="d-flex justify-content-center mt-3">
            <table class="mytable table table-striped col-10 text-center">
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
                        <td v-if="data.is_confirmed == 'accepted'"><font-awesome-icon class="text-success"
                                icon="fa-solid fa-check" />
                        </td>
                        <td v-if="data.is_confirmed == 'pending'"><font-awesome-icon class="text-info"
                                icon="fa-solid fa-arrows-rotate" /></td>
                        <td v-if="data.is_confirmed == 'denied'"><font-awesome-icon class="text-danger"
                                icon="fa-solid fa-x" /></td>
                        <td scope="col">{{ data.dateini }}</td>
                        <td scope="col">{{ data.datefin }}</td>
                        <td scope="col" class="text-center d-flex justify-content-around"
                            v-if="data.is_confirmed == 'pending'">
                            <button type="button" class="btn btn-outline-success border-radius"
                                @click="setReserve('accepted', data.id, data)">V</button>
                            <button type="button" class="btn btn-outline-danger border-radius"
                                @click="setReserve('denied', data.id, data)">X</button>
                        </td>
                        <td scope="col" class="text-center d-flex justify-content-around"
                            v-else>
                            <button disabled class="btn_noOptions">No options</button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>
<script>
import { reactive } from 'vue';
import { useStore } from 'vuex';
import { useGetReservesAdmin, useChangeStatReserve } from "../../../composables/reserve/useReserve"
import { useSendNotification } from '../../../composables/notifications/useNotifications'


export default {
    setup() {
        const store = useStore()
        const state = reactive({
            thematic: useGetReservesAdmin()
        })
        store.getters["user/setNotis"];

        const setReserve = async (status, id, noti) => {
            let date = new Date(noti.dateini)
            let message = date.getDate() + '-' + (date.getMonth() + 1) + '-' + date.getFullYear() + " at " + date.getHours() + "h " + status
            useSendNotification({ id_user: noti.id_user, message: message });
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
.my-custom-scrollbar {
    position: relative;
    height: 81vh;
    overflow: auto;
}

.table-wrapper-scroll-y {
    display: block;
}

.mytable {
    background-color: white;
}

.all_list {
    height: 81vh !important;
}

.btn_noOptions {
    background-color: aliceblue;
}
</style>