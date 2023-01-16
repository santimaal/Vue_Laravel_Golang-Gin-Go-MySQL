import Api from '@/services/Api'
import secret from '../secret.js'

export default {
   
    userRegisterAdmin(data) {
        return Api(secret.LARAVEL_APP_URL).post('user/register', data);
    },
    userLoginAdmin(data) {
        return Api(secret.LARAVEL_APP_URL).post('user/login', data);
    },
    getProfile_Admin() {
        return Api(secret.LARAVEL_APP_URL).get('user/profile');
    },
    getNotificationsAdmin() {
        return Api(secret.LARAVEL_APP_URL).get('reserve');
    },
    getUsersAdmin() {
        return Api(secret.LARAVEL_APP_URL).get('user/GetUsersClient');
    },
    updateStateUser(status, id) {               
        return Api(secret.LARAVEL_APP_URL).put(`user/UpdateStateUser/${id}/${status}`);
    },
    isAdmin() {
        return Api(secret.LARAVEL_APP_URL).get('user/isAdmin');
    },
    userRegister(data) {
        return Api(secret.GO_APP_URL).post('user/register', data);
    },
    userLogin(data) {
        return Api(secret.GO_APP_URL).post('user/login', data);
    },
    getProfile() {
        return Api(secret.GO_APP_URL).get('user/profile');
    },
    getNotificationsClient() {
        return Api(secret.GO_APP_URL).get('reserve');
    },
    getMyReserves() {
        return Api(secret.GO_APP_URL).get('reserve/myreserve');
    },
    userUpdate(data) {
        return Api(secret.GO_APP_URL).put('user/update', data);
    },
    sendNotification(data) {
        return Api(secret.GO_APP_URL).post('sendTel', data);
    }
}