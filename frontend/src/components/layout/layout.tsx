import React, { ReactNode } from 'react'
import { useAuth } from '../../contexts/auth'
import Header from './navbar'
import Login from '../../pages/login'

interface LayoutProps {
  children: ReactNode
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  const { authed } = useAuth()

  if (!authed) {
    return (
      <div className="min-h-screen bg-slate-400">
        <main>
          <Login />
        </main>
      </div>
    )
  }

  return (
    <div className="min-h-screen">
      <Header />
      <main className="p-16">{children}</main>
    </div>
  )
}

export default Layout
