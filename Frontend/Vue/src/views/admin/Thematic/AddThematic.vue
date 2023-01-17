<template>
    <div class="all_list">
        <div class="d-flex justify-content-center align-items-center">
            <div class="d-flex flex-column mt-5">
                <div class="div_border">
                    <h1 class="text-center mt-2 title">Create Thematic</h1>
                    <div class="pl-5 pr-5 pb-4 pt-2">
                        <form>
                            <div class="mb-4">
                                <label> Name:</label><br>
                                <input type="text" v-model="state.thematic.name" required />
                            </div>
                            <div class="mb-4">
                                <label>Location:</label><br>
                                <input type="text" v-model="state.thematic.location" required>
                            </div>
                            <div class="mb-4">
                                <label>Img:</label><br>
                                <input type="text" v-model="state.thematic.img" required>
                            </div>

                            <div class="d-flex justify-content-center">
                                <button type="button" class="btn btn-outline-success mr-2"
                                    @click="addThematic">Add</button>
                                <router-link to="athematic"><button
                                        class="btn btn-outline-danger ml-2">Cancel</button></router-link>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>

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
                console.log("err");
            } else {
                store.dispatch("thematic/" + Constant.ADD_THEMATIC, { thematic: state.thematic })
                router.push({ name: "admin_thematic" });
            }
        }


        return { state, addThematic }
    }
}

</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Lobster&display=swap");

.all_list {
    height: 81vh !important;
}

.div_border {
    border: 2px solid black;
    background-color: white;
    border-radius: 5px;
}

.title {
    text-decoration: underline;
    font-family: "Lobster", cursive;
}
</style>