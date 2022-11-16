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
import Colors, { percentageToHex } from '../../tokens/Colors'

// TODO for sign up check if username contains any special characters (shouldn't)
type SignInDetails = {
  type: string
  identifier: string
  password: string
}

type SignUpDetails = {
  email: string
  username: string
  password: string
  fullName: string
  confirm: string
}

// TODO clean up later
// TODO loading button when signing in
const SignIn = () => {
  const [isSignInMode, setSignInMode] = useState(true)

  const navigate = useNavigate()
  const { isSignedIn } = useAppSelector(userState)
  const dispatch = useAppDispatch()
  const { auth } = useTransport()
  const [signInDetails, setSignInDetails] = useState<SignInDetails>({
    type: '',
    identifier: '',
    password: '',
  })

  const [signUpDetails, setSignUpDetails] = useState<SignUpDetails>({
    email: '',
    username: '',
    password: '',
    fullName: '',
    confirm: '',
  })

  const changeDetails = (e: any, field: string) => {
    setSignInDetails({
      ...signInDetails,
      [field]: e.target.value,
    })

    if (field === 'identifier') {
      const isEmail = String(e.target.value).match(
        /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
      )

      setSignInDetails({
        ...signInDetails,
        type: isEmail ? 'email' : 'username',
        [field]: e.target.value,
      })
    }
  }

  const changeSignUpDetails = (e: any, field: string) => {
    setSignUpDetails({
      ...signUpDetails,
      [field]: e.target.value
    })
  }

  const isValid = () => {}

  const signIn = async () => {
    // TODO validation and type recognition
    // redirect to play (redirect when loading the page also)
    try {
      const resp = await auth.signIn(signInDetails.type, signInDetails.identifier, signInDetails.password)
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

  const signUp = async () => {
    if (signUpDetails.password != '' && signUpDetails.confirm !== signUpDetails.password)
      return
      
    try {
      const resp = await auth.signUp(signUpDetails.email, signUpDetails.username, signUpDetails.fullName, signUpDetails.password)
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
          position: 'absolute',
          width: '100vw',
          height: '100vh',
          zIndex: 20,
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          fontFamily: 'Fira Code',
          textAlign: 'left',
          fontSize: '26px',
        }}>
        <div style={{ width: '900px', marginRight: '200px' }}>
          <div style={{ color: `#ffffff${percentageToHex(15)}` }}>
            01&nbsp;&nbsp;&nbsp;{'<div>'}
          </div>
          <div style={{ color: `#ffffff${percentageToHex(15)}` }}>
            02&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{'<p>'}
          </div>
          <div style={{ color: `#ffffff${percentageToHex(15)}` }}>
            03&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
            <span style={{ color: `#ffffff` }}>{'Hello World.'}</span>
          </div>
          <div style={{ color: `#ffffff${percentageToHex(15)}` }}>
            04&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
            <span style={{ color: `#ffffff` }}>{'Use your'}</span>
            <span style={{ color: `#5488D8` }}>{' CSS Skills '}</span>
            <span style={{ color: `#ffffff` }}>{'to compete_'}</span>
          </div>
          <div style={{ color: `#ffffff${percentageToHex(15)}` }}>
            05&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{'</p>'}
          </div>
          <div style={{ color: `#ffffff${percentageToHex(15)}` }}>
            06&nbsp;&nbsp;&nbsp;{'</div>'}
          </div>
        </div>
      </div>
      {isSignInMode ? (
        <div
          style={{
            zIndex: 100,
            position: 'absolute',
            backgroundColor: Colors.SignInDetails,
            width: '440px',
            right: '350px',
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
            value={signInDetails.identifier}
            onChange={(e: any) => changeDetails(e, 'identifier')}
            placeholder={'Email or username'}
          />
          <Input
            outerStyle={{
              marginTop: '24px',
              width: '75%',
            }}
            value={signInDetails.password}
            onChange={(e: any) => changeDetails(e, 'password')}
            placeholder={'Password'}
            innerProps={{ type: 'password' }}
          />
          <Button
            onClick={signIn}
            text="Enter >"
            outerStyle={{ width: '180px', marginTop: '100px' }}
          />
          <div style={{ color: 'white', marginTop: '20px', cursor: 'pointer' }} onClick={() => setSignInMode(false)}>
            Or create an account here
          </div>
        </div>
      ) : (
        <div
          style={{
            zIndex: 100,
            position: 'absolute',
            backgroundColor: Colors.SignInDetails,
            width: '440px',
            right: '350px',
            top: 0,
            bottom: 0,
            marginTop: 'auto',
            marginBottom: 'auto',
            height: '850px',
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
            value={signUpDetails.email}
            onChange={(e: any) => changeSignUpDetails(e, 'email')}
            placeholder={'Email'}
          />
          <Input
            outerStyle={{
              marginTop: '24px',
              width: '75%',
            }}
            value={signUpDetails.fullName}
            onChange={(e: any) => changeSignUpDetails(e, 'fullName')}
            placeholder={'Name'}
          />
          <Input
            outerStyle={{
              marginTop: '24px',
              width: '75%',
            }}
            value={signUpDetails.username}
            onChange={(e: any) => changeSignUpDetails(e, 'username')}
            placeholder={'Username'}
          />
          <Input
            outerStyle={{
              marginTop: '24px',
              width: '75%',
            }}
            value={signUpDetails.password}
            onChange={(e: any) => changeSignUpDetails(e, 'password')}
            placeholder={'Password'}
            innerProps={{ type: 'password' }}
          />
          <Input
            outerStyle={{
              marginTop: '24px',
              width: '75%',
            }}
            value={signUpDetails.confirm}
            onChange={(e: any) => changeSignUpDetails(e, 'confirm')}
            placeholder={'Confirm password'}
            innerProps={{ type: 'password' }}
          />
          <Button
            onClick={signUp}
            text="Create an account >"
            outerStyle={{ width: '220px', marginTop: '100px' }}
          />
          <div style={{ color: 'white', marginTop: '20px', cursor: 'pointer' }} onClick={() => setSignInMode(true)}>
            Or sign-in here
          </div>
        </div>
      )}
    </div>
  )
}

export default SignIn
