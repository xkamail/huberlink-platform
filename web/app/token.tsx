'use client'

import { Button } from '@/components/ui/button'
import nookies from 'nookies'

const TokenTest = () => {
  const setToken = () => {
    nookies.set(null, 'xxx', 'yyy')
  }
  return (
    <div className="flex w-full">
      {JSON.stringify(nookies.get(null, 'xxx').xxx)}
      <Button
        onClick={() => {
          setToken()
        }}
      >
        press to set a token
      </Button>
    </div>
  )
}

export default TokenTest

TokenTest.displayName = 'TokenTest'
