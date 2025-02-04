import { createRootRoute, Outlet } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/router-devtools'
import { Layout } from '../widjets'

export const Route = createRootRoute({
  component: () => (
    <>
      <Layout>
        <Outlet />
        <TanStackRouterDevtools />
      </Layout>
    </>
  ),
})