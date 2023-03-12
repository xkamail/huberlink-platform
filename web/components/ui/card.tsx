const Card = ({
  children,
  title,
}: {
  title?: string
  children: React.ReactNode
}) => {
  return (
    <div className="rounded-lg shadow bg-white p-4">
      {title && (
        <div>
          <h2 className="mb-6 scroll-m-20 border-b border-b-slate-200 pb-2 text-2xl font-semibold tracking-tight first:mt-0 dark:border-b-slate-700">
            {title}
          </h2>
        </div>
      )}
      {children}
    </div>
  )
}

export default Card

Card.displayName = 'Card'
