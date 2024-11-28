import axios from 'axios'
import Cookies from 'js-cookie'

const client = axios.create({
  baseURL: 'http://localhost:3000/api',
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
  },
})

client.interceptors.request.use(
  config => {
    config.headers = config.headers || {}
    config.headers['X-CSRF-TOKEN'] = Cookies.get('_csrf')
    return config
  },
  error => Promise.reject(error)
)

export const api = {
  auth: {
    login: async (username: string, password: string): Promise<any> => {
      const resp = await client.post('/auth/login', { username, password })
      return resp.data
    },
    logout: async (): Promise<void> => {
      await client.get('/auth/logout')
    },
    session: async (): Promise<any> => {
      const resp = await client.get('/auth/session')
      return resp.data
    },
  },
}
