'use client'
const SetToken = () => {
  localStorage.setItem('accessToken', 'something store forever')
  return <></>
}

export default SetToken

SetToken.displayName = 'SetToken'
