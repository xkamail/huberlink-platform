import HomeSelector from './home-select'

export default function HomeLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <>
      <HomeSelector />
      {children}
    </>
  )
}
