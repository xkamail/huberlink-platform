'use client'
import Card from '@/components/ui/card'
import { cn } from '@/lib/utils'
import HomeService from '@/services/HomeService'
import { FanIcon, LampIcon } from 'lucide-react'
const fetchData = async (id: string) => {
  return await HomeService.findById(id)
}

const HomePage = ({ params }: { params: { id: string } }) => {
  return (
    <div className="grid grid-cols-12 gap-4">
      <div className="col-span-12">
        <div className="grid grid-cols-12 gap-2">
          {[...Array(4)].map((_, i) => (
            <div
              className={cn(
                `transition-colors hover:bg-indigo-600 hover:text-slate-100 col-span-2 cursor-pointer rounded-lg p-2 border text-indigo-600 border-indigo-600 group`,
                i == 2 && 'bg-indigo-600 text-slate-100'
              )}
              key={i}
            >
              <div className="flex flex-row justify-between">
                <div className="flex items-center">
                  {i == 2 ? (
                    <FanIcon className="w-6 h-6 mr-2 animate-spin duration-1000" />
                  ) : (
                    <LampIcon className="w-6 h-6 mr-2" />
                  )}{' '}
                </div>
                <div className="text-right flex flex-col items-end">
                  <span className="text-lg">เปิดไฟ</span>
                  <span
                    className={cn(
                      `text-xs`,
                      `font-thin`,
                      i == 2 && 'text-slate-100'
                    )}
                  >
                    on
                  </span>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
      <div className="col-span-12">
        <div className="rounded-lg shadow bg-indigo-600 text-slate-100 p-4">
          <div className="flex items-center">
            <div className="flex-shrink-0">
              <img
                className="h-10 w-10 rounded-full"
                src="https://via.placeholder.com/400"
                alt=""
              />
            </div>
          </div>
          <div className="mt-4">
            <div className="text-base font-medium ">Pakorn Sangpeth</div>
            <div>
              Lorem Ipsum is simply dummy text of the printing and typesetting
              industry. Lorem Ipsum has been the standard dummy text ever since
              the 1500s, when an unknown printer took a galley of type and
              scrambled it to make a type specimen book.
            </div>
          </div>
        </div>
      </div>
      <div className="col-span-6">
        <Card />
      </div>
      <div className="col-span-6">
        <Card />
      </div>
    </div>
  )
}
export default HomePage
