import Cookies from 'js-cookie'
import React, { createContext, ReactNode, useContext, useEffect, useState } from 'react'
import { api } from '../services/api'

interface AuthContextType {
  hasSession: boolean
  login: Function
  logout: Function
}

interface AuthProviderProps {
  children: ReactNode
}

const AuthContext = createContext<AuthContextType>({
  hasSession: false,
  login: () => {},
  logout: () => {},
})

export const useAuth = () => useContext(AuthContext)

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [hasSession, setHasSession] = useState<boolean>(!!Cookies.get('session'))

  useEffect(() => {
    ;(async () => {
      try {
        await api.auth.session()
      } catch {
        setHasSession(false)
      }
    })()
  }, [])

  const login = async (username: string, password: string) => {
    await api.auth.login(username, password)
    setHasSession(true)
  }

  const logout = async () => {
    try {
      await api.auth.logout()
    } finally {
      setHasSession(false)
    }
  }

  return <AuthContext.Provider value={{ hasSession, logout, login }}>{children}</AuthContext.Provider>
}
