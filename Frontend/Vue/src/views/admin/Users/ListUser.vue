<template>
    <div class="all_list">

        <div class="d-flex justify-content-center mt-3">
            <table class="mytable table table-striped col-10 text-center bg-light">
                <thead class="thead-dark">
                    <tr>
                        <th scope="col">IMG</th>
                        <th scope="col">Name</th>
                        <th scope="col">Email</th>
                        <th scope="col">Type</th>
                        <th scope="col">Status</th>
                        <th scope="col">Options</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="data in state.users" :key="data.id">
                        <td scope="col"><img :src=data.img alt="avatar" class="img_profile rounded-circle img-fluid">
                        </td>
                        <td scope="col">{{ data.name }}</td>
                        <td scope="col">{{ data.email }}</td>
                        <td scope="col">{{ data.type }}</td>
                        <td scope="col">
                            <font-awesome-icon v-if="data.is_active==1" class="confirmed" icon="fa-solid fa-circle" />
                            <font-awesome-icon v-if="data.is_active==0" class="not_confirmed" icon="fa-solid fa-circle" /></td>
                        <td scope="col" class="text-center d-flex justify-content-around">
                            <button type="button" class="btn btn-outline-success border-radius"
                                @click="setUser('1', data.id)">V</button>
                            <button type="button" class="btn btn-outline-danger border-radius"
                                @click="setUser('0', data.id)">X</button>
                        </td>

                    </tr>
                </tbody>
            </table>
        </div>
    </div>

</template>
<script>
import { reactive, computed } from 'vue';
import { useChangeStatUser, useGetUserClient } from "../../../composables/user/UseUser"

export default {
    setup() {

        const state = reactive({
            cusers: computed(() => useGetUserClient()),
            users: []
        })

        setTimeout(() => {
            state.users = state.cusers
        }, 200);




        const setUser = async (status, id) => {
            await useChangeStatUser(status, id)
            state.users = state.users.map(item => {
                if (item.id == id) {
                    item.is_active = status;
                }
                return item;
            });
        }

        return { state, setUser }
    }
}
</script>
<style scoped>
.mytable {
    background-color: white;
}

.img_profile {
    width: 35px;
}

.all_list {
    height: 81vh;
}

.confirmed {
    color: green;
    font-size: x-large;
}

.not_confirmed {
    color: red;
    font-size: x-large;

}
</style>