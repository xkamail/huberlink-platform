'use client'

import { Button } from '@/components/ui/button'
import { IHome } from '@/lib/types'
import HomeService from '@/services/HomeService'
import { parseCookies } from 'nookies'
import { useEffect, useState } from 'react'

const GoHomeButton = () => {
  const [homes, setHomes] = useState<IHome[]>([])
  useEffect(() => {
    HomeService.list().then((res) => {
      if (res.success) {
        setHomes(res.data)
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
  if (homes.length > 0) {
    return (
      <Button to={`/h/${homes[0].id}`} variant="primary">
        <span className="text-sm uppercase">Open Home</span>
      </Button>
    )
  }
  return (
    <Button to={`/h/create`}>
      <span className="text-sm uppercase">Create Home</span>
    </Button>
  )
}

export default GoHomeButton

GoHomeButton.displayName = 'GoHomeButton'
