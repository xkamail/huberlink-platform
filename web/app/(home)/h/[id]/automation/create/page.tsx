'use client'

import Card from '@/components/ui/card'
import Form from '@/components/ui/form'
import PageHeader from '@/components/ui/page-header'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { toSWR } from '@/lib/utils'
import DeviceService from '@/services/DeviceService'
import { useForm } from 'react-hook-form'
import useSWR from 'swr'

const CreateScenePage = () => {
  const ctx = useForm()
  const homeId = useHomeSelector((s) => s.homeId)
  const {
    data: devices,
    error: devicesError,
    isLoading: devicesLoading,
  } = useSWR(
    ['home/automtion/create/devices', homeId],
    toSWR(DeviceService.list(homeId))
  )
  const submit = () => {
    //
  }
  //
  return (
    <div>
      <PageHeader title="Create Automation" />
      <Card>
        <Form ctx={ctx} onSubmit={submit}>
          <div></div>
        </Form>
      </Card>
    </div>
  )
}

export default CreateScenePage

CreateScenePage.displayName = 'CreateScenePage'
