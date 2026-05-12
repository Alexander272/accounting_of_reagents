import { AppBar, Box, Stack, Toolbar, Tooltip, useTheme } from '@mui/material'
import { Link } from 'react-router-dom'
import logo from '@/assets/logo.webp'

import { AppRoutes } from '@/constants/routes'
import { Permissions } from '@/features/access/constants/permissions'
import { useAppSelector } from '@/hooks/redux'
import { useSignOutMutation } from '@/features/auth/authApiSlice'
import { getToken } from '@/features/user/userSlice'
import { ActiveRealm } from '@/features/realms/components/ActiveRealm'
import { Can } from '@/features/access/components/Can'
import { FlaskIcon } from '../Icons/FlaskIcon'
import { ShieldIcon } from '../Icons/ShieldIcon'
import { NavButton } from './header.style'

export const LayoutHeader = () => {
	const { palette } = useTheme()

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
					<Link to={AppRoutes.HOME}>
						<Stack
							display={'flex'}
							direction={'row'}
							height={50}
							overflow={'hidden'}
							alignItems={'center'}
							justifyContent={'center'}
							sx={{ img: { height: '100%', width: 'auto' } }}
						>
							<img height={46} width={157} src={logo} alt='logo' />
							<FlaskIcon fill={'#042245'} />
						</Stack>
					</Link>
				</Box>

				{token && (
					<Stack direction={'row'} spacing={3} minHeight={'100%'}>
						<Can
							anyOf={[
								Permissions.PermissionWrite,
								Permissions.UsersWrite,
								Permissions.RolesWrite,
								Permissions.RealmsWrite,
							]}
						>
							<Link to={AppRoutes.Accesses} aria-label='statistic page'>
								<Tooltip title='Настройки доступа' disableInteractive>
									<NavButton>
										<ShieldIcon
											sx={{
												fontSize: 26,
												transition: '0.3s all ease-in-out',
												stroke: palette.primary.main,
											}}
										/>
									</NavButton>
								</Tooltip>
							</Link>
						</Can>

						<NavButton>
							<ActiveRealm sx={{ mr: -0.5 }} />
						</NavButton>

						<NavButton onClick={signOutHandler}>Выйти</NavButton>
					</Stack>
				)}
			</Toolbar>
		</AppBar>
	)
}
