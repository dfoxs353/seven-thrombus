import { forwardRef } from 'react';
import { clsx } from 'clsx';

type InputProps = React.InputHTMLAttributes<HTMLInputElement> & {
  className?: string;
  error?: string;
  successMessage?: string;
};

export const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ className = '', error, successMessage, ...props }, ref) => {
    return (
      <div className="relative w-full">
        <input
          ref={ref}
          className={clsx(
            'input w-full transition-colors duration-200',
            {
              'border-[var(--error)]': error,
              'border-[var(--approve)]': !error && successMessage,
            },
            className
          )}
          {...props}
        />

        {(error || successMessage) && (
          <span
            className={clsx(
              'absolute top-[2px] left-[10px] text-[12px]',
              {
                'text-[var(--error)]': error,
                'text-[var(--approve)]': successMessage,
              }
            )}
          >
            {error || successMessage}
          </span>
        )}
      </div>
    );
  }
);

Input.displayName = 'Input';