import { IResponse, ResponseCode } from '@/lib/types'
import { NextRequest } from 'next/server'

const apiURL = String(process.env.NEXT_PUBLIC_API_URL)

export const fetchy = {
  get: async <T extends any>(
    req: NextRequest,
    url: string,
    options?: any
  ): Promise<IResponse<T>> => {
    const token = req.cookies.get('accessToken')?.value
    console.log('[INFO] Fetching', `${apiURL}${url}`)

    const res = (await fetch(`${apiURL}${url}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
      method: 'GET',
    })
      .then(async (r) => {
        if (r.status !== 200) {
          throw new Error('Failed to fetch')
        }

        const res = (await r.json()) as IResponse<any>
        console.log(res)

        if (res.code === ResponseCode.TokenExpired) {
          const result = await doRefreshToken(req)
          if (!result) {
            return {
              success: false,
              code: ResponseCode.InvalidInput,
              message: 'invalid token',
              data: null,
            }
          }
          req.cookies.set('accessToken', result.token)
          req.cookies.set('refreshToken', result.refreshToken)
          // retry request
          return fetchy.get(req, url, options)
        }

        return res
      })
      .catch((e) => {
        console.log(e)
        return {
          success: false,
        }
      })) as IResponse<T>
    return res
  },
  post: async (req: NextRequest, url: string, data: any, options?: any) => {},
}

export const doRefreshToken = async (
  req: NextRequest
): Promise<{
  token: string
  refreshToken: string
} | null> => {
  console.log('[INFO] doRefreshToken')

  const refreshToken = req.cookies.get('refreshToken')?.value
  if (!refreshToken) {
    return null
  }
  const res = (await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}/auth/refresh-token?refreshToken=${refreshToken}`,
    {
      method: 'POST',
    }
  ).then((r) => r.json())) as IResponse<{
    token: string
    refreshToken: string
  }>
  if (res.success) {
    return res.data
  }
  return null
}
