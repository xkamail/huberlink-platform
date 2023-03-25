'use client'

import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '@/components/ui/alert-dialog'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { IIRRemoteVirtualDevice, VirtualCategoryEnum } from '@/lib/types'
import DeviceService from '@/services/DeviceService'
import { ChevronRight, TrashIcon } from 'lucide-react'
import Image from 'next/image'
import Link from 'next/link'
import { mutate } from 'swr'
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
        className="w-12 h-12 md:w-16 md:h-16  grayscale"
      />
    </div>
  )
}

const VirtualDevice = ({
  name,
  category,
  deviceId,
  id: virtualId,
}: IIRRemoteVirtualDevice & { deviceId: string }) => {
  const homeId = useHomeSelector((s) => s.homeId)
  const handleDelete = () => {
    //
    DeviceService.ir
      .deleteVirtual({
        homeId,
        virtualId,
        deviceId,
      })
      .finally(() => {
        mutate(`device-ir-remote`)
      })
  }
  //
  let href = `/h/${homeId}/devices/${deviceId}/ir-remote/${virtualId}`

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

            <AlertDialog>
              <AlertDialogTrigger asChild>
                <Button size="sm" variant="link" className="hover:text-red-500">
                  <TrashIcon className="w-4 h-4" />
                </Button>
              </AlertDialogTrigger>
              <AlertDialogContent>
                <AlertDialogHeader>
                  <AlertDialogTitle>
                    Are you sure absolutely sure?
                  </AlertDialogTitle>
                  <AlertDialogDescription>
                    This action cannot be undone. This will permanently delete
                    your virtual devices and all commands.
                  </AlertDialogDescription>
                </AlertDialogHeader>
                <AlertDialogFooter>
                  <AlertDialogCancel>Cancel</AlertDialogCancel>
                  <AlertDialogAction onClick={handleDelete}>
                    Continue
                  </AlertDialogAction>
                </AlertDialogFooter>
              </AlertDialogContent>
            </AlertDialog>
          </div>
          <div className="">
            <Link href={href}>
              <Button size="sm" variant="subtle" block>
                <span className="mb-1">Settings</span>
                <ChevronRight className="w-4 h-4 ml-1" />
              </Button>
            </Link>
          </div>
        </div>
      </div>
    </Card>
  )
}

export default VirtualDevice

VirtualDevice.displayName = 'VirtualDevice'
