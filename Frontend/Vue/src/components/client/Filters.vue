<template>
    <div class="row justify-content-center">
        <div class="col-8">
            <form class="filters justify-content-around rounded p-3 m-2">
                <div class="row justify-content-around">
                    <div class="col-md-3 position-relative">
                        <div class="d-flex justify-content-center">
                            <select name="" id="" v-model="state.filter.location" class="d-flex justify-content-center">
                                <option class="fil_center" value="" selected disabled hidden> LOCATION</option>
                                <option class="fil_center" value="outside">Outside</option>
                                <option class="fil_center" value="inside">Inside</option>
                            </select>
                        </div>
                    </div>
                    <div class="col-md-3 position-relative">
                        <div class="d-flex justify-content-center">
                            <input class="fil_center" type="number" min="1" placeholder="CAPACITY" v-model="state.filter.capacity" />
                        </div>
                    </div>
                    <div class="col-md-3 position-relative">
                        <div class="d-flex justify-content-center">
                            <select v-model="state.filter.id_thematic" class="">
                                <option class="fil_center" value="" selected disabled hidden>THMATIC</option>
                                <option class="fil_center" v-for="(item, id) in state.thematic" :key="id" :value="item.id">{{ item.name }}
                                </option>
                            </select>
                        </div>
                    </div>
                </div>
                <div class="row justify-content-center mt-3 filbuton">
                    <button type="button" class="mr-3" @click="searchFilters"><font-awesome-icon icon="fa-solid fa-magnifying-glass" /></button>
                    <button type="button" @click="resetFilters"><font-awesome-icon icon="fa-solid fa-arrows-rotate" /></button>
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
    margin-top: 5% !important;
}

select, input {
  height: 40px;
  width: 200px;
  display: flex;
  justify-content: center;
  border-radius: 30px;
  border: solid 2px black;
  background-color: white;
}
.fil_center {
    display: flex !important;
    text-align: center !important;
    border-radius: 30px !important;
}
input::placeholder {
  color: black;
}
.filbuton button{
    background-color: black;
    color: white;
    border-radius: 30px;
    font-size: x-large;
    
}
</style>