import Api from '@/services/Api'
import secret from '../secret.js'

export default {
    getHours(date) {
        return Api(secret.GO_APP_URL).get(`reserve/hours/${date.dateini.substring(0, 10)}/${date.id}`)
    },
    addReserve(info) {
        return Api(secret.GO_APP_URL).post(`reserve/add`, info)
    },
    updateReserve(status, id) {
        return Api(secret.LARAVEL_APP_URL).put(`reserve/${id}`, { is_confirmed: status })
    },
}
