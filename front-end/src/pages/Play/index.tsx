import React, { useState } from 'react'
import Container from '../../components/Container'
import ToggleButton from '../../components/ToggleButton'
import Colors from '../../tokens/Colors'
import Game from './Game'
import Train from './Train'

const Play = () => {
  const [isCompMode, setCompMode] = useState(true)

  return (
    <Container>
      <div
        style={{
          width: '100%',
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          marginTop: '40px',
          marginBottom: '20px',
        }}>
        <div>
          <div
            style={{
              width: '100%',
              marginBottom: '3px',
              fontFamily: 'League Spartan',
              fontSize: '12px',
              color: Colors.SectionHeading,
            }}>
            MODE
          </div>
          <ToggleButton
            index={isCompMode ? 0 : 1}
            options={['Game', 'Train']}
            onSelect={(i: number) => setCompMode(i === 0)}
          />
        </div>
      </div>
      {isCompMode ? <Game /> : <Train />}
    </Container>
  )
}

export default Play
