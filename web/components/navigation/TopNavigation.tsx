import { useHomeSelector } from '@/lib/contexts/HomeContext'

type IProps = {
  title: string
  leftContent?: React.ReactNode
}
const TopNavigation = ({ title, leftContent }: IProps) => {
  const loading = useHomeSelector((s) => s.isLoading)
  return (
    <div className="rounded-lg shadow bg-white p-4 mb-4">
      <div className="flex justify-between items-center">
        <div>{leftContent}</div>
        <h1 className="text-xl md:text-2xl font-bold text-right ">
          {!loading && title}
        </h1>
      </div>
    </div>
  )
}

export default TopNavigation

TopNavigation.displayName = 'TopNavigation'
