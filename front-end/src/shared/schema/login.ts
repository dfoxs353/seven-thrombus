import * as yup from 'yup';

export const loginSchema = yup.object().shape({
  login: yup.string().min(6, 'Минимум 6 символов').required('Введите логин'),
  password: yup.string().min(6, 'Минимум 6 символов').required('Введите пароль')
});
