'use client'

import PageHeader from '@/components/ui/page-header'
import { useHomeSelector } from '@/lib/contexts/HomeContext'

const CreateScenePage = () => {
  const homeId = useHomeSelector((s) => s.homeId)
  //
  return (
    <div>
      <PageHeader title="Create Automation" />
    </div>
  )
}

export default CreateScenePage

CreateScenePage.displayName = 'CreateScenePage'
