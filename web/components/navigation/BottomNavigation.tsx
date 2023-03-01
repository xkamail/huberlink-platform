'use client'
import { cn } from '@/lib/utils'
import { BotIcon, HomeIcon, NetworkIcon, SettingsIcon } from 'lucide-react'
import Link from 'next/link'
import { usePathname } from 'next/navigation'

const menuClass = (active: boolean) =>
  cn(
    `w-full focus:text-primary hover:primary justify-center inline-block text-center pt-2 pb-1 text-base`,
    active ? 'text-indigo-500' : 'text-slate-900'
  )

const BottomNavigation = () => {
  const path = usePathname()
  console.log('path', path)

  const currentHome = ''

  return (
    <div
      className={cn`block fixed inset-x-0 bottom-0 z-10 bg-white shadow border-t mb-0 p-safe`}
    >
      <div className="flex justify-between  container mx-auto">
        <Link href={`/h/${currentHome}`} className={menuClass(false)}>
          <HomeIcon className="w-5 h-5 mx-auto" />
          <span className="block text-xs">Home</span>
        </Link>
        <Link
          href={`/h/${currentHome}/automation`}
          className={menuClass(false)}
        >
          <BotIcon className="w-5 h-5 mx-auto" />
          <span className="block text-xs">Automation</span>
        </Link>
        <Link href={`/h/${currentHome}/devices`} className={menuClass(false)}>
          <NetworkIcon className="w-5 h-5 mx-auto" />
          <span className="block text-xs">Devices</span>
        </Link>
        <Link href="/account" className={menuClass(path === '/account')}>
          <SettingsIcon className="w-5 h-5 mx-auto" />
          <span className="block text-xs">Account</span>
        </Link>
      </div>
    </div>
  )
}

export default BottomNavigation

BottomNavigation.displayName = 'BottomNavigation'
