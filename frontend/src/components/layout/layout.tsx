import React, { ReactNode } from 'react'
import { useAuth } from '../../contexts/auth'
import Header from './navbar'

interface LayoutProps {
  children: ReactNode
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  const { user } = useAuth()

  if (!user) {
    return (
      <div className="min-h-screen bg-slate-400">
        <main>{children}</main>
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
