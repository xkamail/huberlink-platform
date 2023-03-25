import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { CommandFlagEnum, IIRRemoteVirtualDevice } from '@/lib/types'
import { toSWR } from '@/lib/utils'
import DeviceService from '@/services/DeviceService'
import { BoxIcon, PowerIcon } from 'lucide-react'
import useSWR from 'swr'
import { Button } from '../ui/button'
import Card from '../ui/card'
import { Dialog, DialogTrigger } from '../ui/dialog'
const ThingVirtualRemote = ({
  deviceId,
  v,
}: {
  v: IIRRemoteVirtualDevice
  deviceId: string
}) => {
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
    <div className="col-span-6 md:col-span-4">
      <Dialog>
        <DialogTrigger asChild>
          <Card className="cursor-pointer hover:shadow-lg transition-all">
            <div className="flex items-center">
              <div className="flex-shrink-0 flex gap-2">
                <BoxIcon className="w-6 h-6" /> {v.name}
              </div>
            </div>
            <div className="mt-4 flex flex-row items-center justify-between">
              <div></div>
              {homeCommands.map((cmd) => (
                <Button
                  key={cmd.id}
                  variant="subtle"
                  size="circle"
                  className="flex items-center"
                >
                  <PowerIcon className="w-8 h-8 cursor-pointer text-indigo-600 " />
                </Button>
              ))}
            </div>
          </Card>
        </DialogTrigger>
      </Dialog>
    </div>
  )
}

export default ThingVirtualRemote

ThingVirtualRemote.displayName = 'ThingVirtualRemote'
