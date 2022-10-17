import React from 'react'
import { useNavigate } from 'react-router-dom'
import { useAppDispatch } from '../../app/hooks'
import { setUserState } from '../../features/user/userSlice'
import Colors from '../../tokens/Colors'
import Button from '../Button'
import NavButton from './NavButton'
import NavButtonDiv from './NavButtonDiv'

export const Navbar = () => {
  const dispatch = useAppDispatch()
  const navigate = useNavigate()

  const onSignOut = () => {
    dispatch(
      setUserState({
        isSignedIn: false,
        jwt: '',
        userId: -1,
        username: '',
      }),
    )

    navigate("/")
  }

  return (
    <div style={{ width: '100%', position: 'relative' }}>
      <div
        style={{
          position: 'absolute',
          backgroundColor: Colors.DarkBg,
          height: '40px',
          width: '100%',
        }}
      />
      <div
        style={{
          paddingTop: '20px',
          paddingBottom: '20px',
          display: 'flex',
          flexDirection: 'row',
          position: 'absolute',
          width: '100%',
        }}>
        <div style={{ width: '60px', height: '1px' }} />
        <NavButton text="Play" route="/play" />
        <NavButton text="History" route="/history" />
        <NavButton text="Inventory" route="/inventory" />
        <NavButton text="Shop" route="/shop" />
        <div style={{ flex: 1 }} />
        <NavButtonDiv text="Sign-out" onClick={onSignOut} />
        <div style={{ width: '60px', height: '1px' }} />
      </div>
      <div style={{ height: '80px', width: '1px' }} />
    </div>
  )
}
