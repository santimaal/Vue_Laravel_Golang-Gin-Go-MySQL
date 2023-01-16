<template>
    <tr>
        <th scope="row">{{ tableitem.id }}</th>
        <td v-if="tableitem.is_active == false">No Reserved</td>
        <td v-if="tableitem.is_active == true">Reserved</td>
        <td>{{ tableitem.capacity }}</td>
        <td>{{ tableitem.location }}</td>
        <td>
            <div v-for="(item, id) in state.thematic" :key="id">
                <span v-if="(item.id == tableitem.id_thematic)">{{ item.name }}</span>
            </div>
        </td>
        <td colspan="2">
            <button class="btn btn-primary m-1" @click.stop="editTable(tableitem.id)">Edit</button>
            <button class="btn btn-primary m-1" @click.stop="deleteTable(tableitem.id)">Delete</button>
        </td>
    </tr>
</template>

<script>
import Constant from '../../Constant';
import { useStore } from 'vuex'
import { useRouter } from 'vue-router';
import { reactive } from 'vue';

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
        const deleteTable = (id) => {
            store.dispatch("table/" + Constant.DELETE_TABLE, { id });
        }
        const editTable = (id) => {
            console.log(props.tableitem);
            store.dispatch("table/" + Constant.INITIALIZE_TABLE, { tableitem: { ...props.tableitem } });
            router.push({ name: 'updateTable', params: { id } })
        }
        setTimeout(() => {
            state.thematic = store.getters["thematic/getThematic"]
        }, 50);

        const state = reactive({
            thematic: store.state.thematic.thematiclist
            // thematic: ThematicService.getAllThematic().data
        })
        return { deleteTable, editTable, checked, state }
    }
}
</script>

<style>
.li {
    margin-top: 2%;
    background-color: ghostwhite;
    padding: 2%;
    list-style: none;
}

td {
    text-align: center;
}
</style>