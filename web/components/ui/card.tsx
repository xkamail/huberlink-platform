import { cn } from '@/lib/utils'
import React from 'react'

const Card = ({
  children,
  title,
  className,
  renderAs,
}: {
  title?: string
  children: React.ReactNode
  className?: string
  renderAs?: React.ElementType
}) => {
  return (
    <div className={cn('rounded-lg shadow bg-white p-4', className)}>
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
