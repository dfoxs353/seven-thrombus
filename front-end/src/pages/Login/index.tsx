import { Link } from '@tanstack/react-router';
import { Form, Typography, Image, Button, Input, ErrorBoundary, Loader } from '@/shared/ui';
import registration from '@/assets/images/regisration.webp';
import { yupResolver } from '@hookform/resolvers/yup';
import { useForm } from 'react-hook-form';
import { loginSchema } from '@/shared/schema/login';
import { useSignIn } from '@/features/hooks/useSignIn';
import { TSignIn } from '@/shared/models/SignIn';

export const LoginPage = () => {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors, isValid, touchedFields }
  } = useForm({
    resolver: yupResolver(loginSchema),
    mode: 'onChange'
  });
  const { error, data, isError, isPending, mutate, isSuccess } = useSignIn()
  const loginValue = watch('username');
  const passwordValue = watch('password');

  const onSubmit = (data: TSignIn) => {
    mutate(data)
  };

  isSuccess && localStorage.setItem("token", data.data.accessToken)

  return (
    <div className="flex h-screen flex-col items-center justify-center p-4">
      <Form classname="flex flex-col gap-4" onSubmit={handleSubmit(onSubmit)}>
        <Image image={registration}></Image>
        <Typography as="h1" className="heading mb-3">
          Добро пожаловать!
        </Typography>

        <div className="relative w-full">
          <Input
            placeholder="Введите логин"
            className={`input w-full ${touchedFields.username && loginValue
                ? errors.username
                  ? 'border-[var(--error)]'
                  : 'border-[var(--approve)]'
                : 'border-gray-300'
              }`}
            {...register('username')}
          />
          {errors.username && (
            <Typography className="absolute top-[2px] left-[10px] text-[12px] text-[var(--error)]">
              {errors.username.message}
            </Typography>
          )}
        </div>

        <div className="relative w-full">
          <Input
            type="password"
            placeholder="Введите пароль"
            className={`input w-full ${touchedFields.password && passwordValue
                ? errors.password && passwordValue
                  ? 'border-[var(--error)]'
                  : 'border-[var(--approve)]'
                : 'border-gray-300'
              }`}
            {...register('password')}
          />
          {errors.password && (
            <Typography className="absolute top-[2px] left-[10px] text-[12px] text-[var(--error)]">
              {errors.password.message}
            </Typography>
          )}
        </div>

        <Button
          className="button button-primary flex items-center justify-center"
          type="submit"
          disabled={!isValid || isPending}
        >
          {isPending ? (
            <Loader color="white" size="sm" />
          ) : (
            <Typography as="p">Войти</Typography>
          )}
        </Button>
      </Form>
      <Link to="/registration" className="mt-5">
        <Typography as="p" className="inline">
          У меня нет аккаунта. Зарегистрироваться
        </Typography>
      </Link>
      {isError && <ErrorBoundary message={error.message} />}
      {isSuccess && <ErrorBoundary message={"Вы успешно зашли, но страницы пользователя пока нет"} />}
    </div>
  );
};
