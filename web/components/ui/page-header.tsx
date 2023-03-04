import { TypographyH2 } from './h2'

const PageHeader = ({ title }: { title: string }) => {
  return (
    <div className="mb-4">
      <TypographyH2>{title}</TypographyH2>
    </div>
  )
}

export default PageHeader
