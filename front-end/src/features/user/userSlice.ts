import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { RootState } from '../../app/store'

export type UserInfo = {
  jwt?: string
  username?: string
  userId?: number
  isSignedIn: boolean
}

export const userSlice = createSlice({
  name: 'user',
  initialState: {
    jwt: localStorage.getItem('jwt'),
    username: localStorage.getItem('username'),
    userId:
      'userId' in localStorage && localStorage.getItem('userId') === '' ? Number(localStorage.getItem('userId')) : null,
    isSignedIn:
      'isSignedIn' in localStorage
        ? localStorage.getItem('isSignedIn') !== '' && localStorage.getItem('isSignedIn') === 'true'
        : false,
  } as UserInfo,
  reducers: {
    setUserState: (state, action: PayloadAction<UserInfo>) => {
      state.jwt = action.payload.jwt
      state.userId = action.payload.userId
      state.username = action.payload.username
      state.isSignedIn = action.payload.isSignedIn

      localStorage.setItem('isSignedIn', `${state.isSignedIn}`)
      state.jwt ? localStorage.setItem('jwt', state.jwt) : localStorage.removeItem('jwt')
      state.username ? localStorage.setItem('username', state.username) : localStorage.removeItem('username')
      state.userId ? localStorage.setItem('userId', String(state.userId)) : localStorage.removeItem('userId')
    },
  },
})

export const { setUserState } = userSlice.actions
export const userState = (state: RootState) => state.user

export default userSlice.reducer
