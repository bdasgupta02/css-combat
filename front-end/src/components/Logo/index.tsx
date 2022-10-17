import React from 'react'
import Colors from '../../tokens/Colors'

export const Logo = () => {
  return (
    <div
      style={{
        fontFamily: 'League Gothic',
        fontSize: '45px',
        color: Colors.Logo,
      }}>
      <span style={{ color: Colors.White, fontSize: '60px' }}>{'{'}</span>CSSCombat
      <span style={{ color: Colors.White, fontSize: '60px' }}>{'}'}</span>
    </div>
  )
}
