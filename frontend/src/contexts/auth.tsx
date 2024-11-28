import Cookies from 'js-cookie'
import React, { createContext, ReactNode, useContext, useEffect, useState } from 'react'
import { api } from '../services/api'

interface AuthContextType {
  loggedIn: boolean
  login: Function
  logout: Function
}

interface AuthProviderProps {
  children: ReactNode
}

const AuthContext = createContext<AuthContextType>({
  loggedIn: false,
  login: () => {},
  logout: () => {},
})

export const useAuth = () => useContext(AuthContext)

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [loggedIn, setLoggedIn] = useState<boolean>(!!Cookies.get('token'))

  useEffect(() => {
    ;(async () => {
      try {
        await api.auth.session()
      } catch {
        setLoggedIn(false)
      }
    })()
  }, [])

  const login = async (username: string, password: string) => {
    await api.auth.login(username, password)
    setLoggedIn(true)
  }

  const logout = async () => {
    await api.auth.logout()
    setLoggedIn(false)
  }

  return <AuthContext.Provider value={{ loggedIn, logout, login }}>{children}</AuthContext.Provider>
}
