import {
  ICreateDeviceForm,
  IDetailDetail,
  IDeviceCard,
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
      .get<IResponse<IDetailDetail>>(`/home/${homeId}/devices/${deviceId}`)
      .then((r) => r.data)
  },

  ir: {
    //
    findDetail({
      homeId,
      deviceId,
      remoteId,
    }: {
      homeId: string
      deviceId: string
      remoteId: string
    }) {
      return fetcher
        .get<
          IResponse<{
            vs: IIRRemoteVirtualDevice[]
            remote: IIRRemote
          }>
        >(`/home/${homeId}/devices/${deviceId}/ir-remote/${remoteId}`)
        .then((r) => r.data)
    },
  },
}

export default DeviceService
