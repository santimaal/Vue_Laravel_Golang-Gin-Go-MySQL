<template>
    <!-- <li>{{ thematicitem }}</li> -->

    <div class="card">
        <img :src="thematicitem.img" alt="" style="width:20%">
        {{ thematicitem.location }}
    </div>
    <!-- <li :class="checked(tableitem.done)" @click="toggleDone(tableitem.id)">
            <span :class="{ pointer:true, 'todo-done':tableitem.done }" :title="'Titulo : ' + tableitem.capacity">
                {{tableitem.id}}
                {{tableitem.done ? " (done)" : ""}}
            </span>
    </li> -->
</template>

<script>
import Constant from '../../Constant';
import { useStore } from 'vuex'
import { useRouter } from 'vue-router';

export default {
    props: {
        thematicitem: Object
    },
    setup(props) {
        const store = useStore();
        const router = useRouter();

        const checked = (done) => {
            return { "list-group-item": true, "list-group-item-success": done };
        }
        // const toggleDone = (id) => {
        //     store.dispatch(Constant.TOGGLE_DONE, { id });
        // }
        const deleteThematic = (id) => {
            store.dispatch(Constant.DELETE_THEMATIC, { id });
        }
        const editThematic = (id) => {
            store.dispatch(Constant.INITIALIZE_THEMATIC, { thematicitem: { ...props.thematicitem } });
            router.push({ name: 'updateThematic', params: { id } })
        }

        return { /*toggleDone,*/ deleteThematic, editThematic, checked }
    }
}
</script>

<style>

</style>