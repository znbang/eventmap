import { createConnectTransport } from '@connectrpc/connect-web'
import { getToken } from 'boot/auth'

const auth = (next) => async (req) => {
  const token = getToken()
  if (token) {
    req.header.set('Authorization', 'Bearer ' + token)
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
