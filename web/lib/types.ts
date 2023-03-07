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

export type IDeviceCard = {
  id: string
  name: string
  icon: string
  kind: DeviceKindEnum
  latestHeartbeatAt: null | Date
}

export enum DeviceKindEnum {
  Unknown,
  IRRemote,
  Sensor,
  Switch,
  Lamp,
}

export const DEVICE_CATEGORY = [
  {
    name: 'Unknown',
    description: '',
    icon: 'unknown',
    kind: DeviceKindEnum.Unknown,
  },
  {
    name: 'IR Remote',
    description: 'Universal Remote Control',
    icon: 'remote',
    kind: DeviceKindEnum.IRRemote,
  },
  {
    name: 'Sensor',
    description: 'Generic Sensor',
    icon: 'sensor',
    kind: DeviceKindEnum.Sensor,
  },
  {
    name: 'Switch',
    description: 'Generic Switch',
    icon: 'switch',
    kind: DeviceKindEnum.Switch,
  },
]
export type ICreateDeviceForm = {
  name: string
  kind: DeviceKindEnum
  icon?: string
  model?: string
}
