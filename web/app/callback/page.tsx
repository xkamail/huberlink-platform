'use client'

import { Button } from '@/components/ui/button'
import { useToast } from '@/hooks/use-toast'
import { useUserDispatch } from '@/lib/contexts/UserContext'
import { setAuthCookie } from '@/lib/utils'
import AuthService from '@/services/AuthService'
import { redirect, useRouter, useSearchParams } from 'next/navigation'
import { parseCookies } from 'nookies'
import { useCallback, useEffect, useState } from 'react'
// sign in with discord
// callback page
const CallbackPage = () => {
  const { toast } = useToast()
  const params = useSearchParams()
  const code = params.get('code')
  const [message, setMessage] = useState('')
  const router = useRouter()
  const userDispatch = useUserDispatch()
  const fetchData = useCallback(async () => {
    if (!code) return
    const res = await AuthService.signInWithDiscord(code)
    if (res.success) {
      setAuthCookie(res)
      toast({
        title: 'Sign in success',
      })
      // pre fetch user when login success
      userDispatch({ type: 'fetch-user' }).then(() => {
        setMessage(`Welcome back!`)
        const cookies = parseCookies(null)
        if (cookies.currentHome) {
          router.push(`/h/${cookies.currentHome}`)
        } else {
          router.push(`/h`)
        }
      })
    } else {
      setMessage(res.message)
      toast({
        variant: 'destructive',
        title: 'Sign in failed',
        description: res.message,
      })
    }
  }, [code, toast])

  useEffect(() => {
    fetchData()
  }, [fetchData])

  if (!code) {
    redirect('/auth/sign-in')
  }

  return (
    <div className=" h-screen container mx-auto my-10">
      <div className="bg-slate-100 rounded-lg space-y-10 text-center mx-auto max-w-xl pt-4 pb-10">
        <h1 className="text-2xl font-bold">Please wait</h1>
        {message === '' ? <p>Signing ....</p> : <p>{message}</p>}
        {message !== '' && <Button to="/">Go to home</Button>}
      </div>
    </div>
  )
}

export default CallbackPage

CallbackPage.displayName = 'CallbackPage'
