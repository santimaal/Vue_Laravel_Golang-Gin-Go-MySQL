<template>
    <form class="filters">
        Location
        <select name="" id="" v-model="state.filter.location">
            <option value="outside">Outside</option>
            <option value="inside">Inside</option>
        </select>
        <br>
        Capacity
        <input type="number" min="1" v-model="state.filter.capacity" />
        <br>
        Thematic
        <select v-model="state.filter.id_thematic" class="">
            <option v-for="(item, id) in state.thematic" :key="id" :value="item.id">{{ item.name }}</option>
        </select>
        <button type="button" @click="searchFilters">Search</button>
        <button type="button" @click="resetFilters">Reset</button>
    </form>

    <div class="row justify-content-center">
        <div class="col-8">
            <form class="row filters justify-content-around bg-light rounded p-3 m-2">
                <div class="col-md-3 position-relative">
                    <div class="d-flex justify-content-center">
                        <h3>Location</h3>
                    </div>
                    <p class="text-center">
                        Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do
                        eiusmod tempor incididunt ut labore et dolore magna aliqua.
                    </p>
                </div>
                <div class="col-md-3 position-relative">
                    <div class="d-flex justify-content-center">
                        <h3>Capacity</h3>
                    </div>
                    <p class="text-center">
                        Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do
                        eiusmod tempor incididunt ut labore et dolore magna aliqua.
                    </p>
                </div>
                <div class="col-md-3 position-relative">
                    <div class="d-flex justify-content-center">
                        <h3>Thematic</h3>
                    </div>
                    <p class="text-center">
                        Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do
                        eiusmod tempor incididunt ut labore et dolore magna aliqua.
                    </p>
                </div>
            </form>
        </div>
    </div>

</template>
<script>
import { reactive, getCurrentInstance, computed } from 'vue';
import { useRoute } from 'vue-router';
import { useStore } from 'vuex';
import Constant from '../../Constant';

export default {
    setup() {
        const store = useStore();
        const currentRoute = useRoute()
        const { emit } = getCurrentInstance();

        if (store.state.thematic.thematiclist.length == 0) {
            store.dispatch("thematic/" + Constant.INITIALIZE_THEMATIC)
        }

        const state = reactive({
            thematic: computed(() => store.getters['thematic/getThematic']),
            filter: {
                location: "", capacity: "", id_thematic: ""
            }
        })

        if (currentRoute.params.filter && currentRoute.params.filter != "all") {
            state.filter = JSON.parse(atob(currentRoute.params.filter))
        }

        const searchFilters = () => {
            emit('filters', state.filter)
        }

        const resetFilters = () => {
            state.filter = {
                location: "",
                capacity: "",
                id_thematic: ""
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