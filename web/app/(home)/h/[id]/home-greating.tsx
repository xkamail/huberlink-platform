import { useUser } from '@/lib/hooks'

const HomeGreating = () => {
  const { profile } = useUser()
  return (
    <div className="rounded-lg shadow bg-indigo-600 text-slate-100 p-4">
      <div className="flex items-center">
        <div className="flex-shrink-0">
          <img
            className="h-10 w-10 rounded-full"
            src="https://via.placeholder.com/400"
            alt="x"
          />
        </div>
      </div>
      <div className="mt-4">
        <div className="text-base font-medium ">{profile.email}</div>
        <div>
          Lorem Ipsum is simply dummy text of the printing and typesetting
          industry. Lorem Ipsum has been the standard dummy text ever since the
          1500s, when an unknown printer took a galley of type and scrambled it
          to make a type specimen book.
        </div>
      </div>
    </div>
  )
}

export default HomeGreating

HomeGreating.displayName = 'HomeGreating'
