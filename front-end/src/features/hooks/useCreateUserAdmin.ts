import { useMutation } from '@tanstack/react-query';
import { post} from '@shared/api/apiAbstractions';
import { TUser } from '@shared/models/User';

export const useCreateAllUsersAdmin = () => {
  return useMutation({
    mutationKey: ['CreateUser'],
    mutationFn: async (user:TUser) => {
        const response = await post<TUser, any>(`/accounts`, user);
        return response.data;
      },
    });
};
