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
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { IHome } from '@/lib/types'
import HomeService from '@/services/HomeService'
import { ListIcon, PlusIcon } from 'lucide-react'
import Link from 'next/link'
import { useCallback, useEffect, useState } from 'react'

const HomeSelector = () => {
  const [homeList, setHomeList] = useState<IHome[]>([])
  const fetchData = useCallback(async () => {
    const r = await HomeService.list()
    if (r.success) {
      setHomeList(r.data)
    } else {
      setHomeList([])
    }
  }, [])
  useEffect(() => {
    fetchData()
  }, [fetchData])

  const homeTitle = useHomeSelector((s) => s.homeName)
  const loading = useHomeSelector((s) => s.isLoading)
  return (
    <TopNavigation
      leftContent={
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <ListIcon className="cursor-pointer hover:text-slate-500 transition-colors h-8 w-8" />
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
