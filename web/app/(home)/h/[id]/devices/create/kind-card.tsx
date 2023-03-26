import { DeviceKindEnum } from '@/lib/types'
import { cn } from '@/lib/utils'
import {
  CircleDotIcon,
  EyeIcon,
  HelpCircle,
  ToggleRightIcon,
} from 'lucide-react'
import { useFormContext } from 'react-hook-form'

interface IProps {
  kind: string | DeviceKindEnum
  icon: string
  label: string
  description: string
  onChange: (kind: DeviceKindEnum) => void
}

const KindCard = ({ kind, icon, label, description, onChange }: IProps) => {
  const { setValue, watch } = useFormContext()
  const v = watch('kind')
  const onClick = () => {
    setValue('kind', kind)
  }
  const active = v === kind

  return (
    <div
      className={cn(
        'cursor-pointer flex flex-col items-center justify-center w-20 h-20 rounded-lg gap-1  transition-all',
        active ? 'bg-indigo-500 shadow' : 'bg-slate-100 hover:bg-slate-200'
      )}
      onClick={onClick}
    >
      {renderIcon(icon, active)}
      <span
        className={cn(
          `block text-sm font-medium text-slate-900`,
          active ? 'text-slate-100' : 'text-slate-800'
        )}
      >
        {label}
      </span>
    </div>
  )
}

const renderIcon = (icon: string, active: boolean) => {
  const iconClass = cn(`w-8 h-8`, active ? 'text-slate-100' : 'text-gray-500')
  if (icon === 'unknown') {
    return <HelpCircle className={iconClass} />
  }
  if (icon === 'sensor') {
    return <EyeIcon className={iconClass} />
  }
  if (icon === 'switch') {
    return <ToggleRightIcon className={iconClass} />
  }
  if (icon === 'remote') {
    return <CircleDotIcon className={iconClass} />
  }
  return <HelpCircle className={iconClass} />
}

export default KindCard

KindCard.displayName = 'KindCard'
