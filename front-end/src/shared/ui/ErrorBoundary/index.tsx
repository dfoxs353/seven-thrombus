import { useEffect, useState } from 'react';
import { Typography } from '@/shared/ui/index';

type TErrorBoundaryProps = {
  message: string;
};

export const ErrorBoundary = ({ message }: TErrorBoundaryProps) => {
  const [secondsRemaining, setSecondsRemaining] = useState<number>(5);

  useEffect(() => {
    if (secondsRemaining <= 0) return;

    const interval = setInterval(() => {
      setSecondsRemaining(prev => {
        if (prev <= 1) {
          clearInterval(interval);
          return 0;
        }
        return prev - 1;
      });
    }, 1000);

    return () => clearInterval(interval);
  }, [secondsRemaining]);

  if (secondsRemaining <= 0) return null;

  return (
    <>
      {secondsRemaining > 0 && (
        <section className='absolute top-4 left-[50%] translate-x-[-50%] flex w-full max-w-[345px] flex-col gap-2.5 px-3 py-4 bg-[var(--additional-light)] rounded-xl'>
          <Typography className='color-[var(--base-dark)] text-base font-bold'>Какая-то ошибка</Typography>
          <Typography className='color-[var(--additional-dark)] text-sm font-normal'>{message}</Typography>
        </section>
      )}
    </>
  );
};
