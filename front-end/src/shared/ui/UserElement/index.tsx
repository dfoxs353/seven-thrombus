import { ReactSVG } from "react-svg"
import { TUser } from "@shared/models/User"
import { Typography, List } from "@shared/ui"
import { ROLES } from "@/shared/constants/roles"
import disciplines from '@/assets/images/icons/disciplines.svg'

export const UserElement = ({ firstName, lastName, roles, username }: TUser) => (
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
      <ReactSVG src={disciplines} className="transition-colors duration-300 hover:stroke-[var(--main-accent)] stroke-[var(--additional-dark)] cursor-pointer"/>
    </div>
  </li>
)