import clsx from "clsx"
import { HTMLAttributes } from "react"

type TFormProps = {
  classname?: string,
} & HTMLAttributes<HTMLFormElement>

export const Form = ({ classname, children, ...props }: TFormProps) => {
  return (
    <form className={clsx('flex flex-col gap-4', classname)} {...props}>
      {children}
    </form>
  )
}