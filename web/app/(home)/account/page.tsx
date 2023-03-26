'use client'

import Card from '@/components/ui/card'
import { TypographyH2 } from '@/components/ui/h2'
import { Separator } from '@/components/ui/sperator'
import SetPasswordForm from './set-password-form'

const AccountPage = () => {
  return (
    <div className="mt-4 space-y-4">
      <div className="flex justify-between items-center">
        <TypographyH2>Account</TypographyH2>
        <div className="w-16 h-16 bg-slate-200 rounded-full"></div>
      </div>
      <Separator />
      <div className="">
        <Card title="Set your new password">
          <div className="">
            <SetPasswordForm />
          </div>
        </Card>
      </div>
    </div>
  )
}

export default AccountPage

AccountPage.displayName = 'AccountPage'
