import axios from 'axios'
import { browserCookie } from '../contexts/auth'

export const apiInstance = axios.create({
  baseURL: '/api',
  withCredentials: true,
  xsrfHeaderName: 'X-CSRF-TOKEN',
  xsrfCookieName: '_csrf',
  headers: {
    'Content-Type': 'application/json',
  },
})

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
