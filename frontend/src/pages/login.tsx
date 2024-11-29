import { FormEvent, useState } from 'react'
import { Navigate } from 'react-router'
import { Button } from '../components/ui/button'
import { Input } from '../components/ui/input'
import { useAuth } from '../contexts/auth'

const LoginPage = () => {
  const { authed, login } = useAuth()

  const [username, setUsername] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const [error, setError] = useState<string | null>(null)

  if (authed) return <Navigate to="/" />

  const handleLogin = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()

    try {
      window.location.href = 'http://localhost:3000/api/auth/launch'
    } catch {
      setError('Invalid Username or Password')
    }
  }

  return (
    <div className="flex min-h-screen items-center justify-center">
      <div className="w-full max-w-md space-y-6 rounded-lg bg-white p-8 shadow-md">
        <h2 className="text-center text-2xl font-semibold text-gray-800">Login</h2>

        {error && <p className="text-red-500">{error}</p>}

        <form className="space-y-4" onSubmit={handleLogin}>
          <div>
            <label className="block text-sm font-medium text-gray-700">Username</label>
            <Input
              id="username"
              placeholder="Enter your username"
              className="mt-1"
              onChange={e => setUsername(e.target.value)}
            />
          </div>
          <div>
            <label htmlFor="password" className="block text-sm font-medium text-gray-700">
              Password
            </label>
            <Input
              type="password"
              id="password"
              placeholder="Enter your password"
              className="mt-1"
              onChange={e => setPassword(e.target.value)}
            />
          </div>
          <Button className="mt-4 w-full">Login</Button>
        </form>

        {/* <p className="text-center text-sm text-gray-600">
          Don't have an account?{' '}
          <a href="/register" className="text-blue-500 hover:underline">
            Sign up
          </a>
        </p> */}
      </div>
    </div>
  )
}

export default LoginPage
