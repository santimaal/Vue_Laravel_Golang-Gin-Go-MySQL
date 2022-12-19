<template>
    <section class="allpage">
        <div class="mask d-flex h-100 align-items-center gradient-custom-3">
            <div class="container h-100">
                <div class="row d-flex justify-content-center align-items-center h-100">
                    <div class="col-12 col-md-9 col-lg-7 col-xl-6">
                        <div class="card" style="border-radius: 15px;">
                            <div class="card-body p-5">
                                <h2 class="text-uppercase text-center mb-5">Login</h2>
                                <form>
                                    <div class="form-outline mb-4">
                                        <label class="form-label" for="email">Your Email</label>
                                        <input type="email" v-model="state.form.email" id="email"
                                            class="form-control form-control-lg" />
                                        <span class="text-danger">{{ state.email }}</span>
                                    </div>

                                    <div class="form-outline mb-4">
                                        <label class="form-label" for="pass">Password</label>
                                        <input type="password" v-model="state.form.password" id="pass"
                                            class="form-control form-control-lg" />
                                        <!-- <span class="text-danger">{{ state.err.password.$model }}</span> -->
                                        <span class="text-danger" v-if="state.err.password.$invalid == true">Password is
                                            required</span>
                                    </div>

                                    <div class="d-flex justify-content-center">
                                        <button
                                            v-if="state.err.email.$invalid == true || state.err.password.$invalid == true"
                                            type="button"
                                            class="btn btn-light btn-block btn-lg gradient-custom-4 text-body border border-dark"
                                            disabled>Login</button>
                                        <button v-else type="button"
                                            class="btn btn-success btn-block btn-lg gradient-custom-4 text-body"
                                            @click="login">Login</button>
                                    </div>

                                    <p class="text-center text-muted mt-5 mb-0">You don't have an account?
                                        <router-link to="register" class="fw-bold text-body">
                                            <u>Register Here</u>
                                        </router-link>
                                    </p>

                                </form>

                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>

<script>
import { reactive, /*computed*/ } from 'vue';
import { useVuelidate } from '@vuelidate/core';
import { required, email } from '@vuelidate/validators';

export default {
    setup() {
        const state = reactive({
            form: {
                email: "",
                password: "",
            },
            err: {
                email: "",
                password: ""
            }
        })
        const rules = {
            email: { required, email },
            password: { required }
        }

        state.err = useVuelidate(rules, state.form);



        // if (state.err.email.$model == "") {
        //     console.log("Esta vacio el email");
        //     state.email = "";
        // } else {
        //     console.log("El email esta lleno");
        //     if (state.form.email == true) {
        //         state.email = "Email is invalid";
        //     } else {
        //         state.email = "";
        //     }
        // }
        // if (state.err.password.$model == "") {
        //     console.log("Esta vacio PASSWORD");
        // } else {
        //     console.log("PASSWORD llena");
        // }


        const login = () => {
            console.log("Hola soy el login");
            console.log(state.form.email);
            console.log(state.err.email);
            // if (state.form.email != "") {
            //     if (state.form.email == true) {
            //         state.email = "Password is required"
            //     } else {
            //         state.email = ""
            //     }

            // } else {
            //     state.pass = ""
            //     console.log(state.form);
            // }
        }

        return { state, login }
    }
}

</script>

<style scoped>
.allpage {
    height: 81vh;
}
</style>