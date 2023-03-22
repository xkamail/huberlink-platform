'use client'

import PageHeader from '@/components/ui/page-header'
import { cn } from '@/lib/utils'
import Image from 'next/image'
import { useState } from 'react'

const AirSettingPage = () => {
  const btnClassName = cn(
    `cursor-pointer w-10 h-10 hover:text-slate-500 transition-all duration-200`
  )

  const [mode, setMode] = useState<'Cool' | 'Auto' | 'loading'>('Auto')
  const toggleMode = () => {
    setMode('loading')
    setTimeout(() => {
      setMode(mode === 'Auto' ? 'Cool' : 'Auto')
    }, 1000)
  }

  return (
    <>
      <PageHeader title={`Air - `} />
      <div className="mx-auto max-w-lg">
        <div className="flex flex-col gap-4">
          <div className="">
            <Image
              src={require(`@/assets/images/air-conditioner.png`)}
              alt="ait"
              width={512}
              height={512}
              className="w-24 h-24 mx-auto"
            />
          </div>
          <div className="flex flex-row gap-4 justify-center items-center"></div>
          <div className=""></div>
        </div>
      </div>
    </>
  )
}

export default AirSettingPage
