import { IResponse } from '@/lib/types'
import { IUser } from './../lib/types'
import { fetcher } from './requests'

const AuthService = {
  signInWithDiscord(code: string) {
    return fetcher
      .post<
        IResponse<{
          token: string
          refreshToken: string
        }>
      >('/auth/sign-in', {
        code,
      })
      .then((r) => r.data)
  },
  me(accessToken?: string) {
    return fetcher
      .get<IResponse<IUser>>(
        '/auth/me',
        accessToken
          ? {
              headers: {
                Authorization: `Bearer ${accessToken}`,
              },
            }
          : {}
      )
      .then((r) => r.data)
  },
}

export default AuthService
