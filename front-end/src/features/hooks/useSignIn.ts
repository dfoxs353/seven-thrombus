import { useMutation } from '@tanstack/react-query';
import { post } from '@/shared/api/apiAbstractions';
import { TSignIn } from '@/shared/models/SignIn';
import { TTokenPair } from '@/shared/models/TokenPair';

export const useSignIn = () => {
  return useMutation({
    mutationFn: (userData: TSignIn) => {
      return post<TTokenPair, TSignIn>('/signin', userData);
    }
  });
};
