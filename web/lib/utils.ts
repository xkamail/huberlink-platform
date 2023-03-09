import { IResponse } from '@/lib/types'
import { ClassValue, clsx } from 'clsx'
import { setCookie } from 'nookies'
import { UseFormReturn } from 'react-hook-form'
import { twMerge } from 'tailwind-merge'
import { ResponseCode } from './types'
export function toSWR<T extends any>(f: Promise<IResponse<T>>) {
  return () =>
    f.then((res) => {
      if (res.success) return res.data
      throw res
    })
}
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

export function setAuthCookie(
  res: IResponse<{
    token: string
    refreshToken: string
  }>
) {
  if (!res.success) return
  const maxAge = 60 * 60 * 24 * 7
  const opts = {
    maxAge,
    expires: new Date(Date.now() + maxAge * 1000),
  }
  setCookie(null, 'accessToken', res.data.token, opts)
  setCookie(null, 'refreshToken', res.data.refreshToken, opts)
}
