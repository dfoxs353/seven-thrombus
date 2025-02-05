import { TRole } from "./Role"

export type TUser = {
  firstName:string,
  id:number,
  lastName:string,
  password:string,
  roles: TRole[],
  username:string,
}

export type TUsers = Array<TUser>