import { IResponse } from '@/lib/types'
import axios, { AxiosResponse } from 'axios'
import nookies, { destroyCookie, parseCookies, setCookie } from 'nookies'
import { ResponseCode } from './../lib/types'
import AuthService from './AuthService'
const apiURL = process.env.NEXT_PUBLIC_API_URL

export const fetcher = axios.create({
  baseURL: apiURL,
  timeout: 10 * 1000,
})

fetcher.interceptors.request.use((config) => {
  const accessToken = nookies.get(null, 'accessToken')?.accessToken
  if (accessToken && config.headers.Authorization === undefined) {
    config.headers.Authorization = `Bearer ${accessToken}`
  }
  return config
})

fetcher.interceptors.response.use(
  async (f: AxiosResponse<IResponse<any>>) => {
    if (f.status !== 200) {
      return {
        ...f,
        data: {
          success: false,
          code: ResponseCode.ClientError,
          errors: [],
          message: `Unexpected status code: ${f.status}`,
        },
      }
    }
    const originalRequest = f.config
    if (
      !f.data.success &&
      f.data.code === ResponseCode.TokenExpired &&
      // @ts-ignore
      !originalRequest._retry
    ) {
      const cookie = parseCookies(null)
      if (cookie['refreshToken']) {
        // @ts-ignore
        originalRequest._retry = true
        const refreshRes = await AuthService.invokeRefreshToken(
          cookie.refreshToken
        )
        if (refreshRes.success) {
          console.log('[Interceptor] Refresh token success')

          destroyCookie(null, 'accessToken')
          destroyCookie(null, 'refreshToken')
          setCookie(null, 'accessToken', refreshRes.data.token)
          setCookie(null, 'refreshToken', refreshRes.data.refreshToken)
          originalRequest.headers.Authorization = `Bearer ${refreshRes.data.token}`
          return fetcher(originalRequest)
        }
      }
      return {
        ...f,
        data: {
          success: false,
          code: ResponseCode.TokenExpired,
          errors: [],
          message: 'Token expired',
        },
      }
    }
    return {
      ...f,
      data: f.data,
    }
  },
  (error) => {
    return Promise.reject(error)
  }
)
