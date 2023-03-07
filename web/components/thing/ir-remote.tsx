'use client'

import { Button } from '@/components/ui/button'
import { BoxIcon, PowerIcon, PowerOffIcon } from 'lucide-react'
import { useState } from 'react'

type IProps = {}
const IRRemoteThingCard = ({}: IProps) => {
  const [virtualDevices, setVirtualDevices] = useState([])
  // fetch virtual device
  // render virtual device first
  // then append real device

  const handlePower = (e: any) => {
    e.preventDefault()
  }
  return (
    <>
      <Card>
        <div className="flex items-center">
          <div className="flex-shrink-0 flex gap-2">
            <BoxIcon className="w-6 h-6" /> TV
          </div>
        </div>
        <div className="mt-4 flex flex-row items-center justify-between">
          <div></div>
          <Button
            onClick={handlePower}
            variant="subtle"
            size="circle"
            className="flex items-center"
          >
            <PowerIcon className="w-8 h-8 cursor-pointer text-indigo-500 " />
          </Button>
        </div>
      </Card>
      <div className="col-span-6 md:col-span-4 bg-white rounded-lg p-4 shadow">
        <div className="flex items-center">
          <div className="flex-shrink-0 flex gap-2">
            <BoxIcon className="w-6 h-6" /> Air
          </div>
        </div>
        <div className="mt-4 flex flex-row items-center justify-between">
          <div></div>
          <Button
            onClick={handlePower}
            variant="subtle"
            size="circle"
            className="flex items-center"
          >
            <PowerOffIcon className="w-8 h-8 cursor-pointer text-red-500 " />
          </Button>
        </div>
      </div>
    </>
  )
}

const Card = ({ children }: { children: React.ReactNode }) => {
  return (
    <a
      href=""
      className="col-span-6 md:col-span-4 bg-white transition-all hover:shadow-lg cursor-pointer rounded-lg p-4 shadow"
    >
      {children}
    </a>
  )
}

export default IRRemoteThingCard

IRRemoteThingCard.displayName = 'IRRemoteThingCard'
