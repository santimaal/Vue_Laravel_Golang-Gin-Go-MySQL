<template>
    <div class="notification dropleft">
            <svg viewBox="0 0 24 24" fill="currentColor" stroke="currentColor" stroke-width="2" stroke-linecap="round"
            stroke-linejoin="round" class="feather feather-bell dropdown-toggle" role="button" data-toggle="dropdown"
            aria-expanded="true" width="3vh">
            <path d="M18 8A6 6 0 006 8c0 7-3 9-3 9h18s-3-2-3-9M13.73 21a2 2 0 01-3.46 0" />
        </svg>
        <span class="translate-middle badge rounded-pill bg-danger" v-if="state.notis != 0">
            {{ state.notis }}
        </span>
        <!-- <a class="dropdown-toggle" href="#" role="button" data-toggle="dropdown" aria-expanded="false">A</a> -->
        <ul class="dropdown-menu mt-5 noti">
            <li><a class="dropdown-item" href="#">20-10-2001 12h Taula 2</a></li>
            <li><a class="dropdown-item" href="#">20-10-2001 12h Taula 2</a></li>
            <li><a class="dropdown-item" href="#">20-10-2001 12h Taula 2</a></li>
            <li><a class="dropdown-item" href="#">20-10-2001 12h Taula 2</a></li>
        </ul>
    </div>
</template>
<script>
import { reactive } from 'vue';


export default {
    setup() {
        const state = reactive({
            notis: 1
        })
        return { state }
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
    top: -1vh;
}

.noti {
    border: 1px solid black;
    background-color: rgba(0, 0, 0,0.95)
}
ul {
    border: 50%;
}
li a {
    color: white;
}
</style>