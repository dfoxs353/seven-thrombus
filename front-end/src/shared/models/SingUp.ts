import { TUser } from './User';

export type TSignUp = Omit<TUser, 'id' | 'roles'>;
