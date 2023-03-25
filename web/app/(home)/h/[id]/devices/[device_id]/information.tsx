import { Button } from '@/components/ui/button'
import Card from '@/components/ui/card'
import { useToast } from '@/hooks/use-toast'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { IDeviceDetail } from '@/lib/types'
import DeviceService from '@/services/DeviceService'
import dayjs from 'dayjs'
import { useEffect, useState } from 'react'
const DeviceInformation = ({ data }: { data: IDeviceDetail }) => {
  const homeId = useHomeSelector((s) => s.homeId)
  const { toast } = useToast()
  const [copied, setCopied] = useState(false)
  const [online, setOnline] = useState<'checking' | 'online' | 'offline'>(
    'checking'
  )
  //
  useEffect(() => {
    DeviceService.ping(homeId, data.id).then((r) => {
      setOnline(r.success && r.data ? 'online' : 'offline')
    })
  }, [])
  //
  const handleCopy = () => {
    navigator.clipboard.writeText(data.token)
    toast.succes('Copied to clipboard')
    setCopied(true)
    setTimeout(() => {
      setCopied(false)
    }, 1000)
  }

  return (
    <Card title="Information">
      <div className="flex flex-col">
        <div className="flex flex-row">
          <div className="flex flex-col w-1/2 gap-1">
            <p className="text-sm">Name</p>
            <p className="text-sm">Type</p>
            <p className="text-sm">Model</p>
            <p className="text-sm">Heartbeat</p>
            <p className="text-sm">Status</p>
          </div>
          <div className="w-1/2 flex flex-col items-end gap-1">
            <p className="text-sm">{data.name}</p>
            <p className="text-sm">{data.kind}</p>
            <p className="text-sm">{data.model}</p>
            <p className="text-sm">
              {data.latestHeartbeatAt ? (
                <span className="text-green-500">
                  {dayjs(data.latestHeartbeatAt).format('HH:mm:ss DD/MM/YYYY')}
                </span>
              ) : (
                <span className="text-red-500">Offline</span>
              )}
            </p>
            {online == 'checking' ? (
              <p>
                <span className="text-yellow-500">Checking...</span>
              </p>
            ) : (
              <p>
                {online == 'online' ? (
                  <span className="text-green-500">Online</span>
                ) : (
                  <span className="text-red-500">Offline</span>
                )}
              </p>
            )}
          </div>
        </div>
        <div className="mt-6 w-full">
          <Button block variant="subtle" onClick={handleCopy}>
            {copied ? 'Copied' : 'Copy Device Token'}
          </Button>
        </div>
      </div>
    </Card>
  )
}

export default DeviceInformation

DeviceInformation.displayName = 'DeviceInformation'
