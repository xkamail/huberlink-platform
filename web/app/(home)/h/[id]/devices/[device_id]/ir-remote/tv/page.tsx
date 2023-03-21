'use client'

import PageHeader from '@/components/ui/page-header'

const TVSettingPage = () => {
  return (
    <>
      <PageHeader title={`TV - `} />
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4"></div>
    </>
  )
}

export default TVSettingPage

TVSettingPage.displayName = 'TVSettingPage'
