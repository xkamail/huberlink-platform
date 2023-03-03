'use client'
import { Button } from '@/components/ui/button'
import { TypographyH2 } from '@/components/ui/h2'
import { PlusIcon } from 'lucide-react'

const AutomationPage = () => {
  return (
    <div>
      <div className="mb-4 flex justify-between items-center">
        <TypographyH2>My Automation</TypographyH2>
        <Button variant="outline-primary">
          <PlusIcon className="w-4 h-4 inline-block mr-1" /> Create
        </Button>
      </div>
    </div>
  )
}

export default AutomationPage

AutomationPage.displayName = 'AutomationPage'
