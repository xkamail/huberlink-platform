'use client'
import { Icons } from '@/components/icons'
import { Button } from '@/components/ui/button'
import { TypographyH2 } from '@/components/ui/h2'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { IDeviceCard } from '@/lib/types'
import DeviceService from '@/services/DeviceService'
import { PlusIcon } from 'lucide-react'
import { useCallback, useEffect, useState } from 'react'
import DeviceSkeletons from './skeleton'

const HomeDevicesPage = () => {
  const homeId = useHomeSelector((s) => s.homeId)

  const [devices, setDevices] = useState<IDeviceCard[]>([])
  const [status, setStatus] = useState<'idle' | 'loading' | 'ok' | 'error'>(
    'idle'
  )
  const fetchData = useCallback(async () => {
    setStatus('loading')
    const res = await DeviceService.list(homeId)
    if (res.success) {
      setDevices(res.data)
      setStatus('ok')
    } else {
      setStatus('error')
    }
  }, [homeId])

  useEffect(() => {
    fetchData()
  }, [fetchData])

  return (
    <div>
      <div className="mb-4 flex justify-between items-center">
        <TypographyH2>Your devices</TypographyH2>
        <Button to={`/h/${homeId}/devices/create`} variant="outline-primary">
          <PlusIcon className="w-4 h-4 inline-block mr-1" /> Create
        </Button>
      </div>
      <div className=" grid gap-4">
        {status === 'loading' && <DeviceSkeletons />}
        {status === 'ok' &&
          devices.map((d, i) => (
            <div
              className="col-span-1 flex justify-between items-center rounded-lg p-4 bg-white shadow transition-all cursor-pointer hover:shadow-lg"
              key={i}
            >
              <div className="">
                <div className="flex items-center">
                  <Icons.bot className="h-8 w-8 text-slate-800" />
                </div>
              </div>
              <div className=" text-left w-full px-4">{d.name}</div>
              <div className="">
                <Button variant="subtle">
                  <span>x</span>
                </Button>
              </div>
            </div>
          ))}
        <div className="text-center col-span-full">
          <Button variant="link">Load more</Button>
        </div>
      </div>
    </div>
  )
}

export default HomeDevicesPage

HomeDevicesPage.displayName = 'HomeDevicesPage'
