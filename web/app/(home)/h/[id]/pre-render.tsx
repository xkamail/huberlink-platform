'use client'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { setCookie } from 'nookies'
import React, { useEffect } from 'react'

const PreRenderHome = ({ children }: { children: React.ReactNode }) => {
  const isNotFound = useHomeSelector((s) => s.isNotFound)
  const isError = useHomeSelector((s) => s.isError)
  const homeId = useHomeSelector((s) => s.homeId)

  useEffect(() => {
    setCookie(null, `currentHome`, homeId, {
      maxAge: 30 * 24 * 60 * 60,
    })
  }, [])

  if (isNotFound)
    return (
      <div>
        <p>Home not found</p>
      </div>
    )
  if (isError)
    return (
      <div>
        <p>Error</p>
      </div>
    )
  return <>{children}</>
}

export default PreRenderHome

PreRenderHome.displayName = 'PreRenderHome'
