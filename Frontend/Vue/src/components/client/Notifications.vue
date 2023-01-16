<template>
    <div class="notification dropleft" @click="getNotis">
        <div class="test" data-toggle="dropdown" aria-expanded="true">
            <svg viewBox="0 0 24 24" fill="currentColor" stroke="currentColor" stroke-width="2" stroke-linecap="round"
                stroke-linejoin="round" class="feather feather-bell dropdown-toggle" role="button" width="3vh">
                <path d="M18 8A6 6 0 006 8c0 7-3 9-3 9h18s-3-2-3-9M13.73 21a2 2 0 01-3.46 0" />
            </svg>
            <span class="translate-middle badge rounded-pill bg-danger" v-if="state.notis != 0">
                {{ state.notis }}
            </span>
        </div>
        
        <ul class="dropdown-menu mt-5 noti">
            <li v-if="state.notification.length == 0">
                <div class="ml-1 text-white">No notifications yet</div>
            </li>
            <li v-for="noti in state.notification" :key="noti.id">
                <div class="dropdown-item text-white">{{ noti.dateini }} is {{ noti.is_confirmed }}</div>
            </li>
        </ul>
    </div>
</template>
<script>
import { reactive, computed } from 'vue';
import { useStore } from 'vuex';
import { useGetNotisClient } from '../../composables/notifications/useNotifications'

export default {
    setup() {
        const store = useStore()
        const state = reactive({
            notis: computed(() => store.getters["user/getNotis"]),
            notification: [],
        })

        const getNotis = async () => {
            store.getters["user/setNotis"];
            state.notification = await useGetNotisClient()
        }
        return { state, getNotis }
    }
}
</script>
<style scoped>
path {
    color: lightgray;
}

.notification {
    display: flex;
    flex-direction: column;
}

span {
    position: absolute;
    top: -1.5vh;
    left: 50%;
}

.noti {
    border: 1px solid black;
    background-color: rgba(0, 0, 0, 0.95)
}

ul {
    border: 50%;
}

li a {
    color: white;
}

.border-radius {
    border-radius: 100%;
    border-color: transparent !important;
    ;
}
</style>