import { ref } from 'vue';
import ReserveService from '../../services/ReserveService';

export const useGHours = (date) => {
    let available = Array.from(Array(24).keys())
    const hours = ref([])
    ReserveService.getHours(date)
        .then(res => {
            console.log(res);
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
            status.value = res.status
        })
        .catch(error => console.error(error))
        return status
};