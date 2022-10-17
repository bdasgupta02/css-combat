import React, { useState } from 'react'
import { useSpring, animated, config } from 'react-spring'
import Colors, { percentageToHex } from '../../tokens/Colors'

const Button = ({
  type = 'primary',
  text,
  onClick,
  outerStyle = {},
  isLoading = false,
}: {
  type?: 'primary' | 'secondary' | 'disabled'
  forceWidth?: string
  text: any
  onClick: any
  outerStyle?: any
  isLoading?: boolean
}) => {
  const [isHover, setHover] = useState(false)
  const [isPressed, setPressed] = useState(false)

  const bgSpring = useSpring({
    backgroundColor: isPressed ? Colors.Button : isHover ? Colors.ButtonHover : Colors.Button,
    config: config.stiff,
  })

  const shadowSpring = useSpring({
    boxShadow: isPressed
      ? `3px 10px 10px ${Colors.ButtonRoot}00`
      : isHover
      ? `8px 10px 20px ${Colors.ButtonRoot}${percentageToHex(40)}`
      : `3px 10px 10px ${Colors.ButtonRoot}${percentageToHex(25)}`,
  })

  return (
    <animated.div
      onMouseEnter={() => setHover(true)}
      onMouseLeave={() => {
        setPressed(false)
        setHover(false)
      }}
      onMouseDown={
        type === 'disabled'
          ? () => setPressed(true)
          : () => {
              setPressed(true)
              onClick()
            }
      }
      onMouseUp={() => {
        setPressed(false)
        setHover(false)
      }}
      style={{
        outline: `4px solid ${Colors.ButtonRoot}${percentageToHex(70)}`,
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '50px',
        width: '100%',
        fontSize: '18px',
        fontWeight: 'bold',
        fontFamily: 'Lato',
        cursor: 'pointer',
        color: Colors.White,
        ...bgSpring,
        ...shadowSpring,
        ...outerStyle,
      }}>
      {text}
    </animated.div>
  )
}

export default Button
