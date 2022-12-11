import { ref } from "vue";
import ThematicService from '../../services/ThematicService'

export const useThematicInfinite = (page = 1, limit = 6) => {
    const thematics = ref([])
    ThematicService.GetThematicInfinite(page, limit)
        .then(res => [thematics.value = res.data])
        .catch(error => console.error(error))
    return thematics;
};


