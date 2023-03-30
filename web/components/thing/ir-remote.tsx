'use client'

import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { toSWR } from '@/lib/utils'
import DeviceService from '@/services/DeviceService'
import useSWR from 'swr'
import ThingVirtualRemote from './virtual-remote'
type IProps = {
  deviceId: string
}
const IRRemoteThingCard = ({ deviceId }: IProps) => {
  const homeId = useHomeSelector((s) => s.homeId)
  const { data, error, isLoading } = useSWR(
    [`ir-remote-card`, deviceId, homeId],
    toSWR(DeviceService.ir.findDetail({ deviceId, homeId }))
  )
  if (!data) return null
  //
  return (
    <>
      {data.virtuals.map((v) => (
        <ThingVirtualRemote deviceId={deviceId} v={v} key={v.id} />
      ))}
    </>
  )
}

const Card = ({ children }: { children: React.ReactNode }) => {
  return (
    <a
      href=""
      className="col-span-6 md:col-span-4 bg-white transition-all hover:shadow-lg cursor-pointer rounded-lg p-4 shadow"
    >
      {children}
    </a>
  )
}

export default IRRemoteThingCard

IRRemoteThingCard.displayName = 'IRRemoteThingCard'
