import { ResponseCode } from '@/lib/types'
import HomeService from '@/services/HomeService'
import { notFound } from 'next/navigation'

const fetchData = async (id: string) => {
  return await HomeService.findById(id)
}

const HomePage = async ({ params }: { params: { id: string } }) => {
  console.log('params', params)

  const res = await fetchData(params.id)
  if (!res.success) {
    if (res.code === ResponseCode.ResourceNotFound) {
      notFound()
    }
    throw new Error(res.message)
  }
  return (
    <div>
      adsasd
      <p></p>
    </div>
  )
}
export default HomePage
