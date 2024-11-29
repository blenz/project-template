import axios from 'axios'
import Cookies from 'js-cookie'
import { useState } from 'react'
import { browserCookie, User } from '../contexts/auth'

const client = axios.create({
  baseURL: '/api',
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

client.interceptors.response.use(
  response => response,
  error => {
    if (window.location.pathname !== '/login' && error.status === 401) {
      browserCookie.remove()
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

const useApi = () => {
  const [loading, setLoading] = useState(false)

  const request = async (func: Function, ...args: any[]) => {
    try {
      setLoading(true)
      const resp = await func(...args)
      return resp.data
    } finally {
      setLoading(false)
    }
  }

  const api = {
    auth: {
      login: async (username: string, password: string): Promise<User> => {
        const resp = await request(client.post, '/auth/login', { username, password })
        return resp.data as User
      },
      logout: async (): Promise<void> => {
        await request(client.get, '/auth/logout')
      },
      session: async (): Promise<User> => {
        const resp = await request(client.get, '/auth/session')
        return resp.data as User
      },
    },
  }

  return { api, loading }
}

export default useApi
