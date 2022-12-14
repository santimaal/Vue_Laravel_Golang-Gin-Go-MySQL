import { ref } from 'vue';
import UserService from '../../services/UserService';

export const useGetNotisAdmin = async() => {
    const noti = ref([])
    await UserService.getNotificationsAdmin()
        .then(res => {
            noti.value = res.data
        })
        .catch(error => console.error(error))
    return noti;
};

export const useGetNotisClient = async() => {
    const noti = ref([])
    await UserService.getNotificationsClient()
        .then(res => {
            noti.value = res.data
        })
        .catch(error => console.error(error))
    return noti;
}