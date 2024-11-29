import { BrowserRouter, Route, Routes } from 'react-router'
import Layout from './components/layout/layout'
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
              <Route path="" element={<HomePage />} />
            </Routes>
          </Layout>
        </AuthProvider>
      </BrowserRouter>
    </>
  )
}

export default App
