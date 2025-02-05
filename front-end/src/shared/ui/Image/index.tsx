import clsx from "clsx"
import { HTMLAttributes } from "react"

type TImageProps = {
  classname?: string,
  image: string,
} & HTMLAttributes<HTMLImageElement>

export const Image = ({ image, classname, ...props }: TImageProps) => (
  <img src={image} className={clsx('w-full h-auto', classname)} {...props} />
)