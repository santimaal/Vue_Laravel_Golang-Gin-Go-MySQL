import { ref } from 'vue';
import { createToaster } from "@meforma/vue-toaster";
import UserService from '../../services/UserService';

const toaster = createToaster();

export const useChangeStatUser = (status, id) => {
    UserService.updateStateUser(status, id)
        .then(res => {
            console.log(res);
            toaster.success("Change status User Correctly", { position: "top-right", duration: 5000, dismissible: true });
        })
        .catch(error => console.error(error))
}

export const useGetUserClient = () => {
    const users = ref([])
    UserService.getUsersAdmin()
        .then(res => {
            users.value = res.data
        })
        .catch(error => console.error(error))
    return users
}
