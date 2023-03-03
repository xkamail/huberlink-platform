import { IResponse } from '@/lib/types'
import { ClassValue, clsx } from 'clsx'
import { UseFormReturn } from 'react-hook-form'
import { twMerge } from 'tailwind-merge'
import { ResponseCode } from './types'

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export function formError(ctx: UseFormReturn<any>, res: IResponse<any>) {
  console.log(res)

  if (
    !res.success &&
    res.code === ResponseCode.InvalidInput &&
    res.errors.length > 0
  ) {
    res.errors.map((err) => {
      if (!err.fieldName) return
      ctx.setError(err.fieldName!, {
        type: 'manual',
        message: err.reason!,
      })
    })
  }
}
