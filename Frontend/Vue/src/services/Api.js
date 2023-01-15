import axios from 'axios'
import router from '../router/index';
import { useStore } from "vuex";
import Constant from "./../Constant";
import secrets from "../secret";
import { createToaster } from "@meforma/vue-toaster";

const toaster = createToaster({position: "top-right", duration: 5000, dismissible: true });


export default (URL) => {
    const store = useStore();

    const axiosInstance =
        axios.create({
            baseURL: URL
        })
    //Depediendo de la URL pasaremos un token u otro
    const token = URL === secrets.LARAVEL_APP_URL ? localStorage.getItem('token_admin') : localStorage.getItem('token');
    if (token) {
        axiosInstance.defaults.headers.common.Authorization = `Bearer ${token}`;
    }

    axiosInstance.interceptors.response.use(
        (response) => response,
        (error) => {
            console.log(error)
            if (error.response.status === 401) {
                store.dispatch('user/'+Constant.LOGOUT);
                toaster.error("You don't have permission LOGOUT!");
                router.push({ name: "login" });
            }
           
            return Promise.reject(error)
        }
    )
    return axiosInstance
}