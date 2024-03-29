import { ref } from 'vue';
import ReserveService from '../../services/ReserveService';
import { createToaster } from "@meforma/vue-toaster";
import UserService from '../../services/UserService';

const toaster = createToaster();

export const useGHours = (date) => {
    let available = Array.from(Array(24).keys())
    const hours = ref([])
    ReserveService.getHours(date)
        .then(res => {
            res.data.map(reserva => {
                available.splice(available.indexOf(parseInt(reserva.dateini.substring(11, 13))), 1)
            })
            hours.value = available
        })
        .catch(error => console.error(error))
    return hours;
};

export const useAddReserve = (info) => {
    const status = ref()
    ReserveService.addReserve(info)
        .then(res => {
            if (res.status == 200) {
                let date = new Date(info.dateini);
                toaster.success("Reserve created at <br>" + date.getDate() + '-' + (date.getMonth() + 1) + '-' + date.getFullYear() + " at " + date.getHours() + "h", { position: "top-right", duration: 10000, dismissible: true });
            }
            status.value = res.status
        })
        .catch(error => console.error(error))
    return status
};

export const useChangeStatReserve = (status, id) => {
    ReserveService.updateReserve(status, id)
        .then(res => {
            console.log(res);
        })
        .catch(error => console.error(error))
}

export const useGetReservesAdmin = () => {
    const reserves = ref([])
    ReserveService.getReservesAdmin()
        .then(res => {
            reserves.value = res.data
        })
        .catch(error => console.error(error))
    return reserves
}


export const useGetReserveClient = async () => {
        const reserva = ref([])
        await UserService.getMyReserves()
            .then(res => {
                res.data.forEach(reserva => {
                    let date = new Date(reserva.dateini)
                    reserva.dateini= date.getDate() + '-' + (date.getMonth() + 1) + '-' + date.getFullYear() + " at " + date.getHours() + "h"
                });
                reserva.value = res.data
            })
            .catch(error => console.error(error))
        return reserva;
}