'use client'
import PageHeader from '@/components/ui/page-header'
import Spinner from '@/components/ui/spinner'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { IDeviceDetail, ResponseCode } from '@/lib/types'
import DeviceService from '@/services/DeviceService'
import { useCallback, useEffect, useState } from 'react'
import DeviceInformation from './information'
import IRRemoteSection from './ir-remote'

//
const DeviceDetailPage = ({
  params: { device_id: deviceId },
}: {
  params: { device_id: string }
}) => {
  const homeId = useHomeSelector((s) => s.homeId)
  const [data, setData] = useState<IDeviceDetail | null>(null)
  const [status, setStatus] = useState<
    'idle' | 'loading' | 'ok' | 'error' | 'notfound'
  >('idle')

  const fetchData = useCallback(async () => {
    const res = await DeviceService.findById({ deviceId, homeId })
    if (!res.success) {
      if (res.code === ResponseCode.ResourceNotFound) {
        setStatus('notfound')
      }
      setStatus('error')
      return
    }
    setData(res.data)
    const resIR = await DeviceService.ir.findDetail({
      homeId,
      deviceId,
    })
    if (!resIR.success) {
      setStatus('error')
      return
    }
    setStatus('ok')
  }, [deviceId, homeId])

  useEffect(() => {
    fetchData()
  }, [])

  if (status === 'notfound')
    return (
      <div className="text-center">
        <p>Device not found</p>
      </div>
    )
  if (status === 'loading' || !data)
    return (
      <div className="text-center w-full flex justify-center my-20">
        <Spinner />
      </div>
    )

  return (
    <>
      <PageHeader title={data.name} />
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div className="col-span-1">
          <DeviceInformation data={data} />
        </div>
        <div className="col-span-2">
          <IRRemoteSection />
        </div>
      </div>
    </>
  )
}

export default DeviceDetailPage

DeviceDetailPage.displayName = 'DeviceDetailPage'
