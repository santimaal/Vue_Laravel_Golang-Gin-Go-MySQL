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

                                    <div class="d-flex justify-content-center">
                                        <button
                                            v-if="state.err.name.$invalid == true || state.err.email.$invalid == true || state.err.password.$invalid == true || state.err.rptpassword.$invalid == true"
                                            type="button"
                                            class="btn btn-success btn-block btn-lg gradient-custom-4 text-body"
                                            disabled>Register</button>
                                        <button v-else type="button"
                                            class="btn btn-success btn-block btn-lg gradient-custom-4 text-body"
                                            @click="register">Register</button>
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

export default {
    setup() {
        const state = reactive({
            form: {
                name: "",
                email: "",
                password: "",
                rptpassword: "",
            },
            err: {
                name: "",
                email: "",
                password: "",
                rptpassword: "",
            },
            pass: ""
        })

        const rules = {
            name: { required },
            email: { required, email },
            password: { required },
            rptpassword: { required }
        }



        state.err = useVuelidate(rules, state.form)
        // console.log(v$);
        const register = () => {
            console.log("mismo");
            if (state.form.password != state.form.rptpassword) {
                console.log("hola");
                state.pass = "Is not the same password"
            } else {
                state.pass = ""
                console.log(state.form);
            }
        }
        return { state, register }
    }
}
</script>
<style>

</style>