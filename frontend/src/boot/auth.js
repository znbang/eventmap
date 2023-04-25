import { boot } from 'quasar/wrappers'
import { reactive } from 'vue'
import { authService } from 'src/lib/service'

const state = reactive({
  user: null
})

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

  await fetchUser().catch(() => {})
})
