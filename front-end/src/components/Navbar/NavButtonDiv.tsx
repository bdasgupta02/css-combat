import React from 'react'
import styled from 'styled-components'

const ButtonRoot = styled.div`
  text-decoration: none;
  font-size: 16px;
  color: white;
  margin-right: 16px;
  padding: 8px 20px 8px;
  outline: 1px solid #ffffff40;
  border: 3px solid #f53f4f00;
  background-color: #ffffff00;
  box-shadow: 0px 0px 0px #97979700;
  cursor: pointer;

  -webkit-transition: all 200ms ease-out;
  -moz-transition: all 200ms ease-out;
  -o-transition: all 200ms ease-out;
  -ms-transition: all 200ms ease-out;
  transition: all 200ms ease-out;

  -webkit-touch-callout: none;
  -webkit-user-select: none;
  -khtml-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;

  &:hover {
    background-color: #ffffff40;
  }
`

const NavButtonDiv = ({ text, onClick }: { text: string; onClick: any }) => {
  return (
    <ButtonRoot onClick={() => onClick()}>
      {text}
    </ButtonRoot>
  )
}

export default NavButtonDiv
