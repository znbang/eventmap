import { boot } from 'quasar/wrappers'
import VueGtag from 'vue-gtag'

export default boot(({ app, router }) => {
  app.use(VueGtag, {
    config: {
      id: 'G-VC1QQDMQHJ'
    }
  }, router)
})
