<template>
  <div class="all_list">
    <div class="d-flex justify-content-center align-items-center">
      <div class="d-flex flex-column mt-5">
        <div class="div_border">
          <h1 class="text-center mt-2 title">Update Thematic</h1>
          <div class="pl-5 pr-5 pb-4 pt-3">
            <div class="form-group">
              <label htmlFor="todo">Name :</label>
              <input type="text" class="form-control" v-model="state.thematicitemlocal.name" />
            </div>
            <div class="form-group">
              <label htmlFor="desc">Location :</label>
              <input type="text" class="form-control" v-model="state.thematicitemlocal.location" />
            </div>
            <div class="form-group mb-4">
              <label htmlFor="desc">Img : </label>&nbsp;
              <input type="text" class="form-control" v-model="state.thematicitemlocal.img" />
            </div>
            <div class="d-flex justify-content-center">
              <button type="button" class="btn btn-outline-success mr-2" @click="updateThematic">Update</button>
              <router-link to="/athematic"><button type="button"
                  class="btn btn-outline-danger ml-2">Cancel</button></router-link>
            </div>
          </div>
        </div>
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
      state.thematicitemlocal = store.state.thematic.thematiclist.find(item => item.id == currentRoute.params.id);
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