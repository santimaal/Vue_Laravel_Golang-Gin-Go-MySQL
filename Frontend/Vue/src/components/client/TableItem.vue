<template>
    <li>{{ tableitem }}</li>
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
        const deleteTable = (id) => {
            store.dispatch(Constant.DELETE_TABLE, { id });
        }
        const editTable = (id) => {
            store.dispatch(Constant.INITIALIZE_TABLE, { tableitem: { ...props.tableitem } });
            router.push({ name: 'updateTable', params: { id } })
        }

        return {deleteTable, editTable, checked }
    }
}
</script>

<style>

</style>