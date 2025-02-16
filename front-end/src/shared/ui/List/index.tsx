import {  HTMLAttributes, ReactNode } from "react";

type TListProps = {
  className?:string;
  children: ReactNode;
} & HTMLAttributes<HTMLUListElement>

export const List = ({children,className, ...props}:TListProps) => (
  <ul className={`${className}`} {...props}>
    {children}
  </ul>
)