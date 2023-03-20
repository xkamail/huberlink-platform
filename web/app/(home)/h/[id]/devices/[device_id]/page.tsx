'use client'
import PageHeader from '@/components/ui/page-header'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import DeviceService from '@/services/DeviceService'
import { useCallback, useEffect, useState } from 'react'

//
const DeviceDetailPage = ({
  params: { device_id: deviceId },
}: {
  params: { device_id: string }
}) => {
  const homeId = useHomeSelector((s) => s.homeId)
  const [data, setData] = useState<any>(null)
  const [status, setStatus] = useState<
    'idle' | 'loading' | 'ok' | 'error' | 'notfound'
  >('idle')
  const fetchData = useCallback(() => {
    DeviceService.findById({ deviceId, homeId })
  }, [deviceId, homeId])

  useEffect(() => {
    fetchData()
  }, [])

  if (status === 'notfound')
    return (
      <>
        <p>Device not found</p>
      </>
    )
  return (
    <>
      <PageHeader title="xxxx" />
      <div className="bg-slate-100 w-full">xx</div>
    </>
  )
}

export default DeviceDetailPage

DeviceDetailPage.displayName = 'DeviceDetailPage'
