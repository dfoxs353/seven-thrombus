import * as yup from 'yup';

export const registrationSchema = yup.object().shape({
  login: yup.string().min(6, 'Минимум 6 символов').required('Введите логин'),
  password: yup
    .string()
    .min(6, 'Минимум 6 символов')
    .required('Введите пароль'),
  confirmPassword: yup
    .string()
    .oneOf([yup.ref('password')], 'Пароли должны совпадать')
    .required('Повторите пароль')
});
