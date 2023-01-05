<template>
    <section class="bg-image allpage mt-4 mb-4">
        <div class="mask d-flex align-items-center h-100 gradient-custom-3">
            <div class="container h-100">
                <div class="row d-flex justify-content-center align-items-center h-100">
                    <div class="col-12 col-md-9 col-lg-7 col-xl-6">
                        <div class="card" style="border-radius: 15px;">
                            <div class="card-body p-5">
                                <h2 class="text-uppercase text-center mb-5">Create an account</h2>

                                <form>
                                    <div class="form-outline mb-4">
                                        <label class="form-label" for="name">Your Name</label>
                                        <input type="text" v-model="state.form.name" id="name"
                                            class="form-control form-control-lg" />
                                        <div class="input-errors" v-for="error of state.err.name.$errors"
                                            :key="error.$uid">
                                            <div class="error-msg">{{ error }}</div>
                                        </div>
                                    </div>

                                    <div class="form-outline mb-4">
                                        <label class="form-label" for="email">Your Email</label>
                                        <input type="email" v-model="state.form.email" id="email"
                                            class="form-control form-control-lg" />
                                        <div class="input-errors" v-for="error of state.err.email.$invalid"
                                            :key="error.$uid">
                                            <div class="error-msg">{{ error }}</div>
                                        </div>
                                        <span class="text-danger">{{ state.email }}</span>
                                    </div>

                                    <div class="form-outline mb-4">
                                        <label class="form-label" for="pass">Password</label>
                                        <input type="password" v-model="state.form.password" id="pass"
                                            class="form-control form-control-lg" />
                                        <span class="text-danger">{{ state.pass }}</span>
                                    </div>

                                    <div class="form-outline mb-4">
                                        <label class="form-label" for="pass">Repeat your password</label>
                                        <input type="password" v-model="state.form.rptpassword" id="pass2"
                                            class="form-control form-control-lg" />
                                    </div>

                                    <div class="form-check mb-3">
                                        <input class="form-check-input" type="checkbox" v-model="state.form.adminserver"
                                            value="" id="adminserver">
                                        <label class="form-check-label" for="adminserver">
                                            Admin Server
                                        </label>
                                    </div>


                                    <div class="d-flex justify-content-center">
                                        <button
                                            v-if="state.err.name.$invalid == true || state.err.email.$invalid == true || state.err.password.$invalid == true || state.err.rptpassword.$invalid == true"
                                            type="button"
                                            class="btn btn-success btn-block btn-lg gradient-custom-4 text-body"
                                            disabled>Register
                                        </button>
                                        <button v-else type="button"
                                            class="btn btn-success btn-block btn-lg gradient-custom-4 text-body"
                                            @click="register">Register
                                        </button>
                                    </div>

                                    <p class="text-center text-muted mt-5 mb-0">Have already an account?
                                        <router-link to="login" class="fw-bold text-body">
                                            <u>Login here</u>
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
import { reactive } from 'vue';
import { useVuelidate } from '@vuelidate/core'
import { required, email } from '@vuelidate/validators'
import { useStore } from 'vuex';
import Constant from '../../Constant';

export default {
    setup() {
        const store = useStore()
        const state = reactive({
            form: {
                name: "",
                email: "",
                password: "",
                rptpassword: "",
                adminserver: false,
            },
            err: {
                name: "",
                email: "",
                password: "",
                rptpassword: "",
            },
            pass: "",
            email: ""
        })


        const rules = {
            name: { required },
            email: { required, email },
            password: { required },
            rptpassword: { required }
        }



        state.err = useVuelidate(rules, state.form);
        const register = () => {
            if (state.form.password != state.form.rptpassword) {
                state.pass = "Is not the same password";
            } else {
                state.pass = "";
                if (state.form.adminserver) {
                    store.dispatch("user/" + Constant.USER_REGISTER_ADMIN, state.form);
                } else {
                    store.dispatch("user/" + Constant.USER_REGISTER, state.form);
                }
            }
        }
        return { state, register };
    }
}
</script>
<style>

</style>