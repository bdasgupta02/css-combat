import React, { createContext, useContext, ReactNode } from 'react'
import axios from 'axios'
import { useAppSelector } from '../../app/hooks'
import { userState } from '../../features/user/userSlice'
import TransportAuth from './TransportAuth'

type ContextType = {
  auth: TransportAuth
}

export const TransportContext = createContext({} as ContextType)

export const useTransport = () => {
  return useContext(TransportContext)
}

// TODO set redux state in page level
export const TransportProvider = ({ children }: { children: ReactNode }) => {
  const { jwt } = useAppSelector(userState)

  const clientApiGateway = axios.create({
    baseURL: 'http://localhost:8010/',
    headers: jwt
      ? {
          Authorization: `BEARER ${jwt}`,
        }
      : {},
  })

  const auth = new TransportAuth(clientApiGateway)

  const value: ContextType = {
    auth,
  }

  return <TransportContext.Provider value={value}>{children}</TransportContext.Provider>
}
