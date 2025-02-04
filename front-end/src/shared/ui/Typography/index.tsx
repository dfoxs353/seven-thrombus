import { HTMLAttributes } from "react";

type TTypographyProps = HTMLAttributes<HTMLParagraphElement>;

export const Typography = ({children, ...props}:TTypographyProps) => (
  <p {...props}>{children}</p>
)