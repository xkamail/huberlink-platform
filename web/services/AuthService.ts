import { IResponse, ISignInForm } from '@/lib/types'
import { IUser } from './../lib/types'
import { fetcher } from './requests'

const AuthService = {
  signIn(form: ISignInForm) {
    return fetcher
      .post<
        IResponse<{
          token: string
          refreshToken: string
        }>
      >(`/auth/sign-in-username`, form)
      .then((r) => r.data)
  },
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
  invokeRefreshToken(refreshToken: string) {
    return fetcher
      .post<
        IResponse<{
          token: string
          refreshToken: string
        }>
      >('/auth/refresh-token?refreshToken=' + refreshToken)
      .then((r) => r.data)
  },
}

export default AuthService
