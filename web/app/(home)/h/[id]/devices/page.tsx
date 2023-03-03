'use client'
import { Icons } from '@/components/icons'
import { Button } from '@/components/ui/button'
import { TypographyH2 } from '@/components/ui/h2'
import { PlusIcon } from 'lucide-react'

const HomeDevicesPage = () => {
  return (
    <div>
      <div className="mb-4 flex justify-between items-center">
        <TypographyH2>Your devices</TypographyH2>
        <Button variant="outline-primary">
          <PlusIcon className="w-4 h-4 inline-block mr-1" /> Create
        </Button>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        {[...new Array(5)].map((_, i) => (
          <div
            className="flex justify-between items-center rounded-lg p-4 bg-white shadow transition-all cursor-pointer hover:shadow-lg"
            key={i}
          >
            <div className="">
              <div className="flex items-center">
                <Icons.bot className="h-8 w-8 text-slate-800" />
              </div>
            </div>
            <div className=" text-left w-full px-4">asd</div>
            <div className="">
              <Button variant="subtle">
                <span>x</span>
              </Button>
            </div>
          </div>
        ))}
        <div className="text-center col-span-2">
          <Button variant="link">Load more</Button>
        </div>
      </div>
    </div>
  )
}

export default HomeDevicesPage

HomeDevicesPage.displayName = 'HomeDevicesPage'
