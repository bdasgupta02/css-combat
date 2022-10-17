import React, { useEffect, useState } from 'react'
import Button from '../../components/Button'
import { Heading } from '../../components/Text'
import useWebSocket from 'react-use-websocket'
import { matchmakingUrl, matchmakingCodes } from '../../api/ws/Websocket'
import { useAppSelector } from '../../app/hooks'
import { userState } from '../../features/user/userSlice'
import { useNavigate } from 'react-router-dom'
import { percentageToHex } from '../../tokens/Colors'

// TODO temporary button for now
// TODO move state to redux for matchmaking and WS (localstorage for ongoing game check with a query too)
// Get rank func
const Game = () => {
  const navigate = useNavigate()
  const { jwt } = useAppSelector(userState)

  // ui
  const [socketUrl, setSocketUrl] = useState<any>()
  const [isSearching, setSearching] = useState(false)

  // ws
  const [messageHistory, setMessageHistory] = useState([] as string[])
  const { sendMessage, lastMessage } = useWebSocket(socketUrl ? socketUrl : '')

  useEffect(() => {
    setSocketUrl(isSearching ? matchmakingUrl : null)
  }, [isSearching])

  useEffect(() => {
    sendMessage(`100 ${jwt}`)
  }, [socketUrl])

  useEffect(() => {
    if (lastMessage) {
      setMessageHistory([...messageHistory, lastMessage.data])
    }
  }, [lastMessage])

  useEffect(() => {
    if (messageHistory.length > 0 && messageHistory[messageHistory.length - 1]) {
      const respCode = messageHistory[messageHistory.length - 1].substring(0, 3)
      if (respCode == matchmakingCodes.timedOut) {
        setSearching(false)
        setSocketUrl('')
      } else if (respCode == matchmakingCodes.found) {
        const gameRoom = messageHistory[messageHistory.length - 1].substring(4)
        navigate(`/play/${gameRoom}`)
      }
    }
  }, [messageHistory])

  const searchButtonStyle = isSearching
    ? {
        width: '200px',
        backgroundColor: `#09FF6B${percentageToHex(54)}`,
        boxShadow: `3px 10px 10px #09FF6B${percentageToHex(15)}`,
        outline: `4px solid #09FF6B${percentageToHex(70)}`,
        fontWeight: 'normal',
      }
    : { width: '200px' }

  return (
    <div>
      <Heading>Competitive mode</Heading>
      <div>{'<rank, game and other info here>'}</div>
      <div style={{ marginTop: '20px' }} />
      <Button
        text={isSearching ? 'Searching' : 'Play >'}
        onClick={() => setSearching(!isSearching)}
        outerStyle={searchButtonStyle}
      />
      <div></div>
    </div>
  )
}

export default Game
