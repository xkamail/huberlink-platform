import { cn } from '@/lib/utils'
import React from 'react'

const Form = ({
  children,
  className,
}: {
  children: React.ReactNode
  className?: string
}) => {
  return <form className={cn(`space-y-4`, className)}>{children}</form>
}

export default Form

Form.displayName = 'Form'
