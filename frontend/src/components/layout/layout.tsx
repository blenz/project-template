import React, { ReactNode } from 'react'
import Header from './navbar'

interface LayoutProps {
  children: ReactNode
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <div>
      <Header />
      <main className="m-10 border-2 border-green-500">{children}</main>
    </div>
  )
}

export default Layout
