import GoHomeButton from '@/app/go-home-button'
import { cn } from '@/lib/utils'

const LandingNavbar = () => {
  return (
    <nav
      className={cn(
        `sticky z-50 top-0 w-full border-b border-b-slate-200 bg-white`
      )}
    >
      <div className="container px-4 md:px-0 mx-auto h-16 flex items-center">
        <div>
          <a href="/" className="">
            <span className="text-indigo-500 text-lg font-bold hover:text-indigo-600 transition-colors">
              HuberLink
            </span>{' '}
          </a>
        </div>
        <div className="flex-end ml-auto">
          <div className="space-x-2">
            <GoHomeButton />
          </div>
        </div>
      </div>
    </nav>
  )
}

export default LandingNavbar

LandingNavbar.displayName = 'LandingNavbar'
