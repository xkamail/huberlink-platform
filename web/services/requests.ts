import axios from 'axios'
import nookies from 'nookies'
const apiURL = process.env.NEXT_PUBLIC_API_URL

export const fetcher = axios.create({
  baseURL: apiURL,
  timeout: 10 * 1000,
})

fetcher.interceptors.request.use((config) => {
  const accessToken = nookies.get(null, 'accessToken').accessToken
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
