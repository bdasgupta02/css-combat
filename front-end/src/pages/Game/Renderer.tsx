import React from 'react'

const Renderer = ({ inner }: { inner: string }) => {
  return (
    <div dangerouslySetInnerHTML={{ __html: inner }} />
  )
}

export default Renderer