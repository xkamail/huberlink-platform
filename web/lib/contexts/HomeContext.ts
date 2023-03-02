import { ResponseCode } from '@/lib/types'
import HomeService from '@/services/HomeService'
import { useCallback, useEffect, useState } from 'react'
import { createProvider } from '.'
import { IHome } from './../types'

type IHomeAction = { type: 'value'; payload: string }

export const [HomeContextProvider, useHomeDispatch, useHomeSelector] =
  createProvider(({ homeId }: { homeId: string }) => {
    const [home, setHome] = useState<IHome | null>(null)
    const [status, setStatus] = useState<
      'idle' | 'loading' | 'done' | 'error' | 'not-found'
    >('idle')
    const fetchDetail = useCallback(async () => {
      setStatus('loading')
      const res = await HomeService.findById(homeId)
      if (res.success) {
        setHome(res.data.home)
        setStatus('done')
        return
      }
      if (res.code === ResponseCode.ResourceNotFound) {
        setStatus('not-found')
        return
      }
      setStatus('error')
    }, [homeId])
    useEffect(() => {
      fetchDetail()
    }, [fetchDetail])
    const dispatch = (action: IHomeAction) => {
      switch (action.type) {
        case 'value':
          break

        default:
          break
      }
    }
    const state = {
      homeId: homeId,
      homeName: home?.name || '',
      isNotFound: status === 'not-found',
      isIdle: status === 'idle',
      isError: status === 'error',
      isLoading: status === 'loading',
    }
    return [state, dispatch]
  })
