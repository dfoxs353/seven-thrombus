import { Link, useNavigate } from '@tanstack/react-router';
import { Form, Typography, Image, Button, Input, ErrorBoundary, Loader } from '@/shared/ui';
import registration from '@/assets/images/regisration.webp';
import { yupResolver } from '@hookform/resolvers/yup';
import { useForm } from 'react-hook-form';
import { loginSchema } from '@/shared/schema/login';
import { useSignIn } from '@/features/hooks/useSignIn';
import { TSignIn } from '@/shared/models/SignIn';
import { useEffect } from 'react';
import { getSuccessMessage } from '@/features/utils/getSuccessMessage';

export const LoginPage = () => {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors, isValid, touchedFields },
  } = useForm({
    resolver: yupResolver(loginSchema),
    mode: 'onChange',
  });
  const { error, data, isError, isPending, mutate, isSuccess } = useSignIn();
  const [usernameValue, passwordValue] = watch(['username', 'password']);
  const navigate = useNavigate({ from: '/login' });

  const onSubmit = (data: TSignIn) => {
    mutate(data);
  };

  useEffect(() => {
    if (isSuccess && data) {
      localStorage.setItem('token', data.data.accessToken);
      localStorage.setItem('refreshToken', data.data.refreshToken);
      navigate({ to: '/user' });
    }
  }, [isSuccess, data, navigate]);

  return (
    <div className='flex h-screen flex-col items-center justify-center p-4'>
      <Form classname='flex flex-col gap-4' onSubmit={handleSubmit(onSubmit)}>
        <Image image={registration}></Image>
        <Typography as='h1' className='heading mb-3'>
          Добро пожаловать!
        </Typography>

        <Input
          placeholder="Введите логин"
          {...register('username')}
          error={errors.username?.message}
          successMessage={getSuccessMessage(
            'username',
            'Логин заполнен корректно',
            touchedFields,
            errors,
            usernameValue
          )}
        />

        <Input
          placeholder="Введите пароль"
          type='password'
          {...register('password')}
          error={errors.password?.message}
          successMessage={getSuccessMessage(
            'password',
            'Пароль заполнен корректно',
            touchedFields,
            errors,
            passwordValue
          )}
        />

        <Button
          className='button button-primary flex items-center justify-center'
          type='submit'
          disabled={!isValid || isPending}
        >
          {isPending ? <Loader color='white' size='sm' /> : <Typography as='p'>Войти</Typography>}
        </Button>
      </Form>
      <Link to='/registration' className='mt-5'>
        <Typography as='p' className='inline'>
          У меня нет аккаунта. Зарегистрироваться
        </Typography>
      </Link>
      {isError && <ErrorBoundary message={error.message} />}
      {isSuccess && <ErrorBoundary message={'Вы успешно зашли, но страницы пользователя пока нет'} />}
    </div>
  );
};
