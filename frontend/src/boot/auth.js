import { boot } from 'quasar/wrappers'
import { Notify } from 'quasar'
import axios from 'axios'
import { reactive } from 'vue'
import { authService } from 'components/service'

const state = reactive({
  user: null
})

export function login(provider, accessToken) {
  return axios.post('/api/login', { provider, accessToken })
    .then(() => {
      return fetchUser()
    })
}

export function logout() {
  return authService.logout({})
    .then(() => {
      state.user = null
    })
}

export function fetchUser() {
  return authService.getUser({})
    .then(response => {
      state.user = response
    })
}

export function isAuthenticated() {
  return state.user !== null
}

export function currentUser() {
  return state.user
}

export default boot(async ({ router }) => {
  router.beforeEach((to, from) => {
    if (to.path === '/_=_') {
      return '/'
    } else if (to.matched.some(record => record.meta.auth) && !isAuthenticated()) {
      return '/login'
    }
  })

  axios.interceptors.request.use(config => {
    if (state.user?.token) {
      config.headers.Authorization = 'Bearer ' + state.user.token
    }
    return config
  })

  axios.interceptors.response.use(response => {
    return response
  }, error => {
    if (error.response?.status === 401) {
      if (error.config?.url !== '/api/user') {
        state.user = null
        router.push('/login')
      }
    } else if (error.response?.status !== 400) {
      Notify.create({
        type: 'negative',
        message: error.message
      })
    }
    return Promise.reject(error)
  })

  await fetchUser().catch(() => {})
})
