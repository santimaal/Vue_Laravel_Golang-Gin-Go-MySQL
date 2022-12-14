import Api from '@/services/Api'
import secret from '../secret.js'

export default {
  getAllThematic() {
    return Api(secret.GO_APP_URL).get(`thematic`)
  },
  getThematicById(id) {
    return Api(secret.GO_APP_URL).get(`thematic/${id}`)
  },
  GetThematicInfinite(page, limit) {
    return Api(secret.GO_APP_URL).get('thematic/infinite', { params: { offset: page, limit: limit } });
  },
  createThematic(data) {
    return Api(secret.LARAVEL_APP_URL).post('thematic', data)
  },
  updateThematic(data, id) {
    return Api(secret.LARAVEL_APP_URL).put(`thematic/${id}`, data)
  },
  changeStatusThematic(data, id) {
    return Api(secret.LARAVEL_APP_URL).put(`thematic/changeStatus/${id}`, data)
  },
  deleteThematicById(id) {
    return Api(secret.LARAVEL_APP_URL).delete(`thematic/${id}`)
  }
}