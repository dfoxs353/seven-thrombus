import { useMutation } from '@tanstack/react-query';
import { post } from '@/shared/api/apiAbstractions';
import { TSignUp } from '@/shared/models/SingUp';

export const useSignUp = () => {
  return useMutation({
    mutationKey: ['SignUp'],
    mutationFn: (userData: TSignUp) => {
      return post<TSignUp, any>('/signup', userData);
    },
  });
};
