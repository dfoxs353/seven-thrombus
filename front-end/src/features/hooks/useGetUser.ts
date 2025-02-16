import { useQuery } from '@tanstack/react-query';
import { get } from '@/shared/api/apiAbstractions';
import { TUser } from '@/shared/models/User';

export const useGetUser = () => {
  return useQuery<TUser>({
    queryKey: ['User'],
    queryFn: async () => {
      const response = await get<TUser>('/accounts/me');
      return response.data;
    },
  });
};
