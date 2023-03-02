import axios from 'axios'
import { cookies } from 'next/headers'
import nookies from 'nookies'
const apiURL = process.env.NEXT_PUBLIC_API_URL

export const fetcher = axios.create({
  baseURL: apiURL,
  timeout: 10 * 1000,
})

fetcher.interceptors.request.use((config) => {
  const nextCookie = cookies()

  const accessToken =
    nookies.get(null, 'accessToken').accessToken ||
    nextCookie.get('accessToken')?.value

  if (accessToken && config.headers.Authorization === undefined) {
    config.headers.Authorization = `Bearer ${accessToken}`
  }
  return config
})

fetcher.interceptors.response.use(
  (f) => {
    return f
  },
  (error) => {
    return Promise.reject(error)
  }
)
