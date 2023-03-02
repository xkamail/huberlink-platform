'use client'
import HomeService from '@/services/HomeService'
const fetchData = async (id: string) => {
  return await HomeService.findById(id)
}

const HomePage = ({ params }: { params: { id: string } }) => {
  return (
    <div>
      <p></p>
    </div>
  )
}
export default HomePage
