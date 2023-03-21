'use client'

import { IIRRemoteVirtualDevice, VirtualCategoryEnum } from '@/lib/types'
import { ChevronRight, TrashIcon } from 'lucide-react'
import Image from 'next/image'
import { Button } from '../ui/button'
import Card from '../ui/card'

const renderCategory = (category: VirtualCategoryEnum) => {
  let img = require(`@/assets/images/remote-control.png`)
  if (category === VirtualCategoryEnum.AirConditioner) {
    img = require(`@/assets/images/air-conditioner.png`)
  }
  if (category === VirtualCategoryEnum.TV) {
    img = require(`@/assets/images/tv.png`)
  }
  if (category === VirtualCategoryEnum.Fan) {
    img = require(`@/assets/images/fan.png`)
  }
  if (category === VirtualCategoryEnum.Speaker) {
    img = require(`@/assets/images/speaker.png`)
  }

  return (
    <div className="mx-auto flex justify-center">
      <Image
        src={img}
        alt={category.toString()}
        width={256}
        height={256}
        className="w-12 h-12 md:w-16 md:h-16"
      />
    </div>
  )
}

const VirtualDevice = ({ name, category }: IIRRemoteVirtualDevice) => {
  return (
    <Card>
      <div className="grid grid-cols-2 gap-4">
        <div className="col-span-1">
          <div className="">{renderCategory(category)}</div>
        </div>
        <div className="col-span-1 space-y-4">
          <div className="flex w-full justify-between items-center">
            <h1 className="text-xl text-slate-900 font-semibold tracking-tight">
              {name}
            </h1>
            <Button size="sm" variant="link" className="hover:text-red-500">
              <TrashIcon className="w-4 h-4" />
            </Button>
          </div>
          <div className="">
            <Button size="sm" variant="subtle" block>
              <span className="mb-1">Settings</span>
              <ChevronRight className="w-4 h-4 ml-1" />
            </Button>
          </div>
        </div>
      </div>
    </Card>
  )
}

export default VirtualDevice

VirtualDevice.displayName = 'VirtualDevice'
