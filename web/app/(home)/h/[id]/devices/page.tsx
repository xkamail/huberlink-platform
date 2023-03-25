'use client'
import { Icons } from '@/components/icons'
import { Button } from '@/components/ui/button'
import { TypographyH2 } from '@/components/ui/h2'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { IDeviceCard } from '@/lib/types'
import DeviceService from '@/services/DeviceService'
import { PlusIcon } from 'lucide-react'
import Link from 'next/link'
import { useEffect, useState } from 'react'
import useSWR from 'swr'
import DeviceSkeletons from './skeleton'

const HomeDevicesPage = () => {
  const homeId = useHomeSelector((s) => s.homeId)
  const { data, error, isLoading } = useSWR(`home-devices-${homeId}`, () =>
    DeviceService.list(homeId)
  )
  const [devices, setDevices] = useState<IDeviceCard[]>([])
  const [status, setStatus] = useState<'idle' | 'loading' | 'ok' | 'error'>(
    'idle'
  )
  useEffect(() => {
    if (isLoading) {
      setStatus('loading')
    }
    if (error) {
      setStatus('error')
    }
    if (data && data.success) {
      setDevices(data.data)
      setStatus('ok')
    }
    if (data && !data.success) {
      setStatus('error')
    }
    return () => {}
  }, [data, error, isLoading])

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
            <Link
              href={`/h/${homeId}/devices/${d.id}`}
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
                <Button to={`/h/${homeId}/devices/${d.id}`} variant="subtle">
                  <span className="text-sm">
                    <Icons.settings className="w-5 h-5" />
                  </span>
                </Button>
              </div>
            </Link>
          ))}
        {status === 'ok' && devices.length === 0 && (
          <div className="text-center col-span-full">
            <p className="text-slate-500 dark:text-slate-400">
              You {`don't`} have any devices yet. Create one!
            </p>
          </div>
        )}
      </div>
    </div>
  )
}

export default HomeDevicesPage

HomeDevicesPage.displayName = 'HomeDevicesPage'
