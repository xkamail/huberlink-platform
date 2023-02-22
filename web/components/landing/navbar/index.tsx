import { Button } from '@/components/ui/button'
import { cn } from '@/lib/utils'

const LandingNavbar = () => {
  return (
    <nav className={cn(`sticky z-50 top-0 w-full border-b border-b-slate-200`)}>
      <div className="container mx-auto h-16 flex items-center">
        <div>
          <a href="/" className="">
            <span className="text-indigo-500 text-lg font-bold hover:text-indigo-600 transition-colors">
              HuberLink
            </span>{' '}
          </a>
        </div>
        <div className="flex-end ml-auto">
          <div className="space-x-2">
            <Button variant="outline">
              <span className="text-sm uppercase">LOGIN</span>
            </Button>
            <Button variant="primary">
              <span className="text-sm uppercase">Try now</span>
            </Button>
          </div>
        </div>
      </div>
    </nav>
  )
}

export default LandingNavbar

LandingNavbar.displayName = 'LandingNavbar'
