import { useEffect, useRef, useState } from 'react';

type TErrorBoundaryProps = {
  message: string;
};

export const ErrorBoundary = ({ message }: TErrorBoundaryProps) => {
  const [secondsTillTheEnd, setSecondsTillTheEnd] = useState<number>(5);
  const [width, setWidth] = useState<number>(100);
  const startTimeRef = useRef(Date.now());

  useEffect(() => {
    const animationFrame = requestAnimationFrame(function animate() {
      const elapsed = Date.now() - startTimeRef.current;
      const remainingTime = Math.max(5000 - elapsed, 0);

      setWidth((remainingTime / 5000) * 100);

      setSecondsTillTheEnd(Math.ceil(remainingTime / 1000));

      if (remainingTime > 0) {
        requestAnimationFrame(animate);
      }
    });

    return () => cancelAnimationFrame(animationFrame);
  }, []);

  return (
    <>
      {secondsTillTheEnd !== 0 && (
        <section className="absolute right-3 bottom-5 flex w-full max-w-[200px] flex-col border-2 border-red-500">
          {message}
          {secondsTillTheEnd}
          <div
            className="h-4 bg-red-500"
            style={{
              width: `${width}%`,
              transition: 'width 50ms linear'
            }}
          />
        </section>
      )}
    </>
  );
};
