<template>
    <form class="filters">
        <select name="" id="" v-model="state.filter.location">
            <option value="outside">Outside</option>
            <option value="inside">Inside</option>
        </select>
        <br>
        <input type="number" min="1" v-model="state.filter.capacity" />
        <br>
        <select v-model="state.filter.id_thematic" class="">
            <option v-for="(item, id) in state.thematic" :key="id" :value="item.id">{{ item.name }}</option>
        </select>
        <button type="button" @click="searchFilters">Search</button>
        <button type="button" @click="resetFilters">Reset</button>
    </form>

</template>
<script>
import { reactive, getCurrentInstance, computed } from 'vue';
import { useStore } from 'vuex';
import Constant from '../../Constant';

export default {
    setup() {
        const store = useStore();
        const { emit } = getCurrentInstance();

        if (store.state.thematic.thematiclist.length == 0) {
            store.dispatch("thematic/" + Constant.INITIALIZE_THEMATIC)
        }

        const state = reactive({
            thematic: computed(() => store.getters['thematic/getThematic']),
            filter: {
                location: "", capacity: 1, id_thematic: ""
            }
        })

        const searchFilters = () => {
            emit('filters', state.filter)
        }

        const resetFilters = () => {
            state.filter= {
                location : "",
                capacity : 1,
                id_thematic:""
            }
        }

        return { state, searchFilters, resetFilters }
    }
}
</script>
<style>
.filters {
    margin: 2% 2%;
}
</style>