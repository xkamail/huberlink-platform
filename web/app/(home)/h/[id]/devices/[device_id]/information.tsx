import { Button } from '@/components/ui/button'
import Card from '@/components/ui/card'
import { IDeviceDetail } from '@/lib/types'
import dayjs from 'dayjs'
import { useEffect, useState } from 'react'
const DeviceInformation = ({ data }: { data: IDeviceDetail }) => {
  const [copied, setCopied] = useState(false)
  useEffect(() => {}, [])
  return (
    <Card title="Information">
      <div className="flex flex-col">
        <div className="flex flex-row">
          <div className="flex flex-col w-1/2 gap-1">
            <p className="text-sm">Name</p>
            <p className="text-sm">Type</p>
            <p className="text-sm">Model</p>
            <p className="text-sm">Heartbeat</p>
          </div>
          <div className="w-1/2 flex flex-col items-end gap-1">
            <p className="text-sm">{data.name}</p>
            <p className="text-sm">{data.kind}</p>
            <p className="text-sm">{data.model}</p>
            <p className="text-sm">
              {data.latestHeartbeatAt ? (
                dayjs(data.latestHeartbeatAt).format('HH:mm:ss DD/MM/YYYY')
              ) : (
                <span className="text-red-500">Offline</span>
              )}
            </p>
          </div>
        </div>
        <div className="mt-6 w-full">
          <Button block variant="subtle">
            {copied ? 'Copied' : 'Copy Device Token'}
          </Button>
        </div>
      </div>
    </Card>
  )
}

export default DeviceInformation

DeviceInformation.displayName = 'DeviceInformation'