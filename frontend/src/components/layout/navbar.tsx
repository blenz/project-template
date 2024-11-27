import React from 'react'
import { Link } from 'react-router'
import { useAuth } from '../../contexts/auth'

const NavBar: React.FC = () => {
  const { logout } = useAuth()

  return (
    <header className="bg-slate-600 px-24 py-3">
      <nav className="flex items-center justify-between text-white">
        <Link className="text-4xl font-semibold" to="/">
          Example
        </Link>

        <div className="flex space-x-6 text-xl">
          <Link to="/">Test</Link>
          <Link to="/about">About</Link>
          <Link to="/login" onClick={() => logout()}>
            Logout
          </Link>
        </div>
      </nav>
    </header>
  )
}

export default NavBar
