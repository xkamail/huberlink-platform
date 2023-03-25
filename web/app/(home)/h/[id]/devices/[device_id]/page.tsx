'use client'
import Card from '@/components/ui/card'
import PageHeader from '@/components/ui/page-header'
import Spinner from '@/components/ui/spinner'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import DeviceService from '@/services/DeviceService'
import { useState } from 'react'
import useSWR from 'swr'
import DeviceInformation from './information'
import IRRemoteSection from './ir-remote'
//
const DeviceDetailPage = ({
  params: { device_id: deviceId },
}: {
  params: { device_id: string }
}) => {
  const homeId = useHomeSelector((s) => s.homeId)
  const [status, setStatus] = useState<
    'idle' | 'loading' | 'ok' | 'error' | 'notfound'
  >('idle')
  const {
    data: resp,
    isLoading,
    error,
  } = useSWR(
    'device-detail',
    () =>
      DeviceService.findById({
        deviceId,
        homeId,
      }),
    {
      refreshInterval: 1000,
    }
  )

  if (error || !resp?.success)
    return (
      <div className="text-center">
        <p>{resp?.message}</p>
      </div>
    )

  if (isLoading || !resp)
    return (
      <div className="text-center w-full flex justify-center my-20">
        <Spinner />
      </div>
    )

  const data = resp.data

  return (
    <>
      <PageHeader title={data.name} />
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div className="col-span-1">
          <DeviceInformation data={data} />
        </div>
        <div className="col-span-1">
          <Card>-</Card>
        </div>
        <div className="col-span-2">
          <IRRemoteSection device={data} />
        </div>
      </div>
    </>
  )
}

export default DeviceDetailPage

DeviceDetailPage.displayName = 'DeviceDetailPage'
