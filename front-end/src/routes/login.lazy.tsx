import { createLazyFileRoute } from '@tanstack/react-router';
import { LoginPage } from '../pages/Login';

export const Route = createLazyFileRoute('/login')({
  component: RouteComponent,
});

function RouteComponent() {
  return <LoginPage />;
}
