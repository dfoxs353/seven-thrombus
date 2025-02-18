import { ReactSVG } from "react-svg"
import { TUser } from "@shared/models/User"
import { Typography, List, Input, Form, Button, ErrorBoundary } from "@shared/ui"
import { ROLES } from "@/shared/constants/roles"
import disciplines from '@/assets/images/icons/disciplines.svg'
import { useVisible } from "@/features/hooks/useVisible"
import { Modal } from "@/widjets"
import { useDeleteAllUsersAdmin } from "@/features/hooks/useDeleteUserAdmin"
import { useQueryClient } from "@tanstack/react-query"
import { MouseEvent } from "react"

export const UserElement = ({ firstName, lastName, roles, username, id }: TUser) => {
  const { visible, close, toggle } = useVisible(false)
  const { mutate, isSuccess, error, isError } = useDeleteAllUsersAdmin()
  const queryClient = useQueryClient()

  const handleClick = (e: MouseEvent<HTMLButtonElement, globalThis.MouseEvent>) => {
    e.preventDefault()
    mutate(id)
  }

  isSuccess && queryClient.invalidateQueries({ queryKey: ['Users'] })

  return (
    <li className="border-[1px] border-[var(--additional-light)] rounded-xl p-3 flex flex-col gap-4">
    <List className="flex flex-wrap gap-1.5 select-none">
      {roles.map(role => <li key={role} className="bg-[var(--additional-light)] rounded-md p-3">
        <Typography className="font-normal text-xs text-[var(--base-dark)]">{ROLES[role]}</Typography>
      </li>)}
    </List>
    <div className="flex justify-between">
      <div className="flex gap-0.5 select-none">
        <Typography>{firstName}</Typography>
        <Typography>{lastName}</Typography>
        <Typography>{username}</Typography>
      </div>
        <ReactSVG src={disciplines} onClick={toggle} className="transition-colors duration-300 hover:stroke-[var(--main-accent)] stroke-[var(--additional-dark)] cursor-pointer" />
    </div>
      <Modal
        title={'Редактор пользователя'}
        isOpen={visible}
        onClose={close}>
        <Form>
          <Input placeholder="Имя пользователя" value={firstName} />
          <Input placeholder="Фамилия пользователя" value={lastName} />
          <Input placeholder="Логин" value={username} />
          <Input placeholder="Роль" value={ROLES[roles[0]]} />
          <Button onClick={(e) => handleClick(e)} className="button button-secondary border-[var(--error)] text-[var(--error)]">Удалить пользователя</Button>
        </Form>
      </Modal>
      {isError && <ErrorBoundary message={error.message} />}
      {isSuccess && <ErrorBoundary message={'Вы успешно удалили пользователя'} />}
  </li>
  )
}