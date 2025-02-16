import * as yup from 'yup';

export const loginSchema = yup.object().shape({
  username: yup.string().min(3, 'Минимум 3 символова').required('Введите логин'),
  password: yup.string().min(3, 'Минимум 3 символова').required('Введите пароль'),
});
