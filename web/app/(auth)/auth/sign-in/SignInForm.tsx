'use client'

import { Button } from '@/components/ui/button'
import Form from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { useToast } from '@/hooks/use-toast'
import { useUserDispatch } from '@/lib/contexts/UserContext'
import { ISignInForm } from '@/lib/types'
import { formError, setAuthCookie } from '@/lib/utils'
import AuthService from '@/services/AuthService'
import { useRouter } from 'next/navigation'
import { parseCookies } from 'nookies'
import { useForm } from 'react-hook-form'
const SignInForm = () => {
  // eslint-disable-next-line react-hooks/rules-of-hooks
  const userDispatch = useUserDispatch()
  const { toast } = useToast()
  const router = useRouter()
  const ctx = useForm({
    defaultValues: {
      username: '',
      password: '',
    },
  })
  const submit = async (data: ISignInForm) => {
    const res = await AuthService.signIn(data)
    if (!res.success) {
      toast.error(res.message)
      formError(ctx, res)
      return
    }
    setAuthCookie(res)
    toast({
      title: 'Sign in success',
    })
    // pre fetch user when login success
    userDispatch({ type: 'fetch-user' }).then(() => {
      const cookies = parseCookies(null)
      if (cookies.currentHome) {
        router.push(`/h/${cookies.currentHome}`)
      } else {
        router.push(`/h`)
      }
    })
  }

  return (
    <Form ctx={ctx} onSubmit={submit} className=" mb-5 mx-auto">
      <Input
        {...ctx.register('username')}
        placeholder="Username or E-mail"
        type="text"
      />
      <Input
        {...ctx.register('password')}
        name="password"
        placeholder="Password"
        type="password"
      />
      <Button
        loading={ctx.formState.isSubmitting}
        block
        type="submit"
        variant="default"
      >
        Sign in
      </Button>
    </Form>
  )
}

export default SignInForm

SignInForm.displayName = 'SignInForm'
