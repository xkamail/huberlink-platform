'use client'

import HomeService from '@/services/HomeService'
import { useRouter } from 'next/navigation'
import { parseCookies } from 'nookies'
import { useEffect } from 'react'

const PreHomePage = () => {
  const router = useRouter()
  const cookies = parseCookies(null)
  const findFirstHome = async () => {
    const res = await HomeService.list()
    if (res.success && res.data.length > 0) {
      router.push(`/h/${res.data[0].id}`)
    } else {
      router.push('/h/create')
    }
  }
  if (cookies.currentHome) {
    router.push(`/h/${cookies.currentHome}`)
  }
  useEffect(() => {
    if (!cookies.currentHome) {
      findFirstHome()
    }
  }, [cookies])
  return (
    <div className="mt-4">
      <div className="bg-slate-200 duration-1000 animate-pulse rounded-lg w-full h-16"></div>

      <div className="gap-4 grid mt-10 grid-cols-3">
        <div className="bg-slate-200 duration-1000 animate-pulse rounded-lg w-full h-16"></div>

        <div className="bg-slate-200 duration-1000 animate-pulse rounded-lg w-full h-16"></div>
        <div className="bg-slate-200 duration-1000 animate-pulse rounded-lg w-full h-16"></div>
        <div className="bg-slate-200 duration-1000 animate-pulse rounded-lg w-full h-16"></div>
        <div className="bg-slate-200 duration-1000 animate-pulse rounded-lg w-full h-16"></div>
        <div className="bg-slate-200 duration-1000 animate-pulse rounded-lg w-full h-16"></div>
      </div>
    </div>
  )
}

export default PreHomePage
