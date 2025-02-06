import clsx from 'clsx';

type TypographyProps<T extends React.ElementType> = {
  as?: T;
  classname?: string;
} & React.ComponentPropsWithoutRef<T>;

export const Typography = <T extends React.ElementType = 'p'>({
  as,
  classname = '',
  children,
  ...props
}: TypographyProps<T>) => {
  const Component = as || 'p';
  return (
    <Component className={clsx('text-center', classname)} {...props}>
      {children}
    </Component>
  );
};
