<template>
    <nav class="nav_pag" aria-label="">
        <ul class="pagination justify-content-center">
            <li class="page-item">
                <a class="page-link" @click="changePage(state.page - 1)" aria-label="Previous">
                    <span aria-hidden="true">&laquo;</span>
                    <span class="sr-only">Previous</span>
                </a>
            </li>
            <li class="page-item" v-for="(row, id) in state.total_pages" :key="id" :class="isActive(id)"
                @click="changePage(id)">
                <a class="page-link num_pag" href="#">{{ row }}</a>
            </li>
            <li class="page-item">
                <a class="page-link" @click="changePage(state.page + 1)" aria-label="Next">
                    <span aria-hidden="true">&raquo;</span>
                    <span class="sr-only">Next</span>
                </a>
            </li>
        </ul>
    </nav>
</template>
<script>
import { reactive, getCurrentInstance } from 'vue';

export default {
    props: {
        totalpages: Number
    },
    setup(props) {
        const { emit } = getCurrentInstance();
        const state = reactive({
            total_pages: props.totalpages,
            page: 0
        })

        const isActive = (id) => {
            return { active: id == state.page };
        };

        const changePage = (page) => {
            if (page < 0) {
                state.page = 0;
            } else if (page == state.total_pages) {
                state.page = state.total_pages - 1;
            } else {
                state.page = page;
            }
            emit('changepage', state.page)
        }
        return { state, changePage, isActive }
    }
}
</script>
<style scoped>
nav {
    background-color: rgb(201, 243, 234) !important;
}
.page-link{
    background-color: transparent;
    border: 0px;
    color:black;
}
.page-link span{
font-size: large;
}
.num_pag{
    border-radius: 20px;
    border: 2px solid black;
}
.page-item{
padding: 3px;
}

.page-item.active .page-link{
    border-radius: 20px;
    border: 2px solid rgb(92, 91, 91);
    background-color: black;
    color: white;
    font-weight: bold;
}

</style>