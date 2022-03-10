import { boot } from 'quasar/wrappers'
import { createI18n } from 'vue-i18n'
import messages from 'src/i18n'

const i18n = createI18n({
  globalInjection: true,
  legacy: false,
  locale: 'zh-TW',
  messages
})

export default boot(({ app }) => {
  // Set i18n instance on app
  app.use(i18n)
})

export { i18n }

export function tr(errors) {
  return Object.fromEntries(Object.entries(errors).map(entry => [entry[0], i18n.global.t(entry[1])]))
}
