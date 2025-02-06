import { useMutation } from '@tanstack/react-query';
import { post } from '@/shared/api/apiAbstractions';
import { TSignUp } from '../../shared/models/SingUp';

export const useSignUp = () => {
  return useMutation({
    mutationFn: (userData: TSignUp) => {
      return post<TSignUp>('/signup', userData);
    }
  });
};
