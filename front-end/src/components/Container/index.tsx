import React from 'react'
import styled from 'styled-components'
import Colors from '../../tokens/Colors'

const Bg = styled.div`
  width: 100%;
  overflow: auto;
`

const Root = styled.div`
  width: 100%;
  margin-left: auto;
  margin-right: auto;
`

type Props = {
  children: React.ReactNode
  breakpoint?: 'sm' | 'md' | 'lg' | string
  style?: any
  rootStyle?: any
  onScroll?: any
}

const Container = ({ children, breakpoint = 'md', style, rootStyle, onScroll }: Props) => {
  const getBreakPoint = () => {
    switch (breakpoint) {
      case 'sm':
        return '600px'
      case 'md':
        return '1000px'
      case 'lg':
        return '1200px'
      default:
        return breakpoint
    }
  }

  return (
    <Bg style={{ ...style }} onScroll={(e: any) => onScroll(e)}>
      <Root
        style={{
          maxWidth: getBreakPoint(),
          color: Colors.White,
          ...rootStyle,
        }}>
        <div style={{ padding: '0px 20px 0px', marginBottom: '50px' }}>{children}</div>
      </Root>
    </Bg>
  )
}

export default Container
