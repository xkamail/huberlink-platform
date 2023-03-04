'use client'
import { Button } from '@/components/ui/button'
import Form from '@/components/ui/form'
import FormInput from '@/components/ui/form-input'
import { Label } from '@/components/ui/label'
import PageHeader from '@/components/ui/page-header'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { DeviceKindEnum, DEVICE_CATEGORY, ICreateDeviceForm } from '@/lib/types'
import DeviceService from '@/services/DeviceService'
import { useForm } from 'react-hook-form'
import KindCard from './kind-card'

const CreateDevicePage = () => {
  const homeId = useHomeSelector((s) => s.homeId)

  const ctx = useForm<ICreateDeviceForm>({
    defaultValues: {
      name: '',
      model: '',
      kind: DeviceKindEnum.Unknown,
    },
  })
  const loading = ctx.formState.isSubmitting

  const submit = async (data: ICreateDeviceForm) => {
    await DeviceService.create(homeId, data)
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
            label="Model"
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
          <div className="flex justify-end">
            <Button type="submit" variant="default" loading={loading}>
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
