import { User } from '../contexts/auth'
import { apiInstance } from '../lib/axios'

export const api = {
  auth: {
    login: async (username: string, password: string): Promise<User> => {
      const resp = await apiInstance.post<User>('/auth/login', { username, password })
      return resp.data
    },
    logout: async (): Promise<void> => {
      await apiInstance.get('/auth/logout')
    },
    session: async (): Promise<User> => {
      const resp = await apiInstance.get<User>('/auth/session')
      return resp.data
    },
    launch: async () => {
      await apiInstance.get('/auth/launch')
    },
  },
}
