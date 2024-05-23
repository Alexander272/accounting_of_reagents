import { AppBar, Box, Stack, Toolbar } from '@mui/material'

import { useAppSelector } from '@/hooks/redux'
import { useSignOutMutation } from '@/features/auth/authApiSlice'
import { getToken } from '@/features/user/userSlice'
import { FlaskIcon } from '../Icons/FlaskIcon'
import { NavButton } from './header.style'
import logo from '@/assets/logo.webp'

export const LayoutHeader = () => {
	const [signOut] = useSignOutMutation()

	const token = useAppSelector(getToken)

	const signOutHandler = () => {
		void signOut(null)
	}

	return (
		<AppBar sx={{ borderRadius: 0 }}>
			<Toolbar
				sx={{
					// maxWidth: '1680px',
					// width: '100%',
					justifyContent: 'space-between',
					alignItems: 'inherit',
					// marginX: 'auto',
				}}
			>
				<Box alignSelf={'center'} display={'flex'} alignItems={'center'}>
					<img height={46} width={157} src={logo} alt='logo' />
					<FlaskIcon fill={'#042245'} />
				</Box>
				{token && (
					<Stack direction={'row'} spacing={3} minHeight={'100%'}>
						{/* <NavLink to='/'>Главная</NavLink> */}

						{/* {CheckPermission({ section: 'criterions', method: 'WRITE' }) && (
						<NavLink to={'criterions'}>Критерии</NavLink>
					)} */}

						<NavButton onClick={signOutHandler}>Выйти</NavButton>
					</Stack>
				)}
			</Toolbar>
		</AppBar>
	)
}
