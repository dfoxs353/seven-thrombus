import { createLazyFileRoute, useNavigate } from '@tanstack/react-router';
import { useGetUser } from '../features/hooks/useGetUser';
import { Button, ErrorBoundary, Header, Loader, Typography } from '@/shared/ui';
import { useSignOut } from '../features/hooks/useSignOut';

export const Route = createLazyFileRoute('/user')({
  component: RouteComponent,
});

function RouteComponent() {
  const { data, error, isError, isLoading, isSuccess } = useGetUser();
  const { mutate } = useSignOut();
  const navigate = useNavigate({ from: '/user' });

  const handleSignOut = () => {
    mutate();
    navigate({ to: '/login' });
  };

  return (
    <main className='p-4 max-w-screen w-full'>
      <div>
        {isLoading && <Loader size='lg' color='white' />}
        {isError && <ErrorBoundary message={error.message} />}
        {isSuccess && <>
          <Typography>{data.username}</Typography>
          <Typography>{data.firstName}</Typography>
          <Typography>{data.lastName}</Typography>
          <Typography>{data.roles}</Typography>
        </>}
        <Button onClick={handleSignOut}>Выйти из аккаунта</Button>
      </div>
      <Header />
    </main>
  );
}
