import React from 'react'
import { Navigate, Outlet } from 'react-router'
import { useAuth } from '../contexts/auth'

const ProtectedRoutes: React.FC = () => {
  const { loggedIn } = useAuth()
  return loggedIn ? <Outlet /> : <Navigate to="/login" />
}

export default ProtectedRoutes
