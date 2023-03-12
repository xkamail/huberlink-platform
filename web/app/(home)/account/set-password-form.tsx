'use client'
import { Button } from '@/components/ui/button'
import Form from '@/components/ui/form'
import FormInput from '@/components/ui/form-input'
import { useToast } from '@/hooks/use-toast'
import AuthService from '@/services/AuthService'
import { useForm } from 'react-hook-form'

const SetPasswordForm = () => {
  const { toast } = useToast()
  const ctx = useForm({
    defaultValues: {
      password: '',
      confirmPassword: '',
    },
  })

  const submit = async (data: {
    password: string
    confirmPassword: string
  }) => {
    if (data.password !== data.confirmPassword) {
      return
    }
    const res = await AuthService.setPassword({
      password: data.password,
    })
    if (!res.success) {
      toast.error(res.message)
      return
    }
    toast.succes(`Password set successfully!`)
  }

  return (
    <Form ctx={ctx} onSubmit={submit}>
      <FormInput
        label="New Password"
        name="password"
        placeholder=""
        type="password"
        options={{
          minLength: {
            value: 8,
            message: 'Password must be at least 8 characters long',
          },
          required: 'Password is required',
        }}
      />
      <FormInput
        label="Confirm New Password"
        name="confirmPassword"
        placeholder=""
        type="password"
        options={{
          minLength: {
            value: 8,
            message: 'Confirm password must be at least 8 characters long',
          },
          validate: (value) => {
            if (value === ctx.watch('password')) return true
            return 'Password do not match'
          },
          required: 'Confirm password is required',
        }}
      />
      <div className="flex justiy-center md:justify-end">
        <Button type="submit" className="ml-auto">
          <span>Set Password</span>
        </Button>
      </div>
    </Form>
  )
}

export default SetPasswordForm

SetPasswordForm.displayName = 'SetPasswordForm'
