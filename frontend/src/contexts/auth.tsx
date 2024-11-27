import React, { createContext, ReactNode, useContext, useEffect, useState } from 'react'
import { useNavigate } from 'react-router'
import { api } from '../services/api'

export interface User {
  username: string
  token: string
}

interface AuthContextType {
  user: User | null
  logout: Function
  login: Function
}

interface AuthProviderProps {
  children: ReactNode
}

const AuthContext = createContext<AuthContextType>({ user: null, login: () => {}, logout: () => {} })

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const navigate = useNavigate()
  const [user, setUser] = useState<User | null>(null)

  useEffect(() => {
    ;(async () => {
      const { user } = await api.auth.session()
      setUser(user)
    })()
  }, [])

  const login = async (username: string, password: string) => {
    await api.auth.login(username, password)
    setUser(user)
    navigate('/')
  }

  const logout = async () => {
    await api.auth.logout()
    setUser(null)
    navigate('/login')
  }

  return <AuthContext.Provider value={{ user, logout, login }}>{children}</AuthContext.Provider>
}

export const useAuth = () => useContext(AuthContext)
