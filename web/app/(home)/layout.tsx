'use client'
import BottomNavigation from '@/components/navigation/BottomNavigation'
import { useUserSelector } from '@/lib/contexts/UserContext'
import { cn } from '@/lib/utils'
import { usePathname, useRouter } from 'next/navigation'

export default function HomeLayout({
  children,
}: {
  children: React.ReactNode
}) {
  const router = useRouter()
  const asPath = usePathname()
  const isLoggedIn = useUserSelector((s) => s.isLoggedIn)
  const done = useUserSelector((s) => s.status === 'success')
  const loading = useUserSelector(
    (s) => s.status === 'loading' || s.status === 'idle'
  )
  if (loading) {
    return <SkeletonAuth />
  }
  if (!isLoggedIn && done) {
    router.push(`/auth/sign-in?redirect=${asPath}`)
    return
  }
  return (
    <div className="bg-slate-100 h-[100vh] pb-[48px] mb-10">
      <div className="mt-2 md:container mx-4 md:mx-auto">{children}</div>
      <BottomNavigation />
    </div>
  )
}

const SkeletonAuth = () => {
  return (
    <div className="bg-slate-100 h-[100vh] pb-[48px] mb-10">
      <div className="mt-2 md:container mx-4 md:mx-auto">
        <div className="bg-white rounded-lg shadow p-4">
          <div className="animate-pulse flex space-x-4">
            <div className="flex-1 space-y-4 py-1">
              <div className="h-4 bg-slate-200 rounded w-3/4"></div>
              <div className="space-y-2">
                <div className="h-4 bg-slate-200 rounded"></div>
                <div className="h-4 bg-slate-200 rounded w-5/6"></div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div
        className={cn`block fixed inset-x-0 bottom-0 z-10 bg-white shadow border-t mb-0 p-safe`}
      ></div>
    </div>
  )
}
