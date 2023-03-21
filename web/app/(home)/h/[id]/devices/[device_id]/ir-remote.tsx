import CreateVirtualDeviceForm from '@/components/ir-remote/create-virtual-device'
import VirtualDevice from '@/components/ir-remote/virtual-device'
import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { IDeviceDetail } from '@/lib/types'
import DeviceService from '@/services/DeviceService'
import { useState } from 'react'
import useSWR from 'swr'
const IRRemoteSection = ({ device }: { device: IDeviceDetail }) => {
  const homeId = useHomeSelector((s) => s.homeId)
  const { data, isLoading, error } = useSWR(
    '/api/ir-remote',
    () =>
      DeviceService.ir.findDetail({
        deviceId: device.id,
        homeId,
      }),
    { refreshInterval: 1000 }
  )
  const [open, setOpen] = useState(false)
  if (isLoading) return <div>Loading...</div>
  if (error) return <div>Error</div>
  return (
    <div className="grid md:grid-cols-2 gap-4">
      {data &&
        data.success &&
        data.data.virtuals.map((v) => (
          <div className="col-span-1" key={v.id}>
            <VirtualDevice {...v} />
          </div>
        ))}
      <div className="col-span-full">
        <Dialog open={open} onOpenChange={setOpen}>
          <DialogTrigger asChild>
            <div className="text-center">
              <Button block size="lg" variant="primary">
                <span>Create new virtual device</span>
              </Button>
            </div>
          </DialogTrigger>
          <DialogContent className="sm:max-w-[425px] md:max-w-[565px]">
            <DialogHeader>
              <DialogTitle>Create Virtual Device</DialogTitle>
            </DialogHeader>

            <CreateVirtualDeviceForm
              deviceId={device.id}
              onSuccess={() => {
                setOpen(false)
              }}
            />
          </DialogContent>
        </Dialog>
      </div>
    </div>
  )
}

export default IRRemoteSection

IRRemoteSection.displayName = 'IRRemoteSection'
