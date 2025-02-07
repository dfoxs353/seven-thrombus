import { createLazyFileRoute } from '@tanstack/react-router'
import { useGetUser } from '../features/hooks/useGetUser'
import { ErrorBoundary, Loader } from '../shared/ui'

export const Route = createLazyFileRoute('/user')({
  component: RouteComponent,
})

function RouteComponent() {

  const {
    data,error,isError,isLoading,isSuccess
  } = useGetUser()

  console.log(data)

  return <div>
    {isLoading && <Loader size='lg' color='white'/>}
    {isError && <ErrorBoundary message={error.message}/>}
    {isSuccess && <div>{data.username}</div>}
  </div>
}
