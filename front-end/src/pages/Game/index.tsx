import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import useWebSocket from 'react-use-websocket'
import { gameplayUrl as base, gameplayCodes } from '../../api/ws/Websocket'
import { useAppSelector } from '../../app/hooks'
import { userState } from '../../features/user/userSlice'
import Editor from 'react-simple-code-editor'
import Prism from 'prismjs'
import DOMPurify from 'dompurify'
import 'prismjs/components/prism-clike'
import 'prismjs/components/prism-javascript'
import './themes/vsc-theme.css'
import './editor.css'
import Colors, { percentageToHex } from '../../tokens/Colors'
import styled from 'styled-components'
import { Heading } from '../../components/Text'

// svg? maybe can have svg vs. other things

// TODO iframe does not work with IE (alternative is dangerousely set inner, render image then show that)
// TODO svg banner
// TODO test security with writing API calls directly within editor
// TODO create editor component separately
// TODO splash loading screen
// TODO custom prism theme is in FYP bookmarks
// TODO color picker
// TODO settings
// TODO mention 300 x 400 canvas
// TODO templates for different types

// TODO wait for stop message rather than use timer from server

type DisplayTime = {
  m: number
  s: number
}

const TopSection = ({ title, body }: { title: string; body: string }) => {
  return (
    <div style={{ marginLeft: '30px' }}>
      <div
        style={{
          fontFamily: 'League Spartan',
          fontSize: '14px',
        }}>
        {title.toUpperCase()}
      </div>
      <div style={{ fontSize: '28px' }}>{body}</div>
    </div>
  )
}

// just using placeholders now
const AvatarBox = ({ username }: { username: string }) => {
  return (
    <div
      style={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        height: '100%',
        justifyContent: 'center',
        width: 'fit-content',
        marginLeft: '15px',
      }}>
      <div style={{ height: '40px', width: '40px', backgroundColor: '#FFFFFF20' }} />
      <div>{username}</div>
    </div>
  )
}

const CodeBit = ({ children }: { children: string }) => {
  return (
    <span
      style={{
        backgroundColor: '#373c49',
        padding: '2px 6px 4px 6px',
        borderRadius: '8px',
        fontStyle: 'fira-code',
        fontSize: '14px',
        color: `#ffffff${percentageToHex(80)}`,
        marginRight: '4px',
      }}>
      {children}
    </span>
  )
}

const Game = () => {
  const [prevTime, setPrevTime] = useState<string>('')
  const [targetImg, setTargetImg] = useState<string>('')
  const [currTime, setCurrTime] = useState<DisplayTime>({ m: 0, s: 0 })
  const [editorVal, setEditorVal] = useState<string>(
    '<!--Write your solution here-->\n<div></div>\n<style></style>',
  )
  const { jwt } = useAppSelector(userState)

  const params = useParams()
  const { sendMessage, lastMessage } = useWebSocket(`${base}/${params.id}`)
  const [isStarted, setStarted] = useState(false)

  useEffect(() => {
    sendMessage && sendMessage(`${gameplayCodes.connected} ${jwt}`)
  }, [])

  useEffect(() => {
    if (lastMessage) {
      const respCode = lastMessage.data.substring(0, 3)
      if (respCode === gameplayCodes.start) {
        setStarted(true)
        setTargetImg(lastMessage.data.substring(8))
      } else if (respCode === gameplayCodes.timeUpdate) {
        const respTime = lastMessage.data.substring(4)
        if (respTime !== prevTime) {
          if (currTime.s === 59) {
            setCurrTime({ m: currTime.m + 1, s: 0 })
          } else {
            setCurrTime({ m: currTime.m, s: currTime.s + 1 })
          }
          setPrevTime(respTime)
        }
      }
    }
  }, [lastMessage])

  const hightlightWithLineNumbers = (input: string) =>
    Prism.highlight(input, Prism.languages['html'], 'html')
      .split('\n')
      .map((line, i) => `<span class='editorLineNumber'>${i + 1}</span>${line}`)
      .join('\n')

  // TODO disable settings and command pallete
  // TODO splash loading screen
  return isStarted ? (
    <div style={{ height: '100vh' }}>
      <div
        style={{
          height: '70px',
          backgroundColor: '#282f3d',
          display: 'flex',
          flexDirection: 'row',
          width: '100%',
          color: 'white',
          alignItems: 'center',
        }}>
        <TopSection title="Time" body={`${String(currTime.m).padStart(2, '0')}:${String(currTime.s).padStart(2, '0')}`} />
        <TopSection title="Accuracy" body={'0%'} />
        <TopSection title="Max accuracy" body={'0%'} />
        <div
          style={{
            flex: 1,
            height: '100%',
            display: 'flex',
            flexDirection: 'row',
            justifyContent: 'flex-end',
          }}>
          <AvatarBox username="test1" />
          <AvatarBox username="test3" />
          <div style={{ width: '30px', height: '1px' }} />
        </div>
      </div>
      <div style={{ display: 'flex', flexDirection: 'row', color: 'white', height: '100%' }}>
        <div style={{ flex: 4, backgroundColor: '#161a24' }}>
          <Editor
            highlight={code => hightlightWithLineNumbers(code)}
            value={editorVal}
            onValueChange={e => setEditorVal(e)}
            padding={10}
            textareaId="codeArea"
            className="editor"
            style={{
              fontFamily: '"Fira code", "Fira Mono", monospace',
              fontSize: '14px',
              color: 'white',
              border: 'none',
              outline: 0,
              height: '100%',
            }}
          />
        </div>
        <div
          style={{
            flex: 2,
            backgroundColor: '#262A35',
            height: '100%',
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}>
          <div
            style={{
              width: '100%',
              display: 'flex',
              flexDirection: 'row',
              padding: '20px 0px 0px',
              alignItems: 'center',
            }}>
            <div style={{ display: 'flex', flexDirection: 'column', paddingRight: '20px' }}>
              <div
                style={{
                  marginLeft: '20px',
                  fontFamily: 'League Spartan',
                  fontSize: '14px',
                  color: `${Colors.White}${percentageToHex(50)}`,
                }}>
                TYPE
              </div>
              <div
                style={{
                  marginLeft: '20px',
                  fontSize: '26px',
                }}>
                Grid
              </div>
            </div>
            <div style={{ display: 'flex', flexDirection: 'column' }}>
              <div
                style={{
                  marginLeft: '20px',
                  fontFamily: 'League Spartan',
                  fontSize: '14px',
                  color: `${Colors.White}${percentageToHex(50)}`,
                }}>
                USE
              </div>
              <div
                style={{
                  marginLeft: '20px',
                }}>
                <CodeBit>grid*</CodeBit>
                <CodeBit>column-gap</CodeBit>
                <CodeBit>row-gap</CodeBit>
              </div>
            </div>
            <div style={{ display: 'flex', flexDirection: 'column' }}>
              <div
                style={{
                  marginLeft: '20px',
                  fontFamily: 'League Spartan',
                  fontSize: '14px',
                  color: `${Colors.White}${percentageToHex(50)}`,
                }}>
                DON'T USE
              </div>
              <div
                style={{
                  marginLeft: '20px',
                }}>
                <CodeBit>svg</CodeBit>
                <CodeBit>flex</CodeBit>
                <CodeBit>margin*</CodeBit>
                <CodeBit>padding*</CodeBit>
                <CodeBit>table</CodeBit>
              </div>
            </div>
          </div>
          <div style={{ flex: 1 }}>
            <div
              style={{
                marginTop: '30px',
                marginBottom: '30px',
                width: '100%',
                fontFamily: 'League Spartan',
                fontSize: '14px',
                transform: 'translateX(-40px)',
                color: `${Colors.White}${percentageToHex(50)}`,
              }}>
              TARGET IMAGE
            </div>
            <OuterDivTraget
              style={{
                display: 'block',
                transform: 'scale(1.2)',
              }}>
              <img src={targetImg} style={{ width: '100%', height: '100%' }} />
            </OuterDivTraget>
          </div>
          <div style={{ flex: 1 }}>
            <div
              style={{
                marginBottom: '25px',
                width: '100%',
                fontFamily: 'League Spartan',
                fontSize: '14px',
                transform: 'translateX(-30px)',
                color: `${Colors.White}${percentageToHex(50)}`,
              }}>
              YOUR SUBMISSION
            </div>
            <iframe
              style={{
                height: '320px',
                width: '420px',
                overflow: 'hidden',
                display: 'block',
                border: 'hidden',
                transform: 'scale(1.2)',
              }}
              frameBorder={0}
              srcDoc={DOMPurify.sanitize(canvasPrefix + editorVal + canvasPostfix)}
            />
          </div>
        </div>
      </div>
    </div>
  ) : (
    <Heading style={{ padding: '20px' }}>Waiting to start</Heading>
  )
}

const OuterDivTraget = styled.div`
  height: 300px;
  width: 400px;
  overflow: hidden;
  background-color: #ffffff00;
  border: 2px solid #00000040;
`

const canvasPrefix = `<div style="height: 300px; width: 400px; overflow: hidden; background-color: #FFFFFF00; border: 2px solid #00000040;">`
const canvasPostfix = '</div>'

export default Game
