import clsx from 'clsx';
import { ButtonHTMLAttributes, DetailedHTMLProps } from 'react';

type TButtonProps = {
  classname?: string;
} & DetailedHTMLProps<ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement>;

export const Button = ({ classname, children, ...props }: TButtonProps) => {
  return (
    <button className={clsx('button', classname)} {...props}>
      {children}
    </button>
  );
};
