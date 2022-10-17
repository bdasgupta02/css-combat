import React from 'react'
import styled from 'styled-components'
import Colors from '../../tokens/Colors'

export const Heading = ({ children, style }: { children: string, style?: any }) => {
  return (
    <TextRoot style={style}>
      {children}
    </TextRoot>
  )
}

const TextRoot = styled.div`
  color: ${Colors.White};
  font-family: Lato, sans-serif;
  font-size: 28px;
`