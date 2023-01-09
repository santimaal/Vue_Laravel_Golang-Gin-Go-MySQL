import Api from '@/services/Api'
import secret from '../secret.js'

export default {
    getHours(date) {
        console.log(date);
        return Api(secret.GO_APP_URL).get(`reserve/hours`, date)
    },
}
