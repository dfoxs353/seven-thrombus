import { useQuery } from '@tanstack/react-query';
import { get } from '@/shared/api/apiAbstractions';

export const useGetUsers = () => {
  const { data, error, isLoading, isError, isFetched } = useQuery({
    queryKey: ['Users'],
    queryFn: () => {
      get('/accounts');
    }
  });

  return {
    data,
    error,
    isError,
    isFetched,
    isLoading
  };
};
