import { ROUTES } from './routes';
import home from '@/assets/images/icons/home.svg'
import disciplines from '@/assets/images/icons/disciplines.svg'
import schedule from '@/assets/images/icons/schedule.svg'
import users from '@/assets/images/icons/users.svg'

export const HEADER = {
  navLinks: [
    {
      title: 'Главная',
      href: ROUTES.HOME,
      icon: home
    },
    {
      title: 'Дисциплины',
      href: ROUTES.DISCIPLINES,
      icon: disciplines
    },
    {
      title: 'Пользователи',
      href: ROUTES.USERS,
      icon: users
    },
    {
      title: 'Расписание',
      href: ROUTES.SCHEDULE,
      icon: schedule
    },
  ],
};
