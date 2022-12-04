import { ref } from 'vue';
import TableService from '../../services/TableService';

export const useTableFilters = (filters = {}) => {
    const table = ref([])
    // TableService.getTableFilter(filters)
    //     .then(res => table.value = res.data)
    //     .catch(error => console.error(error))
    let prueba = TableService.getTableFilter(filters)
    console.log(prueba);
    return table;
};