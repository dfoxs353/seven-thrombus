import { useGetAllUsersAdmin } from "@/features/hooks/useGetAllUsersAdmin"
import { Button, ErrorBoundary, Header, List, Typography, UserElement } from "@/shared/ui"

export const UsersPage = () => {
  const { data, error, isSuccess, isError } = useGetAllUsersAdmin()
  return (
    <main className="page">
      <Typography as="h1" className="heading text-left mb-5">Список пользователей</Typography>
      <Button className="button button-secondary w-full mb-4">Создать пользователя</Button>
      {isSuccess && <List className="flex flex-col gap-4">
        {data.map(user => <UserElement key={user.id} {...user} />)}
      </List>}
      <Header />
      {isError && <ErrorBoundary message={error.message}/>}
    </main>
  )
}