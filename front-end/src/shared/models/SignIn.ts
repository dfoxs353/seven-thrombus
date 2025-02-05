import { TUser } from "./User";

export type TSignIn = Pick<TUser, 'password'> & Pick<TUser, 'username'>