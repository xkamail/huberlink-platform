'use client'
import { useHomeSelector } from '@/lib/contexts/HomeContext'

const CreateDevicePage = () => {
  const homeId = useHomeSelector((s) => s.homeId)
  return <></>
}

export default CreateDevicePage

CreateDevicePage.displayName = 'CreateDevicePage'
