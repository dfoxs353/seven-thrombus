import { createLazyFileRoute, Link } from '@tanstack/react-router';
import { Form, Typography, Image, Button } from '../shared/ui';
import registration from '@/assets/images/regisration.webp';
import { Input } from '../shared/ui/Input';
import { yupResolver } from '@hookform/resolvers/yup';
import { registrationSchema } from '@/shared/schema/registration';
import { useForm } from 'react-hook-form';

export const Route = createLazyFileRoute('/registration')({
  component: RouteComponent
});

function RouteComponent() {
  const {
    register,
    handleSubmit,
    watch,
    formState: { errors, isValid, touchedFields }
  } = useForm({
    resolver: yupResolver(registrationSchema),
    mode: 'onChange'
  });

  const loginValue = watch('login');
  const passwordValue = watch('password');
  const confirmPasswordValue = watch('confirmPassword');

  const onSubmit = (data: any) => {
    console.log('Форма отправлена', data);
  };

  return (
    <div className="flex h-screen flex-col items-center justify-center p-4">
      <Form classname="flex flex-col gap-4" onSubmit={handleSubmit(onSubmit)}>
        <Image image={registration}></Image>
        <Typography as="h1" className="heading mb-3">
          Добро пожаловать!
        </Typography>
        <Input
          placeholder="Введите логин"
          className={`input w-full ${touchedFields.login && loginValue
            ? errors.login
              ? 'border-[var(--error)]'
              : 'border-[var(--approve)]'
            : 'border-gray-300'
            }`}
          {...register('login')}
        />
        {errors.login && (
          <Typography className="absolute top-[2px] left-[10px] text-[12px] text-[var(--error)]">
            {errors.login.message}
          </Typography>
        )}

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

        <div className="relative w-full">
          <Input
            type="password"
            placeholder="Повторите пароль"
            className={`input w-full ${touchedFields.confirmPassword && confirmPasswordValue
              ? errors.confirmPassword
                ? 'border-[var(--error)]'
                : 'border-[var(--approve)]'
              : 'border-gray-300'
              }`}
            {...register('confirmPassword')}
          />
          {errors.confirmPassword && (
            <Typography className="absolute top-[2px] left-[10px] text-[12px] text-[var(--error)]">
              {errors.confirmPassword.message}
            </Typography>
          )}
        </div>
        <Button
          className="button button-primary"
          type="submit"
          disabled={!isValid}
        >
          <Typography as="p">Зарегистрироваться</Typography>
        </Button>
      </Form>
      <Link to="/login" className="mt-5">
        <Typography as="p" className="inline">
          У меня уже есть аккаунт. Войти
        </Typography>
      </Link>
    </div>
  );
}
