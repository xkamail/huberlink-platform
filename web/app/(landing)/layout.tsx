import LandingNavbar from '@/components/landing/navbar'

export default function LandingLayouyt({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <div className="">
      <LandingNavbar />
      <div className="container mx-auto">{children}</div>
    </div>
  )
}
