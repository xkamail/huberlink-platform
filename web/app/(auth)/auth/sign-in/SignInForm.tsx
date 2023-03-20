'use client'

import { Button } from '@/components/ui/button'
import Form from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { useToast } from '@/hooks/use-toast'
import { ISignInForm } from '@/lib/types'
import { formError } from '@/lib/utils'
import AuthService from '@/services/AuthService'
import { useRouter } from 'next/navigation'
import nookies from 'nookies'
import { useForm } from 'react-hook-form'
const SignInForm = () => {
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
    nookies.set(null, 'accessToken', res.data.token)
    nookies.set(null, 'refreshToken', res.data.refreshToken)
    //
    toast.succes('Signed in successfully')

    router.push('/h')
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
