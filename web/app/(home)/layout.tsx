import BottomNavigation from '@/components/navigation/BottomNavigation'

export default function HomeLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <div className="bg-slate-100 pb-[48px] mb-10">
      <div className="mt-2 md:container mx-4 md:mx-auto">{children}</div>
      <BottomNavigation />
    </div>
  )
}
