import { createFileRoute, Link } from '@tanstack/react-router'
import { Form, Typography, Image } from '../shared/ui'
import teamwork from '@/assets/images/team-work.webp'

export const Route = createFileRoute('/')({
  component: Index,
})

function Index() {

  return (
    <div className="p-4 h-screen flex flex-col items-center justify-center">
      <Form>
        <Image image={teamwork}></Image>
        <Typography as='h1' className="heading mb-3">Создайте будущее вместе</Typography>
        <Link to='/login' className='button button-primary'>
          <Typography as='p'>Войти</Typography>
        </Link>
        <Link to='/registration' className='button button-secondary '>
          <Typography as='p'>Зарегистрироваться</Typography>
        </Link>
      </Form>
    </div>
  )
}
