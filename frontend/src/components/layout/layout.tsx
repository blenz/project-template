import React, { ReactNode } from 'react'
import { useAuth } from '../../contexts/auth'
import Header from './navbar'

interface LayoutProps {
  children: ReactNode
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  const { user } = useAuth()

  return (
    <div className={`min-h-screen ${user || 'bg-slate-400'}`}>
      {user && <Header />}
      <main className="p-16">{children}</main>
    </div>
  )
}

export default Layout
