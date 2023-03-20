'use client'
import { Button } from '@/components/ui/button'
import Form from '@/components/ui/form'
import FormInput from '@/components/ui/form-input'
import { Label } from '@/components/ui/label'
import PageHeader from '@/components/ui/page-header'
import { useToast } from '@/hooks/use-toast'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { DeviceKindEnum, DEVICE_CATEGORY, ICreateDeviceForm } from '@/lib/types'
import DeviceService from '@/services/DeviceService'
import { useRouter } from 'next/navigation'
import { useForm } from 'react-hook-form'
import KindCard from './kind-card'

const CreateDevicePage = () => {
  const { toast } = useToast()
  const router = useRouter()
  const homeId = useHomeSelector((s) => s.homeId)

  const ctx = useForm<ICreateDeviceForm>({
    defaultValues: {
      name: '',
      model: '',
      kind: DeviceKindEnum.IRRemote,
    },
  })
  const loading = ctx.formState.isSubmitting

  const submit = async (data: ICreateDeviceForm) => {
    const res = await DeviceService.create(homeId, data)
    if (res.success) {
      toast.succes(`Device ${data.name} created!`)
      router.push(`/h/${homeId}/devices`)
      return
    }
    toast.error(res.message)
  }

  return (
    <>
      <PageHeader title="Create device" />
      <div className="rounded-lg bg-white shadow p-4">
        <Form ctx={ctx} onSubmit={submit}>
          <FormInput
            name="name"
            label="Device name"
            options={{
              required: 'Name is required',
            }}
          />
          <FormInput
            label="Model (optional)"
            name="model"
            options={
              {
                // required: 'Model is required',
              }
            }
          />
          <div className="space-y-1">
            <Label>Category</Label>
            <div className="flex flex-wrap gap-4 w-full">
              {DEVICE_CATEGORY.map((c, i) => (
                <KindCard
                  kind={c.kind}
                  key={i}
                  onChange={(k) => {}}
                  icon={c.icon}
                  label={c.name}
                  description={c.description}
                />
              ))}
            </div>
          </div>
          <div className="flex justify-between">
            <Button
              type="button"
              variant="destructive"
              to={`/h/${homeId}/devices`}
            >
              Cancel
            </Button>
            <Button type="submit" variant="primary" loading={loading}>
              Create device
            </Button>
          </div>
        </Form>
      </div>
    </>
  )
}

export default CreateDevicePage

CreateDevicePage.displayName = 'CreateDevicePage'
