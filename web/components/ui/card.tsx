const Card = () => {
  return (
    <div className="rounded-lg bg-white p-4">
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
        <div className="text-base font-medium text-gray-900">
          Pakorn Sangpeth
        </div>
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

export default Card

Card.displayName = 'Card'
