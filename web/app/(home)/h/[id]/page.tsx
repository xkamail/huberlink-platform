'use client'
import IRRemoteThingCard from '@/components/thing/ir-remote'
import { Button } from '@/components/ui/button'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { useUser } from '@/lib/hooks'
import { DeviceKindEnum, IDeviceCard } from '@/lib/types'
import { toSWR } from '@/lib/utils'
import DeviceService from '@/services/DeviceService'
import { PlusIcon } from 'lucide-react'
import { Metadata } from 'next'
import useSWR from 'swr'
import HomeSenceList from './home-sence'
import SkeletonDisplay from './skeleton'

export const metadata: Metadata = {
  title: 'Home',
}

const HomePage = ({ params: { id: homeId } }: { params: { id: string } }) => {
  const {
    data: devices,
    error,
    isLoading,
  } = useSWR(['home', homeId], toSWR(DeviceService.list(homeId)))

  const { profile } = useUser()
  const renderDeviceCard = (d: IDeviceCard) => {
    if (d.kind === DeviceKindEnum.IRRemote) {
      return <IRRemoteThingCard deviceId={d.id} />
    }
    return null
  }

  const hideScene = true
  const loading = useHomeSelector((s) => s.isLoading)

  if (loading) return <SkeletonDisplay />

  const deviceList = devices || []
  if (deviceList.length === 0) {
    return (
      <div>
        <div className="mb-4 py-20 flex justify-center items-center flex-col gap-2">
          <p className="mb-4">
            You {`don't`} have any devices yet. Click the button below to create
            a
          </p>
          <Button to={`/h/${homeId}/devices/create`} variant="outline-primary">
            <PlusIcon className="w-4 h-4 inline-block mr-1" /> Create a new
            device
          </Button>
        </div>
      </div>
    )
  }
  return (
    <div className="grid grid-cols-12 gap-4">
      {!hideScene && (
        <div className="col-span-12">
          <HomeSenceList />
        </div>
      )}

      {!error && !isLoading && deviceList.map((d) => renderDeviceCard(d))}
    </div>
  )
}
export default HomePage
