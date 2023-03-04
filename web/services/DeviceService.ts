import { IDeviceCard, IResponse } from '@/lib/types'
import { ICreateDeviceForm } from './../lib/types'
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
}

export default DeviceService
