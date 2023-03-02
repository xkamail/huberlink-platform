'use client'
import { HomeContextProvider } from '@/lib/contexts/HomeContext'
import HomeSelector from './home-select'
import PreRenderHome from './pre-render'

export default function HomeLayout({
  children,
  params,
}: {
  children: React.ReactNode
  params: { id: string }
}) {
  return (
    <HomeContextProvider homeId={params.id}>
      <>
        <HomeSelector />
        <PreRenderHome>{children}</PreRenderHome>
      </>
    </HomeContextProvider>
  )
}
