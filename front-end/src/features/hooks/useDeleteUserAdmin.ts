import { useMutation } from '@tanstack/react-query';
import { deleteRequest} from '@/shared/api/apiAbstractions';

export const useDeleteAllUsersAdmin = () => {
  return useMutation({
      mutationKey: ['DeleteUser'],
    mutationFn: async (id: number) => {
        const response = await deleteRequest(`/accounts/${id}`);
        return response.data;
      },
    });
};
