import { BrowserRouter, Route, Routes } from 'react-router'
import Layout from './components/layout/layout'
import ProtectedRoutes from './components/protected-routes'
import { AuthProvider } from './contexts/auth'
import HomePage from './pages/home'
import LoginPage from './pages/login'

function App() {
  return (
    <>
      <BrowserRouter>
        <AuthProvider>
          <Layout>
            <Routes>
              <Route path="/login" element={<LoginPage />} />

              <Route element={<ProtectedRoutes />}>
                <Route path="" element={<HomePage />} />
              </Route>
            </Routes>
          </Layout>
        </AuthProvider>
      </BrowserRouter>
    </>
  )
}

export default App
