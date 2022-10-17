import React from 'react'
import { BrowserRouter as Router, Route, Routes, useLocation } from 'react-router-dom'
import './App.css'
import PrivateRoute from './PrivateRoute'
import { TransportProvider } from './api/rest/TransportContext'
import { Navbar } from './components/Navbar'
import { useAppSelector } from './app/hooks'
import { userState } from './features/user/userSlice'
import Container from './components/Container'
import styled from 'styled-components'
import Colors from './tokens/Colors'
import SignIn from './pages/SignIn'
import Play from './pages/Play'
import Game from './pages/Game'

const Placeholder = () => {
  return <Container>Placeholder</Container>
}

const Pages = () => {
  // to toggle/hide navbar
  const { isSignedIn } = useAppSelector(userState)
  const location = useLocation()

  const inGame = location.pathname.substring(0, 6) === "/play/"

  return (
    <AppOuter>
      {isSignedIn && !inGame && <Navbar />}
      <RouteOuter>
        <Routes>
          <Route path="/profile-settings" element={<PrivateRoute element={<Placeholder />} />} />
          <Route path="/profile" element={<PrivateRoute element={<Placeholder />} />} />
          <Route path="/shop" element={<PrivateRoute element={<Placeholder />} />} />
          <Route path="/inventory" element={<PrivateRoute element={<Placeholder />} />} />
          <Route path="/history" element={<PrivateRoute element={<Placeholder />} />} />
          <Route path="/play/train" element={<PrivateRoute element={<Placeholder />} />} />
          <Route path="/play/:id" element={<PrivateRoute element={<Game />} />} />
          <Route path="/play" element={<PrivateRoute element={<Play />} />} />
          <Route path="/sign-up" element={<Placeholder />} />
          <Route path="/" element={<SignIn />} />
        </Routes>
      </RouteOuter>
    </AppOuter>
  )
}

function App() {
  return (
    <Router>
      <TransportProvider>
        <Pages />
      </TransportProvider>
    </Router>
  )
}

const RouteOuter = styled.div`
  background-color: ${Colors.DarkBg};
  flex: 1;
`

const AppOuter = styled.div`
  min-height: 100vh;
  display: flex;
  flex-direction: column;
`

export default App
