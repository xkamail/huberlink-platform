import { Button } from '@/components/ui/button'
import { Dialog, DialogContent, DialogTrigger } from '@/components/ui/dialog'
import { IIRRemoteVirtualDeviceCommand } from '@/lib/types'
import { EditIcon } from 'lucide-react'

const CommandCard = ({ data }: { data: IIRRemoteVirtualDeviceCommand }) => {
  return (
    <div className="p-4 bg-white shadow-sm rounded-lg flex justify-between items-center">
      <p className="capitalize">{data.name || 'undefined'}</p>

      <Dialog>
        <DialogTrigger asChild>
          <Button variant="ghost">
            <EditIcon className="w-5 h-5" />
          </Button>
        </DialogTrigger>
        <DialogContent>asd</DialogContent>
      </Dialog>
    </div>
  )
}

export default CommandCard

CommandCard.displayName = 'CommandCard'
