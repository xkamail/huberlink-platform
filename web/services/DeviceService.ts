import {
  ICreateDeviceForm,
  IDeviceCard,
  IDeviceDetail,
  IIRRemote,
  IResponse,
} from '@/lib/types'
import { IIRRemoteVirtualDevice } from './../lib/types'
import { fetcher } from './requests'

const DeviceService = {
  list(homeId: string) {
    return fetcher
      .get<IResponse<IDeviceCard[]>>(`/home/${homeId}/devices/all`)
      .then((r) => r.data)
  },
  create(homeId: string, form: ICreateDeviceForm) {
    return fetcher
      .post<
        IResponse<{
          id: string
        }>
      >(`/home/${homeId}/devices`, form)
      .then((r) => r.data)
  },

  findById({ homeId, deviceId }: { homeId: string; deviceId: string }) {
    return fetcher
      .get<IResponse<IDeviceDetail>>(`/home/${homeId}/devices/${deviceId}`)
      .then((r) => r.data)
  },

  ir: {
    //
    findDetail({ homeId, deviceId }: { homeId: string; deviceId: string }) {
      return fetcher
        .get<
          IResponse<{
            vs: IIRRemoteVirtualDevice[]
            remote: IIRRemote
          }>
        >(`/home/${homeId}/devices/${deviceId}/ir-remote`)
        .then((r) => r.data)
    },
  },
}

export default DeviceService
