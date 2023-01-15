import { ref } from 'vue';
import UserService from '../../services/UserService';

export const useGetNotisAdmin = async () => {
    const noti = ref([])
    await UserService.getNotificationsAdmin()
        .then(res => {
            noti.value = res.data
        })
        .catch(error => console.error(error))
    return noti;
};

export const useGetNotisClient = async () => {
    const noti = ref([])
    await UserService.getNotificationsClient()
        .then(res => {
            res.data.forEach(noti => {
                let date = new Date(noti.dateini)
                noti.dateini = date.getDate() + '-' + (date.getMonth() + 1) + '-' + date.getFullYear() + " at " + date.getHours() + "h"
            });
            noti.value = res.data
        })
        .catch(error => console.error(error))
    return noti;
}

export const useSendNotification = async (id, message) => {
    await UserService.sendNotification(id, message)
        .then(res => {
            console.log(res);
        })
        .catch(error => console.error(error))
}