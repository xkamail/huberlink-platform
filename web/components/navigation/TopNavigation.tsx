import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { Icons } from '../icons'

type IProps = {
  title: string
  leftContent?: React.ReactNode
}
const TopNavigation = ({ title, leftContent }: IProps) => {
  const loading = useHomeSelector((s) => s?.isLoading || false)
  return (
    <div className="rounded-lg shadow bg-white p-4 mb-4">
      <div className="flex justify-between items-center">
        <div>{leftContent}</div>
        <h1 className="text-xl md:text-xl font-bold text-right text-slate-700 ">
          {!loading && title}
          <Icons.settings className="cursor-pointer inline ml-2 w-6 h-6 text-slate-500 hover:text-slate-600 transition-colrs" />
        </h1>
      </div>
    </div>
  )
}

export default TopNavigation

TopNavigation.displayName = 'TopNavigation'
