import React from 'react'
import { Link } from 'react-router'

const NavBar: React.FC = () => {
  return (
    <header className="bg-slate-600 px-12 py-3">
      <nav className="flex items-center text-white">
        <div className="justify-start">
          <Link className="text-xl font-semibold" to="/">
            Example
          </Link>
        </div>

        <div className="justify-end">
          <Link to="/">Test</Link>
          <Link to="/about" state={{ name: 'test' }}>
            About
          </Link>
        </div>
      </nav>
    </header>
  )
}

export default NavBar
