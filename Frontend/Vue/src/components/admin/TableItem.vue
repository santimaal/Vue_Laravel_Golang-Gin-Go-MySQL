<template>
    <li class="li">{{ tableitem }}
        <div class="float-right">
            <button type="button" class="btn btn-primary m-1" @click="changeStatus(tableitem)">Status</button>
            <button class="btn btn-primary m-1" @click.stop="editTable(tableitem.id)">Edit</button>
            <button class="btn btn-primary m-1" @click.stop="deleteTable(tableitem.id)">Delete</button>
        </div>
    </li>
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
        const changeStatus = (tableitem) => {
            tableitem.is_active = tableitem.is_active == false ? true : false;
            store.dispatch("table/"+Constant.UPDATE_TABLE, { tableitem: tableitem });
        }
        const deleteTable = (id) => {
            store.dispatch("table/" + Constant.DELETE_TABLE, { id });
        }
        const editTable = (id) => {
            console.log(props.tableitem);
            store.dispatch("table/" + Constant.INITIALIZE_TABLE, { tableitem: { ...props.tableitem } });
            router.push({ name: 'updateTable', params: { id } })
        }

        return {deleteTable, editTable, checked, changeStatus }
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
</style>