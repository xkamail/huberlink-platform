'use client'

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
  return (
    <div className="rounded-lg bg-gradient-to-r from-blue-500  to-indigo-600 p-4 mb-4 text-white">
      <div className="flex justify-between items-center">
        <div>
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
        </div>
        <h1 className="text-xl md:text-2xl font-bold text-right ">My Home</h1>
      </div>
    </div>
  )
}

export default HomeSelector

HomeSelector.displayName = 'HomeSelector'
