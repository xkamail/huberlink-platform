export interface IResponse<T extends any> {
  success: boolean
  code: number
  message: string
  data: T
  errors: any[]
}
export enum ResponseCode {
  Success = 0,
  InvalidInput = 6,
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
