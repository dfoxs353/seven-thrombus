import { Link, useLocation } from '@tanstack/react-router';
import { HEADER } from '@/shared/constants/header';
import { Typography } from '@shared/ui';
import { ReactSVG } from 'react-svg';

export const Header = () => {
  const pathname = useLocation({
    select: (location) => location.pathname
  })
  return (
    <header className={'header'}>
      <ul className='flex justify-center gap-2'>
        {HEADER.navLinks.map((link) => (
          <li className='shrink group' key={link.title + link.href}>
            <Link className='flex flex-col items-center justify-center gap-0.5 h-[100%] px-3 py-2.5' to={link.href}>
              <ReactSVG className={`${pathname === link.href ? 'stroke-[var(--main-accent)]' : 'stroke-[var(--additional-dark)]'} transition-colors duration-300 group-hover:stroke-[var(--main-accent)]`} src={link.icon}></ReactSVG>
              <Typography className={`${pathname === link.href ? 'text-[var(--main-accent)]' : 'text-[var(--additional-dark)]'} transition-colors duration-300 text-header group-hover:text-[var(--main-accent)]`}>
                {link.title}
              </Typography>
            </Link>
          </li>
        ))}
      </ul>
    </header>
  );
};
