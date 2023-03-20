import CreateVirtualDeviceForm from '@/components/ir-remote/create-virtual-device'
import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import { IDeviceDetail } from '@/lib/types'
import { useState } from 'react'

const IRRemoteSection = ({ device }: { device: IDeviceDetail }) => {
  const [open, setOpen] = useState(false)

  return (
    <div className="grid md:grid-cols-2 gap-4">
      <div></div>
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
