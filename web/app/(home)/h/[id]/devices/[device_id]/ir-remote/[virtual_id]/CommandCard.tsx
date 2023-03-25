import { Button } from '@/components/ui/button'
import { Dialog, DialogContent, DialogTrigger } from '@/components/ui/dialog'
import Form from '@/components/ui/form'
import FormInput from '@/components/ui/form-input'
import { IIRRemoteVirtualDeviceCommand } from '@/lib/types'
import { EditIcon } from 'lucide-react'
import { useForm } from 'react-hook-form'

const CommandCard = ({ data }: { data: IIRRemoteVirtualDeviceCommand }) => {
  const ctx = useForm()
  const submit = (data) => {
    //
  }
  return (
    <div className="p-4 bg-white shadow-sm rounded-lg flex justify-between items-center">
      <p className="capitalize">{data.name || 'undefined'}</p>

      <Dialog>
        <DialogTrigger asChild>
          <Button variant="ghost">
            <EditIcon className="w-5 h-5" />
          </Button>
        </DialogTrigger>
        <DialogContent>
          <Form ctx={ctx} onSubmit={submit}>
            <FormInput name="" />
          </Form>
        </DialogContent>
      </Dialog>
    </div>
  )
}

export default CommandCard

CommandCard.displayName = 'CommandCard'
