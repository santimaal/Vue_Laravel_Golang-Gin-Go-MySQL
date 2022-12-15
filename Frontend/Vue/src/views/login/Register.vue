<template>
    <section class="bg-image allpage">
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
                                        <span style="color:red">{{ state.err.name }}</span>
                                    </div>

                                    <div class="form-outline mb-4">
                                        <label class="form-label" for="email">Your Email</label>
                                        <input type="email" v-model="state.form.email" id="email"
                                            class="form-control form-control-lg" />
                                        <span style="color:red">{{ state.err.email }}</span>

                                    </div>

                                    <div class="form-outline mb-4">
                                        <label class="form-label" for="pass">Password</label>
                                        <input type="password" v-model="state.form.password" id="pass"
                                            class="form-control form-control-lg" />
                                        <span style="color:red">{{ state.err.password }}</span>
                                    </div>

                                    <div class="form-outline mb-4">
                                        <label class="form-label" for="pass">Repeat your password</label>
                                        <input type="password" v-model="state.form.rptpassword" id="pass2"
                                            class="form-control form-control-lg" />
                                    </div>

                                    <div class="d-flex justify-content-center">
                                        <button type="button"
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
                password: ""
            }
        })

        const register = () => {
            if (check()==false) {
                console.log("mal");
            } else {
                console.log("bien");
            }
            console.log(state.form);
        }

        function check() {
            let errors = 0
            if (state.form.name.length == 0) {
                state.err.name = "Can't be blank"
                errors++
            } else {
                state.err.name = ""
            }
            if (state.form.email.length == 0) {
                state.err.email = "Can't be blank"
                errors++
            } else {
                state.err.email = ""
            }
            if (state.form.password != state.form.rptpassword) {
                state.err.password = "Is not the same password"
                errors++
            } else {
                state.err.password = ""
            }
            if (state.form.password.length == 0) {
                state.err.password = "Can't be blank"
                errors++
            }
            if (errors != 0) {
                return false
            } else {
                return true
            }
        }
        return { state, register }
    }
}
</script>
<style>

</style>