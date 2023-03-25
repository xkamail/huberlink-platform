'use client'
import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import PageHeader from '@/components/ui/page-header'
import Spinner from '@/components/ui/spinner'
import { useStatus } from '@/hooks/use-status'
import { useToast } from '@/hooks/use-toast'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import {
  IIRRemoteVirtualDevice,
  IIRRemoteVirtualDeviceCommand,
} from '@/lib/types'
import DeviceService from '@/services/DeviceService'
import { PlusIcon } from 'lucide-react'
import Image from 'next/image'
import { useEffect, useState } from 'react'
import CommandCard from './CommandCard'

import useSWR, { mutate } from 'swr'

const IRRemoteSettingPage = ({
  params: { id, device_id: deviceId, virtual_id: virtualId },
}: {
  params: {
    id: string
    device_id: string
    virtual_id: string
  }
}) => {
  const [isLearning, setIsLearning] = useState(false)
  const [status, setStatus] = useStatus('loading')
  const [data, setData] = useState<IIRRemoteVirtualDevice | null>(null)

  const { toast } = useToast()
  const homeId = useHomeSelector((s) => s.homeId)
  const {
    data: _data,
    error,
    isLoading,
  } = useSWR(`remote-setting-${deviceId}`, () =>
    DeviceService.ir.findVirtual({
      homeId,
      deviceId,
      virtualId,
    })
  )
  // watch data changes
  useEffect(() => {
    if (error || !_data?.success) {
      setStatus('error')
    }
    if (isLoading) {
      setStatus('loading')
    }
    if (_data && _data.success) {
      setData(_data.data)
      setButtons(_data.data.buttons)
      setIsLearning(_data.data.isLearning)
      setStatus('ok')
    }
  }, [error, isLoading, _data, setStatus])

  const [buttons, setButtons] = useState<IIRRemoteVirtualDeviceCommand[]>([])

  useEffect(() => {
    if (isLearning) {
      let xx = setInterval(() => {
        mutate(`remote-setting-${deviceId}`)
      }, 1000)
      return () => {
        clearInterval(xx)
      }
    }
  }, [isLearning])

  const onStartLearning = async () => {
    const res = await DeviceService.ir.startLearning({
      homeId,
      deviceId,
      virtualId,
    })
    if (!res.success) {
      toast.error(res.message)
    } else {
      setIsLearning(true)
    }
  }

  const onStopLearning = async () => {
    const res = await DeviceService.ir.stopLearning({
      homeId,
      deviceId,
      virtualId,
    })
    if (!res.success) {
      toast.error(res.message)
    } else {
      // toast.succes('Learning stopped')
      setIsLearning(false)
    }
  }

  if (status === 'notfound') {
    return (
      <div className="text-center w-full flex justify-center my-20">
        <h1>Not Found</h1>
      </div>
    )
  }
  if (status === 'error') {
    return (
      <div className="text-center w-full flex justify-center my-20">
        <h1>Error</h1>
      </div>
    )
  }
  if (status === 'loading' || !data)
    return (
      <div className="text-center w-full flex justify-center my-20">
        <Spinner />
      </div>
    )
  return (
    <div>
      <PageHeader title={data.name} />
      <div>
        <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
          {buttons.map((b) => (
            <CommandCard deviceId={deviceId} key={b.id} data={b} />
          ))}
          <Dialog open={isLearning}>
            <DialogTrigger asChild>
              <div className="p-4 bg-white shadow-sm rounded-lg flex items-center">
                <Button
                  onClick={onStartLearning}
                  variant="link"
                  className="mx-auto"
                >
                  <PlusIcon className="w-5 h-5 mr-2" />
                  Add New Button
                </Button>
              </div>
            </DialogTrigger>
            <DialogContent closeBtn={false} className="sm:max-w-[425px]">
              <DialogHeader>
                <DialogTitle>Start learning a new button</DialogTitle>
                <DialogDescription className="my-10">
                  <div className="mx-auto mb-10 p-4">
                    <Image
                      src={require('@/assets/images/remote-control.png')}
                      width={128}
                      height={128}
                      alt="remote control"
                      className="mx-auto"
                    />
                  </div>
                  <p>
                    Put your remote control in front of the IR sensor and press
                    a button.
                  </p>
                </DialogDescription>
              </DialogHeader>
              <DialogFooter>
                <Button variant="outline" block onClick={onStopLearning}>
                  Stop
                </Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        </div>
      </div>
    </div>
  )
}

export default IRRemoteSettingPage

IRRemoteSettingPage.displayName = 'IRRemoteSettingPage'
