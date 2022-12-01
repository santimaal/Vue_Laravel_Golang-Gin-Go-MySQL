<template>
  <!-- <div class="container">
      <div class="row"> -->
  <div class="col p-3">
    <h2>Update thematic: {{ state.thematicitemlocal.id }}</h2>
  </div>
  <div class="row">
    <div class="col">
      <div class="form-group">
        <label htmlFor="todo">Name :</label>
        <input type="text" class="form-control" v-model="state.thematicitemlocal.name" />
      </div>
      <div class="form-group">
        <label htmlFor="desc">Location :</label>
        <input type="text" class="form-control" v-model="state.thematicitemlocal.location" />
      </div>
      <div class="form-group">
        <label htmlFor="desc">Img : </label>&nbsp;
        <input type="text" class="form-control" v-model="state.thematicitemlocal.img" />
      </div>
      <div class="form-group">
        <button type="button" class="btn btn-primary m-1" @click="updateThematic">Update</button>
        <router-link to="/athematic"><button type="button" class="btn btn-primary m-1">Cancel</button></router-link>
      </div>
    </div>
  </div>
</template>
  
<script>
import { reactive } from 'vue'
import { useStore } from 'vuex'
import { useRouter, useRoute } from 'vue-router';
import Constant from '../../../Constant';
import ThematicService from "@/services/ThematicService";

export default {
  setup() {
    const store = useStore();
    const router = useRouter();
    const currentRoute = useRoute();
    const thematicitem = store.state.thematic.thematiclist.find(item => item.id == currentRoute.params.id)
    const state = reactive({
      thematicitemlocal: { ...thematicitem },
    });

    if (store.state.thematic.thematiclist.length != 0) {
      state.thematicitemlocal =  store.state.thematic.thematiclist.find(item => item.id == currentRoute.params.id);
    } else {
      ThematicService.getThematicById(currentRoute.params.id).then(data => {
        state.thematicitemlocal = data.data;
      })

    }

    const updateThematic = () => {
      router.push({ name: "admin_thematic" });
      store.dispatch("thematic/" + Constant.UPDATE_THEMATIC, { thematicitem: state.thematicitemlocal });
    }

    return { state, updateThematic }

  }
}
</script>
  
<style>

</style>