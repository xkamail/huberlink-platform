'use client'
import PageHeader from '@/components/ui/page-header'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { IDeviceDetail, ResponseCode } from '@/lib/types'
import DeviceService from '@/services/DeviceService'
import { useCallback, useEffect, useState } from 'react'

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
    return <p className="text-center">Loading...</p>

  return (
    <>
      <PageHeader title={data.name} />
      <div className="w-full bg-white p-4">xx</div>
    </>
  )
}

export default DeviceDetailPage

DeviceDetailPage.displayName = 'DeviceDetailPage'
