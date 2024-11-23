import React, { ReactNode } from 'react'
import Header from './navbar'

interface LayoutProps {
  children: ReactNode
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <div>
      <Header />
      <main style={{ padding: '20px' }}>{children}</main>
    </div>
  )
}

export default Layout
