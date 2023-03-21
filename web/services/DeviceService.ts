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
            virtuals: IIRRemoteVirtualDevice[]
            remote: IIRRemote
          }>
        >(`/home/${homeId}/devices/${deviceId}/ir-remote`)
        .then((r) => r.data)
    },
    createVirtual({
      homeId,
      deviceId,
      name,
      kind,
      icon,
    }: {
      homeId: string
      deviceId: string
      name: string
      kind: string
      icon: string
    }) {
      return fetcher
        .post<
          IResponse<{
            id: string
          }>
        >(`/home/${homeId}/devices/${deviceId}/ir-remote/virtual`, {
          name,
          kind,
          icon,
        })
        .then((r) => r.data)
    },
    listVirtual({ homeId, deviceId }: { homeId: string; deviceId: string }) {
      return fetcher.get<IResponse<{}>>(
        `/home/${homeId}/devices/${deviceId}/ir-remote/ir-remote`
      )
    },
    deleteVirtual({
      homeId,
      deviceId,
      virtualId,
    }: {
      homeId: string
      deviceId: string
      virtualId: string
    }) {
      return fetcher.delete<IResponse<{}>>(
        `/home/${homeId}/devices/${deviceId}/ir-remote/virtual/${virtualId}`
      )
    },
  },
}

export default DeviceService
