import Card from '@/components/ui/card'

const PreHomePage = () => {
  return (
    <div className="grid md:grid-cols-12 grid-cols-6 gap-4">
      <div className="col-span-6">
        <Card />
      </div>
      <div className="col-span-6">
        <Card />
      </div>
    </div>
  )
}

export default PreHomePage
