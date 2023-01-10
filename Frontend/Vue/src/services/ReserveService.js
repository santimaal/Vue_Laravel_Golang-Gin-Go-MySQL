import Api from '@/services/Api'
import secret from '../secret.js'

export default {
    getHours(date) {
        console.log(date.dateini.substring(0,10));
        return Api(secret.GO_APP_URL).get(`reserve/hours/${date.dateini.substring(0,10)}/${date.id}`)
    },
    addReserve(info) {
        console.log(info);
        return Api(secret.GO_APP_URL).post(`reserve/add`, info)
    },
}
