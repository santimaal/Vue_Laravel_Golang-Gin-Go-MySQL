<template>
    <div class="container">

        <div class="row justify-content-center">

            <div class="col-md-6">

                <div class="form">
                    <i class="fa fa-search">lupa</i>
                    <input type="text" v-model="state.filter.search" class="form-control form-input"
                        placeholder="Search anything..." v-on:change="search" v-on:keyup.enter="search">
                    <span class="left-pan"><i class="fa fa-microphone">mic</i></span>
                </div>

            </div>

        </div>

    </div>
</template>

<script>
import { reactive, getCurrentInstance } from 'vue';
import { useRoute } from 'vue-router';

export default {
    setup() {
        const { emit } = getCurrentInstance();
        const currentRoute = useRoute()

        const state = reactive({
            filter: {
                search: ""
            }
        })

        if (currentRoute.params.filter && currentRoute.params.filter != "all") {
            state.filter.search = JSON.parse(atob(currentRoute.params.filter)).search
        }

        const search = () => {
            emit('search', state.filter)
        }
        return { search, state }
    }

}
</script>

<style>
body {

    background: #d1d5db;
}

.height {

    height: 100vh;
}

.form {

    position: relative;
}

.form .fa-search {

    position: absolute;
    top: 20px;
    left: 20px;
    color: #9ca3af;

}

.form span {

    position: absolute;
    right: 17px;
    top: 13px;
    padding: 2px;
    border-left: 1px solid #d1d5db;

}

.left-pan {
    padding-left: 7px;
}

.left-pan i {

    padding-left: 10px;
}

.form-input {

    height: 55px;
    text-indent: 33px;
    border-radius: 10px;
}

.form-input:focus {

    box-shadow: none;
    border: none;
}
</style>
