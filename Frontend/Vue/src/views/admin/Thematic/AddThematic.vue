<template>
    <form>
        <h1>Create Thematic</h1>
        Name <input type="text" v-model="state.thematic.name" required />
        <br><br>
        Location <input type="text" v-model="state.thematic.location" required>
        <br><br>
        Img <input type="text" v-model="state.thematic.img" required>
        <br><br>
        <!-- <button type="submit" @click="addThematic">Add</button> -->
        <button type="button" @click="addThematic">Add</button>
        <router-link to="athematic"><button>Cancel</button></router-link>
    </form>

</template>

<script>
import Constant from '../../../Constant';
import { reactive } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';

export default {
    setup() {
        const store = useStore();
        const router = useRouter();
        const state = reactive({
            thematic: { name: "", location: "", img: "" }
        });

        const addThematic = () => {
            if (state.thematic.name == "" || state.thematic.location == "" || state.thematic.img == "") {
                console.log("todo mal");
            } else {
                console.log("buen")
                store.dispatch("thematic/" + Constant.ADD_THEMATIC, { thematic: state.thematic })
                router.push({ name: "admin_thematic" });
            }
        }


        return { state, addThematic }
    }
}

</script>

<style>

</style>