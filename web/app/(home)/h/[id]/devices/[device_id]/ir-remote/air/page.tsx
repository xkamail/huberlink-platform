'use client'

import { Icons } from '@/components/icons'
import PageHeader from '@/components/ui/page-header'
import { cn } from '@/lib/utils'
import { MinusIcon, PlusIcon, SnowflakeIcon } from 'lucide-react'
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
          <div className="flex flex-row gap-4 justify-center items-center">
            <div className="">
              <MinusIcon className={btnClassName} />
            </div>
            <div>
              <h1 className="text-6xl font-bold">25</h1>
            </div>
            <div>
              <PlusIcon className={btnClassName} />
            </div>
          </div>
          <div className="">
            <div className="flex flex-row gap-4 justify-center items-center">
              <ToggleButton
                label={`Mode: ${mode}`}
                loading={mode === 'loading'}
                onToggle={toggleMode}
              />
              <a href="" className="p-4">
                Speed: Auto
              </a>
            </div>
          </div>
        </div>
      </div>
    </>
  )
}

const ToggleButton = ({
  onToggle,
  loading,
  label,
}: {
  onToggle: () => void
  loading: boolean
  label: string
}) => {
  return (
    <button onClick={onToggle} className="p-4 gap-2">
      <div className="bg-white transition-all  shadow cursor-pointer text-center rounded-full p-4 flex justify-center items-center group">
        {loading ? (
          <Icons.spin className="w-10 h-10 group-hover:text-slate-700 animate-spin" />
        ) : (
          <SnowflakeIcon className="w-10 h-10 group-hover:text-slate-700" />
        )}
      </div>
      <p className="text-sm flex items-center">{label}</p>
    </button>
  )
}

export default AirSettingPage
