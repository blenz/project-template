import { BrowserRouter, Route, Routes } from 'react-router'
import Layout from './components/layout/layout'
import HomePage from './pages/home'

function App() {
  return (
    <>
      <BrowserRouter>
        <Layout>
          <Routes>
            <Route path="/" element={<HomePage />} />
          </Routes>
        </Layout>
      </BrowserRouter>
    </>
  )
}

export default App
