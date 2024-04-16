import { Link } from 'react-router-dom'
import { Button, Divider, Stack, Typography } from '@mui/material'

import { Container } from './notFound.style'
import logo from '@/assets/logo192.webp'

export default function NotFound() {
	return (
		<Container>
			<Stack
				direction={'row'}
				divider={<Divider orientation='vertical' flexItem />}
				spacing={4}
				alignItems={'center'}
			>
				<img width={148} height={148} src={logo} alt='logo' />
				<Typography variant='h4' sx={{ fontSize: '8rem', fontWeight: 'bold', color: '#8fb2cb' }}>
					404
				</Typography>
			</Stack>

			<Typography mt={3} mb={3} sx={{ fontSize: '2rem', color: '#8fb2cb' }}>
				Страница не найдена
			</Typography>

			<Link to='/'>
				<Button variant='outlined' size='large' sx={{ borderRadius: '12px', padding: '8px 32px' }}>
					На главную
				</Button>
			</Link>
		</Container>
	)
}
