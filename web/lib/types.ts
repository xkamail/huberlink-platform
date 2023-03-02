export interface IResponse<T extends any> {
  success: boolean
  code: number
  message: string
  data: T
  errors: any[]
}
export enum ResponseCode {
  Success = 0,
  ResourceNotFound = 2,
  TokenExpired = 5,
  InvalidInput,
  InvalidToken,
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
