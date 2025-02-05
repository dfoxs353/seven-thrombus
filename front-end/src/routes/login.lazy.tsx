import { createLazyFileRoute, Link } from '@tanstack/react-router'
import { Form, Typography, Image, Button } from '../shared/ui'
import registration from '@/assets/images/regisration.webp'
import { Input } from '../shared/ui/Input'
import { yupResolver } from "@hookform/resolvers/yup";
import { useForm } from 'react-hook-form';
import { loginSchema } from '../shared/schema/login';

export const Route = createLazyFileRoute('/login')({
  component: RouteComponent,
})

function RouteComponent() {

  const {
    register,
    handleSubmit,
    watch,
    formState: { errors, isValid, touchedFields },
  } = useForm({
    resolver: yupResolver(loginSchema),
    mode: "onChange",
  });

  const loginValue = watch("login");
  const passwordValue = watch("password");

  const onSubmit = (data: any) => {
    console.log("Форма отправлена", data);
  };

  return (
    <div className="p-4 h-screen flex flex-col items-center justify-center">
      <Form classname='flex flex-col gap-4' onSubmit={handleSubmit(onSubmit)}>
        <Image image={registration}></Image>
        <Typography as='h1' className="heading mb-3">Добро пожаловать!</Typography>
        <Input
          placeholder="Введите логин"
          className={`input w-full ${touchedFields.login && loginValue
            ? errors.login
              ? "border-[var(--error)]"
              : "border-[var(--approve)]"
            : "border-gray-300"
            }`}
          {...register("login")}
        />
        {errors.login && <Typography className="text-[var(--error)] absolute top-[2px] left-[10px] text-[12px]">{errors.login.message}</Typography>}

        <div className='w-full relative'>
          <Input
            type="password"
            placeholder="Введите пароль"
            className={`input w-full ${touchedFields.password && passwordValue
              ? errors.password && passwordValue
                ? "border-[var(--error)]"
                : "border-[var(--approve)]"
              : "border-gray-300"
              }`}
            {...register("password")}
          />
          {errors.password && <Typography className="text-[var(--error)] absolute top-[2px] left-[10px] text-[12px]">{errors.password.message}</Typography>}
        </div>

        <Button className='button button-primary' type="submit" disabled={!isValid}>
          <Typography as='p'>Войти</Typography>
        </Button>
      </Form>
      <Link to='/registration' className='mt-5'><Typography as='p' className='inline'>У меня нет аккаунта. Зарегистрироваться</Typography></Link>
    </div>
  )
}
