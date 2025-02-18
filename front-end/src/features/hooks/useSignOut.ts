import { useMutation } from '@tanstack/react-query';
import { put } from '@/shared/api/apiAbstractions';

export const useSignOut = () => {
  return useMutation({
    mutationKey: ['SignOut'],
    mutationFn: () => {
      return put('/signout');
    },
  });
};
