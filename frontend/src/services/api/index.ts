import axios from 'axios'
import { User } from '../../contexts/auth'

const client = axios.create({
  baseURL: 'http://localhost:3000',
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
  },
})

export const api = {
  auth: {
    login: async (username: string, password: string): Promise<void> => {
      await client.post('/api/auth/login', { username, password })
    },
    logout: async (): Promise<void> => {
      await client.get('/api/auth/logout')
    },
    session: async (): Promise<{ id: string; user: User }> => {
      const resp = await client.get('/api/auth/session')
      return { id: '1', user: resp.data as User }
    },
  },
}
