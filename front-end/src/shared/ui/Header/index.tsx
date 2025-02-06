import { Link } from '@tanstack/react-router';
import { HEADER } from '@/shared/constants/header';
import { Typography } from '../Typography';

export const Header = () => {
  return (
    <header className={'h-10 w-screen'}>
      <ul className="flex gap-2">
        {HEADER.navLinks.map((link) => (
          <li key={link.title + link.href}>
            <Link to={link.href}>
              <Typography>{link.title}</Typography>
            </Link>
          </li>
        ))}
      </ul>
    </header>
  );
};
