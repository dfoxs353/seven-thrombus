import { TUser } from './User';

export type TUpdateProfile = Omit<TUser, 'id' | 'roles' | 'username'>;
