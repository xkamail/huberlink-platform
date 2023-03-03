export type IResponse<T> =
  | {
      success: true
      code: number
      message: string
      data: T
    }
  | {
      success: false
      code: number
      message: string
      errors: any[]
    }
export enum ResponseCode {
  Success = 0,
  ResourceNotFound = 2,
  TokenExpired = 5,
  InvalidInput,
  InvalidToken,

  ClientError = 777,
}

export type IUser = {
  id: string
  username: string
  email: string
  avatar: string
  discriminator: string
  createdAt: Date
  updatedAt: Date
}

export type ISignInForm = {
  username: string
  password: string
}

export type ICreateHomeForm = {
  name: string
}

export type IHome = {
  id: string
  name: string
  userId: string
}
