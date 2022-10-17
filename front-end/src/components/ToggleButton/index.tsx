import React, { useState } from 'react'
import { animated, config, useSpring } from 'react-spring'
import Colors, { percentageToHex } from '../../tokens/Colors'

const ToggleElement = ({
  e,
  i,
  index,
  onSelect,
}: {
  e: string
  i: number
  index: number
  onSelect: any
}) => {
  const [isHover, setHover] = useState(false)

  const root = {
    padding: '10px 20px 10px',
    outline: '1px solid #ffffff40',
    cursor: 'pointer',
    backgroundColor: `${Colors.Blue}${percentageToHex(0)}`,
  }

  const inactive = {
    ...root,
  }

  const active = {
    ...root,
    backgroundColor: `${Colors.Blue}${percentageToHex(93)}`,
    outline: `1px solid ${Colors.Blue}${percentageToHex(93)}`,
  }

  const hover = {
    backgroundColor: `${Colors.White}40`,
  }

  const combined = i === index ? active : isHover ? hover : inactive

  const animSpring = useSpring({ ...combined, config: config.stiff })
  return (
    <animated.div
      onMouseEnter={() => setHover(true)}
      onMouseLeave={() => setHover(false)}
      onClick={() => onSelect(i)}
      style={animSpring}>
      {e}
    </animated.div>
  )
}

const ToggleButton = ({
  options,
  index,
  onSelect,
}: {
  options: string[]
  index: number
  onSelect: any
}) => {
  return (
    <div style={{ display: 'flex', flexDirection: 'row' }}>
      {options.map((e, i) => {
        return <ToggleElement e={e} i={i} index={index} onSelect={onSelect} />
      })}
    </div>
  )
}

export default ToggleButton
