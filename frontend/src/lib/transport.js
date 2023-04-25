import { createConnectTransport } from '@bufbuild/connect-web'
import { currentUser} from 'boot/auth'

const auth = (next) => async (req) => {
  const user = currentUser()
  if (user?.token) {
    req.header.set('Authorization', 'Bearer ' + user.token)
  }
  return await next(req)
}

const transport = createConnectTransport({
  baseUrl: '/api',
  useBinaryFormat: true,
  interceptors: [auth]
})

export function useTransport() {
  return transport
}
