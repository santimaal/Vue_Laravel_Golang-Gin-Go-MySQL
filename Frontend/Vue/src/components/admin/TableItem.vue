<template>
    <li>{{tableitem}}
        <div class="float-right">
            <span class="badge badge-secondary pointer ml-1" @click.stop="editTodo(tableitem.id)">Edit</span>
            <span class="badge badge-secondary pointer ml-1" @click.stop="deleteTodo(tableitem.id)">Delete</span>
        </div>
    </li>
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
        tableitem: Object
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
        const deleteTable = (id) => {
            store.dispatch(Constant.DELETE_TABLE, { id });
        }
        const editTable = (id) => {
            store.dispatch(Constant.INITIALIZE_TABLE, { tableitem: { ...props.tableitem } });
            router.push({ name: 'updateTable', params: { id } })
        }

        return { /*toggleDone,*/ deleteTable, editTable, checked }
    }
}
</script>

<style>

</style>