import { cn } from '@/lib/utils'
import React from 'react'
import { FormProvider, UseFormReturn } from 'react-hook-form'

const Form = ({
  children,
  className,
  onSubmit,
  ctx,
}: {
  children: React.ReactNode
  className?: string
  onSubmit: (data: any) => void
  ctx: UseFormReturn<any>
}) => {
  return (
    <FormProvider {...ctx}>
      <form
        onSubmit={ctx.handleSubmit(onSubmit)}
        className={cn(`space-y-4`, className)}
      >
        {children}
      </form>
    </FormProvider>
  )
}

export default Form

Form.displayName = 'Form'
