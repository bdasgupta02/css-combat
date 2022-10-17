import React from 'react'
import { Navigate } from 'react-router-dom'
import { useAppSelector } from './app/hooks'
import { userState } from './features/user/userSlice'

const PrivateRoute = ({ element }: { element: JSX.Element }) => {
  const { isSignedIn } = useAppSelector(userState)
  return isSignedIn ? element : <Navigate to={{ pathname: '/' }} />
}

export default PrivateRoute
