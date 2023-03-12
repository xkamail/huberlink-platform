'use client'
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
        options={{
          validate: {
            minLength: (value) =>
              value.length >= 8 ||
              'Password must be at least 8 characters long',
          },
        }}
      />
      <FormInput
        label="Confirm New Password"
        name="confirmPassword"
        placeholder=""
        options={{
          validate: {
            minLength: (value) =>
              value.length >= 8 ||
              'Confirm password must be at least 8 characters long',
            validate: (value) => {
              if (value === ctx.getValues('confirmPassword')) {
                return true
              }
              return 'Password do not match'
            },
          },
        }}
      />
    </Form>
  )
}

export default SetPasswordForm

SetPasswordForm.displayName = 'SetPasswordForm'
