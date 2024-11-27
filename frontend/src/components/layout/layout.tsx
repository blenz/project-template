import React, { ReactNode } from 'react'
import { useAuth } from '../../contexts/auth'
import Header from './navbar'

interface LayoutProps {
  children: ReactNode
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  const { user } = useAuth()

  return (
    <div>
      {user && <Header />}
      <main>{children}</main>
    </div>
  )
}

export default Layout
