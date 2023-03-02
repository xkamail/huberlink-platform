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
      >(`/homes`, data)
      .then((r) => r.data)
  },
  list() {
    return fetcher.get<IResponse<IHome[]>>(`/home`).then((r) => r.data)
  },
}

export default HomeService
