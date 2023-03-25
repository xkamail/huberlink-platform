import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import Form from '@/components/ui/form'
import FormInput from '@/components/ui/form-input'
import { useToast } from '@/hooks/use-toast'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { IIRRemoteVirtualDeviceCommand } from '@/lib/types'
import DeviceService from '@/services/DeviceService'
import { EditIcon, SaveIcon, TrashIcon } from 'lucide-react'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { mutate } from 'swr'

const CommandCard = ({
  data,
  deviceId,
}: {
  data: IIRRemoteVirtualDeviceCommand
  deviceId: string
}) => {
  const [open, setOpen] = useState(false)
  const { toast } = useToast()
  const homeId = useHomeSelector((s) => s.homeId)
  const ctx = useForm({
    defaultValues: {
      name: data.name,
      remark: data.remark,
    },
  })
  const onDelete = async () => {
    const res = await DeviceService.ir.deleteCommand({
      homeId,
      deviceId,
      virtualId: data.virtualId,
      commandId: data.id,
    })
    if (!res.success) {
      toast.error(res.message)
      return
    }
    mutate(`remote-setting-${deviceId}`)
    setOpen(false)
  }
  const submit = async (payload: { name: string; remark: string }) => {
    //
    const res = await DeviceService.ir.updateCommand(
      {
        homeId,
        deviceId,
        virtualId: data.virtualId,
        commandId: data.id,
      },
      payload
    )
    if (!res.success) {
      toast.error(res.message)
      return
    }
    mutate(`remote-setting-${deviceId}`)
    setOpen(false)
  }
  return (
    <div className="p-4 bg-white shadow-sm rounded-lg flex justify-between items-center">
      <p className="capitalize">{data.name || 'undefined'}</p>

      <Dialog open={open} onOpenChange={setOpen}>
        <DialogTrigger asChild>
          <Button variant="ghost">
            <EditIcon className="w-5 h-5" />
          </Button>
        </DialogTrigger>
        <DialogContent className="sm:max-w-[425px]">
          <Form ctx={ctx} onSubmit={submit}>
            <DialogHeader>
              <DialogTitle>Edit button</DialogTitle>
            </DialogHeader>
            <div className="py-4 space-y-4">
              <FormInput
                name="name"
                placeholder="Name"
                label="Name"
                options={{
                  required: `Name is required`,
                }}
              />
              <FormInput name="remark" placeholder="Remark" label="Remark" />
            </div>
            <DialogFooter>
              <div className="flex justify-between w-full">
                <Button type="button" variant="destructive" onClick={onDelete}>
                  <TrashIcon className="w-4 h-4 mr-2" /> Delete
                </Button>
                <Button type="submit" variant="primary">
                  <SaveIcon className="w-4 h-4 mr-2" /> Save Setting
                </Button>
              </div>
            </DialogFooter>
          </Form>
        </DialogContent>
      </Dialog>
    </div>
  )
}

export default CommandCard

CommandCard.displayName = 'CommandCard'
