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
                                        <span style="color:red">{{ state.err.email }}</span>
                                    </div>

                                    <div class="form-outline mb-4">
                                        <label class="form-label" for="pass">Password</label>
                                        <input type="password" v-model="state.form.password" id="pass"
                                            class="form-control form-control-lg" />
                                        <span style="color:red">{{ state.err.password }}</span>
                                    </div>

                                    <div class="d-flex justify-content-center">
                                        <button type="button"
                                            class="btn btn-success btn-block btn-lg gradient-custom-4 text-body"
                                            @click="login">Register</button>
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
import { reactive } from 'vue';

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

        const login = () => {
            if (check() == false) {
                console.log("mal");
            } else {
                console.log("bien");
            }
            console.log(state.form);
        }

        function check() {
            let errors = 0
            if (state.form.email.length == 0) {
                state.err.email = "Can't be blank"
                errors++
            } else {
                state.err.email = ""
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

        return { state, login }
    }
}
</script>
<style>
.allpage {
    height: 81vh;
}
</style>