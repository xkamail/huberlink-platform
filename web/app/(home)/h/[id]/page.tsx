'use client'
import IRRemoteThingCard from '@/components/thing/ir-remote'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { useUser } from '@/lib/hooks'
import { DeviceKindEnum, IDeviceCard } from '@/lib/types'
import { toSWR } from '@/lib/utils'
import DeviceService from '@/services/DeviceService'
import HomeService from '@/services/HomeService'
import useSWR from 'swr'
import HomeSenceList from './home-sence'
import SkeletonDisplay from './skeleton'
const fetchData = async (id: string) => {
  return await HomeService.findById(id)
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
