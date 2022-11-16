import axios from 'axios'
import React, { useState } from 'react'
import { Navigate, useNavigate } from 'react-router-dom'
import { useTransport } from '../../api/rest/TransportContext'
import { useAppDispatch, useAppSelector } from '../../app/hooks'
import Button from '../../components/Button'
import Input from '../../components/Input'
import { Logo } from '../../components/Logo'
import { setUserState, userState } from '../../features/user/userSlice'
import jwtDecoder from '../../helpers/jwtDecoder'
import Colors from '../../tokens/Colors'

// TODO for sign up check if username contains any special characters (shouldn't)
type SignInDetails = {
  type: string
  identifier: string
  password: string
}

// TODO clean up later
// TODO loading button when signing in
const SignIn = () => {
  const navigate = useNavigate()
  const { isSignedIn } = useAppSelector(userState)
  const dispatch = useAppDispatch()
  const { auth } = useTransport()
  const [details, setDetails] = useState<SignInDetails>({
    type: '',
    identifier: '',
    password: '',
  })
  
  const changeDetails = (e: any, field: string) => {
    setDetails({
      ...details,
      [field]: e.target.value,
    })

    if (field === 'identifier') {
      const isEmail = String(e.target.value).match(
        /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
      )

      setDetails({
        ...details,
        type: isEmail ? 'email' : 'username',
        [field]: e.target.value,
      })
    }
  }

  const isValid = () => {}

  const signIn = async () => {
    // TODO validation and type recognition
    // redirect to play (redirect when loading the page also)
    try {
      const resp = await auth.signIn(details.type, details.identifier, details.password)
      const jwtDecoded = jwtDecoder(resp.data.token)
      dispatch(
        setUserState({
          isSignedIn: true,
          jwt: resp.data.token,
          userId: jwtDecoded.userId,
          username: jwtDecoded.username,
        }),
      )

      navigate('/play')
    } catch (e) {}
    //const jwt = resp.data.token
  }

  // detect if email or username
  return isSignedIn ? (
    <Navigate to="/play" />
  ) : (
    <div
      style={{
        height: '100vh',
        width: '100vw',
        backgroundColor: Colors.Background,
        position: 'relative',
      }}>
      <div
        style={{
          zIndex: 3,
          position: 'absolute',
          backgroundColor: Colors.SignInFrontBox,
          height: '33%',
          width: '50%',
          top: 0,
          left: 0,
        }}
      />
      <div
        style={{
          zIndex: 3,
          position: 'absolute',
          backgroundColor: Colors.SignInFrontBox,
          height: '53%',
          width: '70%',
          bottom: 0,
          right: 0,
        }}
      />
      <div
        style={{
          zIndex: 1,
          position: 'absolute',
          backgroundColor: Colors.SignInBackBox,
          height: '70%',
          width: '70%',
          top: 0,
          bottom: 0,
          left: 0,
          right: 0,
          margin: 'auto',
        }}
      />
      <div
        style={{
          zIndex: 5,
          position: 'absolute',
          backgroundColor: Colors.SignInDetails,
          width: '440px',
          right: '200px',
          top: 0,
          bottom: 0,
          marginTop: 'auto',
          marginBottom: 'auto',
          height: '650px',
          backdropFilter: 'blur(10px)',
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}>
        <div style={{ height: '90px' }} />
        <Logo />
        <Input
          outerStyle={{
            marginTop: '90px',
            width: '75%',
          }}
          value={details.identifier}
          onChange={(e: any) => changeDetails(e, 'identifier')}
          placeholder={'Email or username'}
        />
        <Input
          outerStyle={{
            marginTop: '24px',
            width: '75%',
          }}
          value={details.password}
          onChange={(e: any) => changeDetails(e, 'password')}
          placeholder={'Password'}
          innerProps={{type: 'password'}}
        />
        <Button
          onClick={signIn}
          text="Enter >"
          outerStyle={{ width: '180px', marginTop: '100px' }}
        />
        <div style={{ color: 'white', marginTop: '20px' }}>Or create an account here</div>
      </div>
    </div>
  )
}

export default SignIn
