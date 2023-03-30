'use client'
import { Button } from '@/components/ui/button'
import { TypographyH2 } from '@/components/ui/h2'
import Spinner from '@/components/ui/spinner'
import { useHomeSelector } from '@/lib/contexts/HomeContext'
import { toSWR } from '@/lib/utils'
import HomeService from '@/services/HomeService'
import { PlusIcon } from 'lucide-react'
import useSWR from 'swr'

const AutomationPage = () => {
  const homeId = useHomeSelector((s) => s.homeId)
  const { data, error, isLoading } = useSWR(
    ['home/automtion', homeId],
    toSWR(HomeService.scenes.list(homeId))
  )
  if (error) return <div>error</div>

  return (
    <div>
      <div className="mb-4 flex justify-between items-center">
        <TypographyH2>My Automation</TypographyH2>
        <Button to={`/h/${homeId}/automation/create`} variant="outline-primary">
          <PlusIcon className="w-4 h-4 inline-block mr-1" /> Create
        </Button>
      </div>
      <div>
        {isLoading && (
          <div className="flex justify-center py-10">
            <Spinner />
          </div>
        )}

        {data && data.map((scene) => <div key={scene.id}>{scene.name}</div>)}
      </div>
    </div>
  )
}

export default AutomationPage

AutomationPage.displayName = 'AutomationPage'
