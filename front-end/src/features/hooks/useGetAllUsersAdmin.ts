import { useQuery } from '@tanstack/react-query';
import { get } from '@/shared/api/apiAbstractions';
import { TUsers } from '../../shared/models/User';

export const useGetAllUsersAdmin = () => {
  return useQuery<TUsers>({
      queryKey: ['Users'],
      queryFn: async () => {
        const response = await get<TUsers>('/accounts');
        return response.data;
      },
    });
};
