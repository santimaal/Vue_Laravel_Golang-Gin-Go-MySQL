<template>
    <section class>
        <div class="container py-5">
            <div class="row">
                <div class="col-lg-4">
                    <div class="card mb-4">
                        <div class="card-body text-center">
                            <img :src=state.user.img alt="avatar" class="img_profile rounded-circle img-fluid">
                            <hr>
                            <form class="mt-4">
                                <!-- Username -->
                                <div class="form-outline mb-4">
                                    <div class="row ml-2 div_row">
                                        <p @click="state.check.name = ''" class="p_icon"><font-awesome-icon
                                                class="U_icon " icon="fa-solid fa-user" /></p>
                                        <h5 v-if="state.check.name != true"
                                            class="d-flex align-items-center col-8 ml-2 mr-2">{{ state.user.name }}</h5>
                                        <input type="text" v-if="state.check.name == true" v-model="state.form.name"
                                            :placeholder=state.user.name id="name"
                                            class="ml-2 mr-2 text-center col-8 form-control form-control-md" />
                                        <p v-if="state.check.name != true" @click="state.check.name = true"
                                            class="btn_edit"><font-awesome-icon class="edit_icon"
                                                icon="fa-solid fa-pen-to-square" /></p>
                                        <p v-if="state.check.name == true" @click=UpdateProfile() class="p_icon">
                                            <font-awesome-icon class="U_icon check_icon" icon="fa-solid fa-check" />
                                        </p>
                                    </div>
                                </div>
                                <hr>
                                <!-- Email -->
                                <div class="form-outline mb-4">
                                    <div class="row ml-2">
                                        <p @click="state.check.email = ''" class="p_icon"><font-awesome-icon
                                                class="U_icon" icon="fa-solid fa-envelope" /></p>
                                        <h5 v-if="state.check.email != true"
                                            class="d-flex align-items-center ml-2 mr-2 col-8">{{
                                                state.user.email.substring(0, 5)
                                            }}@gmail.com</h5>
                                        <input type="email" v-if="state.check.email == true" v-model="state.form.email"
                                            :placeholder=state.user.email id="email"
                                            class="ml-2 mr-2 text-center col-8 form-control form-control-md" />
                                        <p v-if="state.check.email != true" @click="state.check.email = true"
                                            class="btn_edit"><font-awesome-icon class="edit_icon"
                                                icon="fa-solid fa-pen-to-square" /></p>
                                        <p v-if="state.check.email == true" @click=UpdateProfile() class="p_icon">
                                            <font-awesome-icon class="U_icon check_icon" icon="fa-solid fa-check" />
                                        </p>
                                    </div>
                                    <div class="input-errors" v-for="error of state.err.email.$invalid"
                                        :key="error.$uid">
                                        <div class="error-msg">{{ error }}</div>
                                    </div>
                                </div>
                                <hr>
                                <!-- Password -->
                                <div class="form-outline mb-4">
                                    <div class="row ml-2">
                                        <p @click="state.check.password = ''" class="p_icon"><font-awesome-icon
                                                class="U_icon" icon="fa-solid fa-key" />
                                        </p>
                                        <h5 v-if="state.check.password != true"
                                            class="d-flex align-items-center ml-2 mr-2 col-8">**********</h5>
                                        <input type="password" v-if="state.check.password == true"
                                            v-model="state.form.password" placeholder="**********" id="password"
                                            class="ml-2 mr-2 text-center col-8 form-control form-control-md" />
                                        <p v-if="state.check.password != true" @click="state.check.password = true"
                                            class="btn_edit"><font-awesome-icon class="edit_icon"
                                                icon="fa-solid fa-pen-to-square" /></p>
                                        <p v-if="state.check.password == true" @click=UpdateProfile() class="p_icon">
                                            <font-awesome-icon class="U_icon check_icon" icon="fa-solid fa-check" />
                                        </p>
                                    </div>
                                </div>
                                <hr>
                                <!-- Img -->
                                <div class="form-outline mb-4">
                                    <div class="row ml-2">
                                        <p @click="state.check.img = ''" class="p_icon"><img :src=state.user.img
                                                alt="avatar" class="rounded img-fluid" style="width: 35px;"></p>
                                        <h5 v-if="state.check.img != true"
                                            class="d-flex align-items-center ml-2 mr-2 col-8">http://{{
                                                state.user.name.substring(0, 6)
                                            }}.png</h5>
                                        <input type="text" v-if="state.check.img == true" v-model="state.form.img"
                                            placeholder="http//url_imagen.png" id="img"
                                            class="ml-2 mr-2 text-center col-8 form-control form-control-md" />
                                        <p v-if="state.check.img != true" @click="state.check.img = true"
                                            class="btn_edit"><font-awesome-icon class="edit_icon"
                                                icon="fa-solid fa-pen-to-square" /></p>
                                        <p v-if="state.check.img == true" @click=UpdateProfile() class="p_icon">
                                            <font-awesome-icon class="U_icon check_icon" icon="fa-solid fa-check" />
                                        </p>
                                    </div>
                                </div>
                                <hr>
                                <!-- Telegram -->
                                <div class="form-outline mb-4">
                                    <div class="row ml-1">
                                        <button type="button" class="btn_tele">
                                            <img @click="telegram" class="img_tele"
                                                src="https://i.postimg.cc/L5J3J8kf/telegram.png" alt="og_tel"
                                                v-b-tooltip.hover placement="right" variant="info"
                                                title="Do you want to activate notifications by telegram? ¡Click on the icon!">
                                        </button>
                                        <input v-if="state.check.chat_id == true" type="text" readonly
                                            v-model="state.form.chat_id" id="chat_id"
                                            class="ml-2 mr-1 text-center col-8 form-control form-control-md" />
                                        <button v-if="state.check.chat_id == true" type="button" class="btn_nube">
                                            <img @click="saveChatId" class="img_nube"
                                                src="https://i.postimg.cc/13pY6Fzs/nube.png" alt="og_tel"
                                                v-b-tooltip.hover placement="right" variant="info"
                                                title="Click me for Coppy, Save and Redirect">
                                        </button>
                                    </div>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
                <div class="col-lg-8">
                    <div class="card mb-4">
                        <h1 class="title">My Reserves</h1>
                        <ListReserves></ListReserves>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>

<script>
import { reactive, computed } from 'vue';
import { useVuelidate } from '@vuelidate/core'
import { email } from '@vuelidate/validators'
import { useStore } from 'vuex';
import Constant from '../../Constant';
import { createToaster } from "@meforma/vue-toaster";
import ListReserves from "../../components/client/List_MyReserves.vue";

const toaster = createToaster();

export default {
    setup() {
        const store = useStore()
        const state = reactive({
            form: {
                name: "",
                email: "",
                password: "",
                img: "",
                chat_id: "",
            },
            check: {
                name: "",
                email: "",
                password: "",
                img: "",
                chat_id: "",
            },
            err: {
                email: "",
            },
            user: computed(() => store.getters["user/getUser"]),
        })
        const rules = {
            email: { email }
        }

        state.err = useVuelidate(rules, state.form);

        if (state.err.email.$model != "") {
            toaster.error("Email is ivalid", { position: "bottom", duration: 5000, dismissible: true });
        } else {
            state.err.email = "";
        }


        const UpdateProfile = () => {
            state.check.name = false;
            state.check.email = false;
            state.check.password = false;
            state.check.img = false;
            store.dispatch("user/" + Constant.USER_UPDATE, state.form);


        }
        const telegram = () => {
            const chatId = generateChatId();
            state.form.chat_id = chatId;
        }
        const generateChatId = () => {
            state.check.chat_id = true;
            var characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
            var result = "";
            for (var i = 0; i < 10; i++) {
                result += characters.charAt(Math.floor(Math.random() * characters.length));
            }
            return result;
        }

        const saveChatId = async () => {
            state.check.chat_id = "";
            await navigator.clipboard.writeText(state.form.chat_id);
            store.dispatch("user/" + Constant.USER_UPDATE, state.form);
            window.open("https://web.telegram.org/z/#5944160111", "_blank");

        }
        return { state, UpdateProfile, telegram, generateChatId, saveChatId, ListReserves }
    },
    components: { ListReserves },
}
</script>
<style>
@import url("https://fonts.googleapis.com/css2?family=Lobster&display=swap");

.img_profile {
    width: 150px;
}

.img_tele {
    width: 35px;
}

.img_nube {
    width: 32px;
    border-radius: 4px;
}

.btn_nube {
    border: 0px;
    width: auto;
    max-width: 35px;
    background-color: black;
    border-radius: 4px;
    display: flex;
    justify-content: center;
    align-items: center;
}

.btn_tele {
    border: 0px;
    width: auto;
    max-width: 35px;
    background-color: transparent;
}

.U_icon {
    width: 35px;
    color: white;
}

.p_icon {
    background-color: black;
    width: 35px;
    height: 35px;
    display: flex;
    align-items: center;
    border-radius: 4px;
}

.btn_edit {
    background-color: transparent;
    width: 35px;
    height: 35px;
    display: flex;
    align-items: center;
    border: 0px;
    font-size: x-large;
}

.edit_icon {
    color: black;
}

.check_icon:hover {
    font-size: x-large;
    color: green;
}

.title{
    font-family: "Lobster", cursive;
    text-align: center;
    text-decoration: underline;
}
</style>