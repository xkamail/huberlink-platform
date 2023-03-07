'use client'
import PageHeader from '@/components/ui/page-header'
import { useCallback, useEffect, useState } from 'react'

//
const DeviceDetailPage = ({ params: { id } }: { params: { id: string } }) => {
  const [data, setData] = useState<any>(null)
  const [status, setStatus] = useState<
    'idle' | 'loading' | 'ok' | 'error' | 'notfound'
  >('idle')
  const fetchData = useCallback(() => {}, [])

  useEffect(() => {
    fetchData()
  }, [])

  if (status === 'notfound')
    return (
      <>
        <p>Device not found</p>
      </>
    )
  return (
    <>
      <PageHeader title="xxxx" />
      <div className="bg-slate-100 w-full">xx</div>
    </>
  )
}

export default DeviceDetailPage

DeviceDetailPage.displayName = 'DeviceDetailPage'
