import { ICreateHomeForm, IResponse } from '@/lib/types'
import { IHome } from './../lib/types'
import { fetcher } from './requests'

const HomeService = {
  create(data: ICreateHomeForm) {
    return fetcher
      .post<
        IResponse<{
          id: string
        }>
      >(`/home`, data)
      .then((r) => r.data)
  },
  list() {
    return fetcher.get<IResponse<IHome[]>>(`/home`).then((r) => r.data)
  },
  findById(id: string, accessToken?: string) {
    return fetcher
      .get<
        IResponse<{
          home: IHome
          member: any[]
        }>
      >(`/home/${id}`, {
        headers:
          (accessToken && {
            Authorization: `Bearer ${accessToken}`,
          }) ||
          {},
      })
      .then((r) => r.data)
  },
}

export default HomeService
