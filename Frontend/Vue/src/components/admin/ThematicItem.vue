<template>
  <tr>
    <th scope="row">{{ thematicitem.id }}</th>
    <td>{{ thematicitem.name }}</td>
    <td>{{ thematicitem.location }}</td>
    <td>{{ thematicitem.img }}</td>
    <td colspan="2"> 
      <button class="btn btn-primary m-1" @click.stop="editThematic(thematicitem.id)">Edit</button>
      <button class="btn btn-primary m-1" @click.stop="deleteThematic(thematicitem.id)">Delete</button>
    </td>
  </tr>
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
    const deleteThematic = (id) => {
      store.dispatch("thematic/" + Constant.DELETE_THEMATIC, { id });
    }
    const editThematic = (id) => {
      store.dispatch("thematic/" + Constant.INITIALIZE_THEMATIC, { thematicitem: { ...props.thematicitem } });
      router.push({ name: 'updateThematic', params: { id } })
    }

    return { deleteThematic, editThematic, checked }
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
td{
    text-align: center;
}
</style>