import { createLazyFileRoute } from '@tanstack/react-router';
import { RegistrationPage } from '../pages/Registration';

export const Route = createLazyFileRoute('/registration')({
  component: RouteComponent,
});

function RouteComponent() {
  return <RegistrationPage />;
}
