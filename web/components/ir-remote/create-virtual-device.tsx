import Form from '@/components/ui/form'
import { useToast } from '@/hooks/use-toast'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { VIRTUAL_CATEGORY } from '@/lib/types'
import DeviceService from '@/services/DeviceService'
import { useForm } from 'react-hook-form'
import { Button } from '../ui/button'
import { DialogFooter } from '../ui/dialog'
import FormInput from '../ui/form-input'
import VirtualKindCard from './virtual-kind-card'

const CreateVirtualDeviceForm = ({
  deviceId,
  onSuccess,
}: {
  deviceId: string
  onSuccess: () => void
}) => {
  const { toast } = useToast()
  const homeId = useHomeSelector((s) => s.homeId)
  const form = useForm({
    defaultValues: {
      name: '',
      kind: '',
      icon: '',
    },
  })
  const onSubmit = async (data: any) => {
    if (data.kind === '') {
      toast.error('Please select a kind')
      return
    }
    //
    const res = await DeviceService.ir.createVirtual({
      homeId,
      deviceId,
      ...data,
    })
    if (!res.success) {
      toast.error(res.message)
      return
    }
    onSuccess()
    toast.succes(`Virtual device ${data.name} created`)
  }
  return (
    <Form ctx={form} onSubmit={onSubmit}>
      <div className="grid gap-4 py-4">
        <FormInput
          name="name"
          className="col-span-3"
          label="Name"
          options={{
            required: 'Name is required',
            maxLength: {
              value: 10,
              message: 'Name must be less than 10 characters',
            },
          }}
        />
        <div className="mt-4 gap-4 flex-row flex justify-center">
          {VIRTUAL_CATEGORY.map((v) => (
            <VirtualKindCard
              kind={v.kind}
              key={v.name}
              icon={v.icon}
              label={v.name}
            />
          ))}
        </div>
      </div>
      <DialogFooter>
        <Button
          type="submit"
          loading={form.formState.isSubmitting}
          variant="primary"
        >
          Create
        </Button>
      </DialogFooter>
    </Form>
  )
}

export default CreateVirtualDeviceForm

CreateVirtualDeviceForm.displayName = 'CreateVirtualDeviceForm'
