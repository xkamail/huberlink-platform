const Card = ({ children }: { children: React.ReactNode }) => {
  return <div className="rounded-lg shadow bg-white p-4">{children}</div>
}

export default Card

Card.displayName = 'Card'
