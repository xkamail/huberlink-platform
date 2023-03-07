import { cn } from '@/lib/utils'
import { FanIcon, LampIcon } from 'lucide-react'

const HomeSenceList = () => {
  return (
    <div className="grid grid-cols-12 gap-2">
      {[...Array(4)].map((_, i) => (
        <div
          className={cn(
            `transition-colors hover:bg-indigo-600 hover:text-slate-100 col-span-4 lg:col-span-2 cursor-pointer rounded-lg p-2 border text-indigo-600 border-indigo-600 group`,
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
  )
}

export default HomeSenceList

HomeSenceList.displayName = 'HomeSenceList'
