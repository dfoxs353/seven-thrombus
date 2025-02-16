import { Link } from '@tanstack/react-router';
import { Form, Typography, Image, Button, Input, Loader, ErrorBoundary } from '@/shared/ui';
import registration from '@/assets/images/regisration.webp';
import { yupResolver } from '@hookform/resolvers/yup';
import { registrationSchema } from '@/shared/schema/registration';
import { useForm } from 'react-hook-form';
import { useSignUp } from '@/features/hooks/useSignUp';
import { useState } from 'react';
import { TSignUp } from '@/shared/models/SingUp';
import { getSuccessMessage } from '../../features/utils/getSuccessMessage';

export const RegistrationPage = () => {
  const [step, setStep] = useState(1);
  const {
    register,
    handleSubmit,
    watch,
    trigger,
    formState: { errors, isValid, touchedFields, isSubmitting },
  } = useForm<TSignUp & { confirmPassword: string }>({
    resolver: yupResolver(registrationSchema),
    mode: 'onChange',
    shouldUnregister: true,
  });

  const [firstNameValue, lastNameValue, usernameValue, passwordValue, confirmPasswordValue] = watch(['firstName', 'lastName', 'username', 'password', 'confirmPassword']);
  const { error, isError, isPending, mutate, isSuccess } = useSignUp();

  const handleContinue = async () => {
    const isValidStep1 = await trigger(['firstName', 'lastName']);
    if (isValidStep1) {
      setStep(2);
    } else {
      return <ErrorBoundary message='Ошибка заполнения, проверьте все ли поля заполнены корректно' />;
    }
  };

  const onSubmit = (data: TSignUp & { confirmPassword: string }) => {
    const { confirmPassword, ...signUpData } = data;
    mutate(signUpData);
  };

  return (
    <div className='flex h-screen flex-col items-center justify-center p-4'>
      <Form className='flex flex-col gap-4' onSubmit={handleSubmit(onSubmit)}>
        <Image image={registration} />
        <Typography as='h1' className='heading mb-3'>
          Добро пожаловать!
        </Typography>

        {step === 1 ? (
          <>
            <Input
              placeholder="Введите имя"
              {...register('firstName')}
              error={errors.firstName?.message}
              successMessage={getSuccessMessage(
                'firstName',
                'Имя заполнено корректно',
                touchedFields,
                errors,
                firstNameValue
              )}
            />

            <Input
              placeholder="Введите фамилию"
              {...register('lastName')}
              error={errors.lastName?.message}
              successMessage={getSuccessMessage(
                'lastName',
                'Фамилия заполнена корректно',
                touchedFields,
                errors,
                lastNameValue
              )}
            />

            <Button className='button button-primary' type='button' onClick={handleContinue} disabled={isSubmitting}>
              Продолжить регистрацию
            </Button>
          </>
        ) : (
          <>
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
                type='password'
                placeholder="Введите пароль"
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

              <Input
                type='password'
                placeholder="Введите пароль"
                {...register('confirmPassword')}
                error={errors.confirmPassword?.message}
                successMessage={getSuccessMessage(
                  'confirmPassword',
                  'Пароли совпадают',
                  touchedFields,
                  errors,
                  confirmPasswordValue
                )}
              />

              <Button
                className='button button-primary flex items-center justify-center'
                type='submit'
                disabled={!isValid || isPending}
              >
                {isPending ? <Loader color='white' size='sm' /> : <Typography as='p'>Зарегистрироваться</Typography>}
              </Button>
          </>
        )}
      </Form>

      <Link to='/login' className='mt-5'>
        <Typography as='p' className='inline'>
          У меня уже есть аккаунт. Войти
        </Typography>
      </Link>

      {isError && <ErrorBoundary message={error.name} />}
      {isSuccess && <ErrorBoundary message='Вы успешно зарегистрировались. Перейдите на страницу входа.' />}
    </div>
  );
};
