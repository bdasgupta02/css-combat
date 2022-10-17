import React, { useState } from 'react'
import { animated, config, useSpring } from 'react-spring'
import styled from 'styled-components'
import Colors, { percentageToHex } from '../../tokens/Colors'

// animations
const Input = ({
  value,
  onChange,
  placeholder,
  outerStyle,
  innerProps,
}: {
  value: string
  onChange: any
  placeholder: string
  outerStyle?: any
  innerProps?: any
}) => {
  const [isHover, setHover] = useState(false)
  const [isFocused, setFocused] = useState(false)

  const bgSpring = useSpring({
    backgroundColor: isFocused ? Colors.InputFocused : Colors.Input,
  })

  const textSpring = useSpring({
    color: isFocused ? Colors.InputTextFocused : Colors.InputText,
    config: config.stiff,
  })

  const borderSpring = useSpring({
    outline: isFocused
      ? `4px solid ${Colors.InputBorderFocused}`
      : isHover
      ? `4px solid ${Colors.InputBorderHover}`
      : `4px solid ${Colors.Transparent}`,
  })

  const AnimatedOuter = animated(Outer)
  const AnimatedInner = animated(Inner)

  return (
    <AnimatedOuter
      onMouseOver={() => setHover(true)}
      onMouseLeave={() => setHover(false)}
      style={{
        ...bgSpring,
        ...outerStyle,
        ...borderSpring,
      }}>
      <AnimatedInner
        value={value}
        onChange={(e: any) => onChange(e)}
        placeholder={placeholder}
        onFocus={() => setFocused(true)}
        onBlur={() => setFocused(false)}
        style={{
          ...textSpring,
        }}
        {...innerProps}
      />
    </AnimatedOuter>
  )
}

const Outer = styled.div`
  width: 100%;
  padding: 0px 12px 0px;
  position: relative;
`

const Inner = styled.input`
  width: 100%;
  height: 48px;
  font-size: 16px;
  outline: none;
  border: none;
  background-color: #00000000;
  font-family: Lato;

  :focus::placeholder {
    color: ${Colors.White + percentageToHex(70)};
  }
`

export default Input
