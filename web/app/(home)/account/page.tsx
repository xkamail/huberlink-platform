'use client'

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
        Lorem Ipsum is simply dummy text of the printing and typesetting
        industry. Lorem Ipsum has been the industry standard dummy text ever
        since the 1500s, when an unknown printer took a galley of type and
        scrambled it to make a type specimen book. It has survived not only five
        centuries, but also the leap into electronic typesetting, remaining
        essentially unchanged. It was popularised in the 1960s with the release
        of Letraset sheets containing Lorem Ipsum passages, and more recently
        with desktop publishing software like Aldus PageMaker including versions
        of Lorem Ipsum.
      </div>
    </div>
  )
}

export default AccountPage

AccountPage.displayName = 'AccountPage'
