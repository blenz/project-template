import axios from 'axios'
import { User } from '../../contexts/auth'

const client = axios.create({
  baseURL: 'http://localhost:3000/api',
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
  },
})

export const api = {
  auth: {
    login: async (username: string, password: string): Promise<User> => {
      const resp = await client.post('/auth/login', { username, password })
      return resp.data as User
    },
    logout: async (): Promise<void> => {
      await client.get('/auth/logout')
    },
    session: async (): Promise<User> => {
      const resp = await client.get('/auth/session')
      return resp.data as User
    },
  },
}
