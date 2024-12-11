import axios from 'axios'
import Cookies from 'js-cookie'
import { browserCookie } from '../contexts/auth'

export const apiInstance = axios.create({
  baseURL: '/api',
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
  },
})

apiInstance.interceptors.request.use(
  config => {
    config.headers = config.headers || {}
    config.headers['X-CSRF-TOKEN'] = Cookies.get('_csrf')
    return config
  },
  error => Promise.reject(error)
)

apiInstance.interceptors.response.use(
  response => response,
  error => {
    if (window.location.pathname !== '/login' && error.status === 401) {
      browserCookie.remove()
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)
