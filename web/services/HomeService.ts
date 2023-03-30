import { ICreateHomeForm, IResponse, IScene } from '@/lib/types'
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
  scenes: {
    list(homeId: string) {
      return fetcher
        .get<IResponse<IScene[]>>(`/home/${homeId}/scenes`)
        .then((r) => r.data)
    },
    findById(homeId: string, sceneId: string) {
      return fetcher
        .get<IResponse<IScene>>(`/home/${homeId}/scenes/${sceneId}`)
        .then((r) => r.data)
    },
    create({ homeId }: { homeId: string }) {
      return fetcher
        .post<IResponse<{ id: string }>>(`/home/${homeId}/scenes`, { homeId })
        .then((r) => r.data)
    },
    delete({ homeId, sceneId }: { homeId: string; sceneId: string }) {
      return fetcher
        .delete<
          IResponse<{
            id: string
          }>
        >(`/home/${homeId}/scenes/${sceneId}`)
        .then((r) => r.data)
    },
  },
}

export default HomeService
