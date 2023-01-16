<template>
  <div class="table-wrapper-scroll-y my-custom-scrollbar">

    <table class="text-center table table-bordered table-striped mb-0 table-light">
      <thead class="bord_thead">
        <tr>
          <th scope="col">IMG</th>
          <th scope="col">Name</th>
          <th scope="col">Day</th>
          <th scope="col">Hour</th>
          <th scope="col">Satus</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="reserve in state.reservas" :key="reserve.id" >
          <th scope="row"><img class="img_reserve" src="https://architecturesideas.com/wp-content/uploads/2020/10/Guide-to-Restaurant-Tables-5.jpg" alt=""></th>
          <td>Table {{ reserve.id_table }}</td> 
          <td>{{ reserve.dateini.split(' ')[0] }}</td>
          <td>{{ reserve.dateini.split(' ')[2].split('h')[0]}}:00</td>
          <td  v-if="reserve.is_confirmed == 'accepted'"><font-awesome-icon class="bg-green" icon="fa-solid fa-check"/></td>
          <td  v-if="reserve.is_confirmed == 'pending'"><font-awesome-icon class="bg-blue" icon="fa-solid fa-arrows-rotate"/></td>
          <td  v-if="reserve.is_confirmed == 'denied'"><font-awesome-icon class="bg-red" icon="fa-solid fa-x"/></td>
        </tr>
      </tbody>
    </table>

  </div>

</template>

<script>
import { reactive } from 'vue';
import { useGetReserveClient } from '../../composables/reserve/useReserve'


export default {
  setup() {
    const state = reactive({
      myreserves: useGetReserveClient(),
      reservas: []
    })
  
    state.myreserves.then(resolvedValue => {
     state.reservas = resolvedValue._rawValue;
     console.log(state.reservas);
    });
    
    return { state }
  }
}
</script>

<style>
.my-custom-scrollbar {
  position: relative;
  height: 600px;
  overflow: auto;
}

.table-wrapper-scroll-y {
  display: block;
}
.bg-green {
 color: lightgreen !important;
 background-color: black;
 font-size: x-large;
 border-radius: 50%;
}
.bg-red {
  color: lightcoral !important;
  background-color: black;
  font-size: x-large;
}
.bg-blue {
  color: lightblue !important;
  background-color: black;
  font-size: x-large;
  border-radius: 20%;
}
.img_reserve{
  width: 50px;
  background-color:transparent;
  border: 0px solid black;
  border-radius: 10px;
}
</style>
