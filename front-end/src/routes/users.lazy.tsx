import { createLazyFileRoute } from '@tanstack/react-router'
import { UsersPage } from '../pages/Users'

export const Route = createLazyFileRoute('/users')({
  component: RouteComponent,
})

function RouteComponent() {
  return <UsersPage />
}
