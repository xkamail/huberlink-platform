import { Button } from './button'

const Card = () => {
  return (
    <div className="rounded-lg shadow bg-white p-4">
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
          <p className="text-sm my-4">
            Lorem Ipsum is simply dummy text of the printing and typesetting
            industry.
          </p>
        </div>
        <div className="flex justify-end">
          <Button variant="primary" size="sm">
            <span>Read more</span>
          </Button>
        </div>
      </div>
    </div>
  )
}

export default Card

Card.displayName = 'Card'
