import Cookies from 'js-cookie'
import React, { createContext, ReactNode, useContext, useEffect, useState } from 'react'
import { useNavigate } from 'react-router'
import useApi from '../hooks/use-api'

export const browserCookie = {
  set: () => Cookies.set('_auth', '1'),
  remove: () => Cookies.remove('_auth'),
  exists: () => !!Cookies.get('_auth'),
}

export interface User {
  username: string
}

interface AuthContextType {
  authed: boolean
  login: Function
  logout: Function
}

interface AuthProviderProps {
  children: ReactNode
}

const AuthContext = createContext<AuthContextType>({
  authed: false,
  login: () => {},
  logout: () => {},
})

export const useAuth = () => useContext(AuthContext)

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const { api } = useApi()
  const navigate = useNavigate()

  const [authed, setAuthed] = useState<boolean>(browserCookie.exists())

  useEffect(() => {
    ;(async () => {
      await api.auth.session()
      browserCookie.set()
      setAuthed(true)
    })()
  }, [])

  const login = async (username: string, password: string) => {
    await api.auth.login(username, password)
    browserCookie.set()
    setAuthed(true)
  }

  const logout = async () => {
    try {
      await api.auth.logout()
    } finally {
      browserCookie.remove()
      setAuthed(false)
      navigate('/login')
    }
  }

  return <AuthContext.Provider value={{ authed, logout, login }}>{children}</AuthContext.Provider>
}
