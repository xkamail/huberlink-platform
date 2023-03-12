'use client'

import Card from '@/components/ui/card'
import { TypographyH2 } from '@/components/ui/h2'
import { Separator } from '@/components/ui/sperator'

const AccountPage = () => {
  return (
    <div className="mt-4 space-y-4">
      <div className="flex justify-between items-center">
        <TypographyH2>Account</TypographyH2>
        <div className="w-24 h-24 bg-slate-200 rounded-full"></div>
      </div>
      <Separator />
      <div className="">
        <Card>
          <div className="">
            <div className="flex justify-between items-center">
              <div className="">
                <div className="text-sm text-gray-500">Name</div>
                <div className="text-lg font-medium">John Doe</div>
              </div>
            </div>
          </div>
        </Card>
      </div>
    </div>
  )
}

export default AccountPage

AccountPage.displayName = 'AccountPage'
