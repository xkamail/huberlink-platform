'use client'

import TopNavigation from '@/components/navigation/TopNavigation'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { ListIcon, PlusIcon } from 'lucide-react'
import Link from 'next/link'

const HomeSelector = () => {
  const homeList = [
    {
      id: 1,
      name: 'Dorm',
    },
    {
      id: 2,
      name: 'Home',
    },
  ]
  const homeTitle = 'Home'
  return (
    <TopNavigation
      leftContent={
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <ListIcon className="cursor-pointer hover:text-slate-200 transition-colors h-8 w-8" />
          </DropdownMenuTrigger>
          <DropdownMenuContent align="start">
            <DropdownMenuLabel>Select home</DropdownMenuLabel>
            <DropdownMenuSeparator />
            {homeList.map((home) => (
              <Link href={`/h/${home.id}`} key={home.id}>
                <DropdownMenuItem>{home.name}</DropdownMenuItem>
              </Link>
            ))}
            <DropdownMenuSeparator />
            <Link href="/h/create">
              <DropdownMenuItem>
                <PlusIcon className="h-4 w-4 mr-2" />
                Create new home
              </DropdownMenuItem>
            </Link>
          </DropdownMenuContent>
        </DropdownMenu>
      }
      title={homeTitle}
    />
  )
}

export default HomeSelector

HomeSelector.displayName = 'HomeSelector'
