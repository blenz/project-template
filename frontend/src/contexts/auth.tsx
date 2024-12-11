import Cookies from 'js-cookie'
import React, { createContext, ReactNode, useContext, useEffect, useState } from 'react'
import { useNavigate } from 'react-router'
import { api } from '../services/api'

export const browserCookie = {
  create: () => Cookies.set('_auth', '1'),
  remove: () => Cookies.remove('_auth'),
  exists: () => !!Cookies.get('_auth'),
}

export interface User {
  username: string
}

interface AuthContextType {
  user: User | null
  authed: boolean
  login: Function
  logout: Function
}

interface AuthProviderProps {
  children: ReactNode
}

const AuthContext = createContext<AuthContextType>({
  user: null,
  authed: false,
  login: () => {},
  logout: () => {},
})

export const useAuth = () => useContext(AuthContext)

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const navigate = useNavigate()

  const [authed, setAuthed] = useState<boolean>(browserCookie.exists())
  const [user, setUser] = useState<User | null>(null)

  useEffect(() => {
    ;(async () => {
      const user = await api.auth.session()
      browserCookie.create()
      setAuthed(true)
      setUser(user)
    })()
  }, [])

  const login = async (username: string, password: string) => {
    await api.auth.login(username, password)
    browserCookie.create()
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

  return <AuthContext.Provider value={{ user, authed, logout, login }}>{children}</AuthContext.Provider>
}
