import { useToast } from '@/hooks/use-toast'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import {
  CommandFlagEnum,
  IIRRemoteVirtualDevice,
  IIRRemoteVirtualDeviceCommand,
} from '@/lib/types'
import { toSWR } from '@/lib/utils'
import DeviceService from '@/services/DeviceService'
import { BoxIcon } from 'lucide-react'
import Link from 'next/link'
import { useState } from 'react'
import useSWR from 'swr'
import { Button } from '../ui/button'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '../ui/dialog'
const ThingVirtualRemote = ({
  deviceId,
  v,
}: {
  v: IIRRemoteVirtualDevice
  deviceId: string
}) => {
  const { toast } = useToast()
  const homeId = useHomeSelector((s) => s.homeId)
  const { data, error, isLoading } = useSWR(
    ['home-virtual-device', v.id],
    toSWR(
      DeviceService.ir.listCommand({
        deviceId,
        homeId,
        virtualId: v.id,
      })
    )
  )

  const commands = data || []

  const homeCommands = commands.filter((x) => {
    return (x.flag & CommandFlagEnum.HomeScreen) == CommandFlagEnum.HomeScreen
  })

  //
  return (
    <div className="col-span-12 ">
      <Dialog>
        <DialogTrigger asChild>
          <div className="w-full rounded-lg shadow bg-white p-4  cursor-pointer hover:shadow-lg transition-all">
            <div className="flex items-center">
              <div className="flex-shrink-0 flex gap-2">
                <BoxIcon className="w-6 h-6" /> {v.name}
              </div>
            </div>
            <div className="mt-4 flex flex-row gap-2 justify-end">
              <div></div>
              {homeCommands.length === 0 && <div className="h-[40px]"></div>}
              {homeCommands.map((cmd) => (
                <CommandButton
                  deviceId={deviceId}
                  virtualId={v.id}
                  cmd={cmd}
                  key={cmd.id}
                />
              ))}
            </div>
          </div>
        </DialogTrigger>
        <DialogContent
          onOpenAutoFocus={(e) => {
            // disable auto focus button
            e.preventDefault()
          }}
          className=""
        >
          <DialogHeader>
            <DialogTitle>{v.name}</DialogTitle>
            <DialogDescription>
              <div className="py-4 flex gap-2">
                {commands.map((cmd) => (
                  <CommandButton
                    deviceId={deviceId}
                    virtualId={v.id}
                    cmd={cmd}
                    key={cmd.id}
                  />
                ))}
              </div>
              <div className="mt-10 mx-auto w-full flex justify-center">
                <Link
                  href={`/h/${homeId}/devices/${deviceId}/ir-remote/${v.id}`}
                >
                  <Button variant="link">Setting</Button>
                </Link>
              </div>
            </DialogDescription>
          </DialogHeader>
        </DialogContent>
      </Dialog>
    </div>
  )
}

const CommandButton = ({
  cmd,
  deviceId,
  virtualId,
}: {
  cmd: IIRRemoteVirtualDeviceCommand
  deviceId: string
  virtualId: string
}) => {
  const [isLoading, setIsLoading] = useState(false)
  const { toast } = useToast()
  const homeId = useHomeSelector((s) => s.homeId)

  const runCommand = (cmd: IIRRemoteVirtualDeviceCommand) => {
    return async (e: any) => {
      e.preventDefault()
      setIsLoading(true)
      const res = await DeviceService.ir
        .executeCommand(
          {
            deviceId,
            homeId,
            virtualId,
          },
          cmd.id
        )
        .finally(() => {
          setIsLoading(false)
        })
      if (!res.success) {
        toast.error(res.message)
      }
    }
  }
  return (
    <Button
      key={cmd.id}
      variant="subtle"
      onClick={runCommand(cmd)}
      className="flex items-center capitalize"
      disabled={isLoading}
    >
      {cmd.name}
    </Button>
  )
}
export default ThingVirtualRemote

ThingVirtualRemote.displayName = 'ThingVirtualRemote'
