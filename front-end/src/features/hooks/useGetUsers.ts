import { useQuery } from '@tanstack/react-query';
import { get } from '@/shared/api/apiAbstractions';

export const useGetUsers = () => {
  return useQuery({
    queryKey: ['Users'],
    queryFn: () => {
      get('/accounts');
    }
  });
};
