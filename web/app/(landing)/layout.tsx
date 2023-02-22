import LandingNavbar from '@/components/landing/navbar'

export default function LandingLayouyt({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <div className="">
      <LandingNavbar />
      {children}
    </div>
  )
}
