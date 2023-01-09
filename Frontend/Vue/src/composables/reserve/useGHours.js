import { ref } from 'vue';
import ReserveService from '../../services/ReserveService';

export const useGHours = (date) => {
    const hours = ref([])
    ReserveService.getHours(date)
        .then(res =>{ 
            console.log(res); 
            hours.value = res.data
        })
        .catch(error => console.error(error))
    return hours;
};