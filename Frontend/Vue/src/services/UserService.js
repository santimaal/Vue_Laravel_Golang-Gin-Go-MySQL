import Api from '@/services/Api'
import secret from '../secret.js'

export default {
    userRegister(data) {
        console.log(data);
        return Api(secret.GO_APP_URL).post('user/register', data)
    },
}