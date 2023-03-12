import { RegisterOptions, useFormContext } from 'react-hook-form'
import { Input, InputProps } from './input'
import { Label } from './label'
interface IProps extends InputProps {
  name: string
  label?: string
  options?: RegisterOptions
}
const FormInput = ({ name, label, options, ...rest }: IProps) => {
  const {
    register,
    formState: { errors, dirtyFields, isSubmitting },
  } = useFormContext()
  const invalid = !!errors[name]
  return (
    <div className="space-y-1">
      {label && <Label htmlFor={name}>{label}</Label>}
      <Input
        id={name}
        {...register(name, options)}
        invalid={invalid}
        {...rest}
      />
      {/* render form hook error by name */}
      {invalid && errors[name]?.message && (
        <span className="text-red-500 text-xs inline-block">
          {`${errors[name]?.message}`}
        </span>
      )}
    </div>
  )
}

export default FormInput

FormInput.displayName = 'FormInput'
