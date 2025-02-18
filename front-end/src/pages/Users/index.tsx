import { useGetAllUsersAdmin } from "@/features/hooks/useGetAllUsersAdmin"
import { Button, ErrorBoundary, Form, Header, Input, List, Typography, UserElement } from "@/shared/ui"
import { useVisible } from "@/features/hooks/useVisible"
import { Modal } from "@/widjets"

export const UsersPage = () => {
  const { data, error, isSuccess, isError } = useGetAllUsersAdmin()
  const { visible, close, toggle } = useVisible(false)
  return (
    <main className="page">
      <Typography as="h1" className="heading text-left mb-5">Список пользователей</Typography>
      <Button className="button button-secondary w-full mb-4" onClick={toggle}>Создать пользователя</Button>
      {isSuccess && <List className="flex flex-col gap-4">
        {data.map(user => <UserElement key={user.id} {...user} />)}
      </List>}
      <Header />
      {isError && <ErrorBoundary message={error.message}/>}
      <Modal
        title={'Создание пользователя'}
        isOpen={visible}
        onClose={close}>
        <Form>
          <Input placeholder="Фамилия пользователя" />
          <Input placeholder="Имя пользователя" />
          <Input placeholder="Логин" />
          <Input placeholder="Пароль" />
          <Input placeholder="Роль" />
        </Form>
      </Modal>
    </main>
  )
}