'use client'
import Card from '@/components/ui/card'
import HomeService from '@/services/HomeService'
const fetchData = async (id: string) => {
  return await HomeService.findById(id)
}

const HomePage = ({ params }: { params: { id: string } }) => {
  return (
    <div className="grid grid-cols-12 gap-4">
      <div className="col-span-12">
        <div className="rounded-lg shadow bg-indigo-600 text-slate-100 p-4">
          <div className="flex items-center">
            <div className="flex-shrink-0">
              <img
                className="h-10 w-10 rounded-full"
                src="https://via.placeholder.com/400"
                alt=""
              />
            </div>
          </div>
          <div className="mt-4">
            <div className="text-base font-medium ">Pakorn Sangpeth</div>
            <div>
              Lorem Ipsum is simply dummy text of the printing and typesetting
              industry. Lorem Ipsum has been the standard dummy text ever since
              the 1500s, when an unknown printer took a galley of type and
              scrambled it to make a type specimen book.
            </div>
          </div>
        </div>
      </div>
      <div className="col-span-6">
        <Card />
      </div>
      <div className="col-span-6">
        <Card />
      </div>
    </div>
  )
}
export default HomePage
