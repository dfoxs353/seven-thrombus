import { useMutation } from '@tanstack/react-query';
import { deleteRequest} from '@/shared/api/apiAbstractions';

export const useDeleteAllUsersAdmin = (id:number) => {
  return useMutation({
      mutationKey: ['DeleteUser'],
      mutationFn: async () => {
        const response = await deleteRequest(`/accounts/${id}`);
        return response.data;
      },
    });
};
