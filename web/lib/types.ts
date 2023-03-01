export interface IResponse<T extends any> {
  success: boolean
  code: number
  message: string
  data: T
  errors: any[]
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
