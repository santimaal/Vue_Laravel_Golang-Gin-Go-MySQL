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
                                        <span class="text-danger">{{ state.passwd }}</span>
                                    </div>

                                    <div class="d-flex justify-content-center">
                                        <button
                                            v-if="state.err.email.$invalid == true || state.err.password.$invalid == true"
                                            type="button"
                                            class="btn btn-light btn-block btn-lg gradient-custom-4 text-body border border-dark"
                                            @click="login">Login</button>
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
import Constant from '../../Constant';
import { useStore } from 'vuex';

export default {
    setup() {
        const store = useStore();
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

        const login = () => {
            let count = 0

            if (state.err.email.$invalid == true) {
                if (state.err.email.$model != "") {
                    state.email = "Email is invalid";
                } else {
                    state.email = "Email is required";
                }

            } else {
                state.email = "";
                count++;
            }
            if (state.err.password.$invalid == true) {
                state.passwd = "Password is required"
            } else {
                state.passwd = "";
                count++;
            }
            if (count == 2) {
                store.dispatch("user/" + Constant.USER_LOGIN, state.form);
            }
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