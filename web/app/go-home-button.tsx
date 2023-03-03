'use client'

import { Button } from '@/components/ui/button'
import { IHome } from '@/lib/types'
import HomeService from '@/services/HomeService'
import { parseCookies } from 'nookies'
import { useEffect, useState } from 'react'

const GoHomeButton = () => {
  const [homes, setHomes] = useState<IHome[]>([])
  const [login, setLogin] = useState(false)
  useEffect(() => {
    HomeService.list().then((res) => {
      if (res.success) {
        setLogin(true)
        setHomes(res.data)
      } else {
        setLogin(false)
      }
    })
  }, [])

  const cookie = parseCookies(null)
  if (cookie['currentHome']) {
    return (
      <Button to={`/h/${cookie['currentHome']}`} variant="primary">
        <span className="text-sm uppercase">Open Home</span>
      </Button>
    )
  }
  if (homes.length > 0 && login) {
    return (
      <Button to={`/h/${homes[0].id}`} variant="primary">
        <span className="text-sm uppercase">Open Home</span>
      </Button>
    )
  }
  if (!login)
    return (
      <Button to={`/auth/sign-in?redirect=/h`}>
        <span className="text-sm uppercase">Sign in</span>
      </Button>
    )
  return (
    <Button to={`/h/create`}>
      <span className="text-sm uppercase">Create Home</span>
    </Button>
  )
}

export default GoHomeButton

GoHomeButton.displayName = 'GoHomeButton'
