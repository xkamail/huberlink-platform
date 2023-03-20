export type IResponse<T> =
  | {
      success: true
      code: ResponseCode
      message: string
      data: T
    }
  | {
      success: false
      code: ResponseCode
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
export interface IDeviceDetail {
  id: string
  name: string
  icon: string
  model: string
  kind: DeviceKindEnum
  homeId: string
  userId: string
  token: string
  ipAddress: string | null
  location: string | null
  latestHeartbeatAt: Date | null
  createdAt: Date
  updatedAt: Date
}

export interface IIRRemote {
  id: string
  deviceId: string
  homeId: string
  createdAt: Date
  updatedAt: Date
}

export interface IIRRemoteVirtualDevice {
  id: string
  remoteId: string
  name: string
  kind: VirtualCategoryEnum
  icon: string
  isLearning: boolean
  properties: { [key: string]: any }
  createdAt: Date
  updatedAt: Date
}

export enum VirtualCategoryEnum {
  Other,
  TV,
  AirConditioner,
  Light,
  Fan,
  Speaker,
  Projector,
  DVD,
  WaterHeart,
}
