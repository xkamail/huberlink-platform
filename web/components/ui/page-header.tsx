import { ArrowLeftIcon } from 'lucide-react'
import { useRouter } from 'next/navigation'
import { Button } from './button'

const PageHeader = ({ title }: { title: string }) => {
  const router = useRouter()
  return (
    <div className="mb-4 flex flex-row justify-between items-center">
      <Button
        variant="ghost"
        onClick={() => {
          router.back()
        }}
      >
        <ArrowLeftIcon />
      </Button>

      <h2 className="scroll-m-20 pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0 dark:border-b-slate-700">
        {title}
      </h2>
    </div>
  )
}

export default PageHeader
