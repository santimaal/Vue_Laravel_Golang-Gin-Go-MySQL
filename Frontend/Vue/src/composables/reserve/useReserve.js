import { ref } from 'vue';
import ReserveService from '../../services/ReserveService';
import { createToaster } from "@meforma/vue-toaster";
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
    ReserveService.updateReserve(status,id)
    .then(res => {
        console.log(res);
    })
    .catch(error => console.error(error)) 
}