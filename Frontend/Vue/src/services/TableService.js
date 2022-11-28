import Api from '@/services/Api'
import secret from '../secret.js'

export default {
  getAllTable() {
    return Api(secret.LARAVEL_APP_URL).get(`table/`)
  },
  getTableById(id) {
    return Api(secret.LARAVEL_APP_URL).get(`table/${id}`)
  },
  createTable(data) {
    return Api(secret.LARAVEL_APP_URL).post('table/', data)
  },
  updateTable(data, id) {
    return Api(secret.LARAVEL_APP_URL).put(`table/${id}`, data)
  },
  changeStatusTable(data, id) {
    return Api(secret.LARAVEL_APP_URL).put(`table/changeStatus/${id}`, data)
  },
  deleteTableById(id) {
    return Api(secret.LARAVEL_APP_URL).delete(`table/${id}`)
  }
}