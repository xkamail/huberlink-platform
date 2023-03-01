import BottomNavigation from '@/components/navigation/BottomNavigation'
import React from 'react'

export default function HomeLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <div className="bg-slate-100 h-screen">
      <div className="mt-2 md:container mx-4 md:mx-auto">{children}</div>
      <BottomNavigation />
    </div>
  )
}
