import { AppBar, Box, Divider, Stack, Toolbar, Tooltip, useMediaQuery, useTheme } from '@mui/material'
import { Link, NavLink } from 'react-router-dom'
import logo from '@/assets/logo.webp'

import { AppRoutes } from '@/constants/routes'
import { Permissions } from '@/features/access/constants/permissions'
import { useAppSelector } from '@/hooks/redux'
import { Can } from '@/features/access/components/Can'
import { useSignOutMutation } from '@/features/auth/authApiSlice'
import { getToken } from '@/features/user/userSlice'
import { ActiveRealm } from '@/features/realms/components/ActiveRealm'
import { FlaskIcon } from '../Icons/FlaskIcon'
import { ShieldIcon } from '../Icons/ShieldIcon'
import { NotebookIcon } from '../Icons/NotebookIcon'
import { LogoutIcon } from '../Icons/LogoutIcon'
import { NavButton } from './header.style'

export const LayoutHeader = () => {
	const { palette, breakpoints } = useTheme()

	const showLabels = useMediaQuery(breakpoints.up('md'))

	const [signOut] = useSignOutMutation()

	const token = useAppSelector(getToken)

	const signOutHandler = () => {
		void signOut(null)
	}

	return (
		<AppBar sx={{ borderRadius: 0 }}>
			<Toolbar sx={{ justifyContent: 'space-between', alignItems: 'inherit' }}>
				<Box alignSelf={'center'} display={'flex'} alignItems={'center'}>
					<Link to={AppRoutes.Home}>
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
							<FlaskIcon fill={palette.primary.dark} />
						</Stack>
					</Link>

					<ActiveRealm />
				</Box>

				{token && (
					<Stack direction='row' alignItems='center' spacing={1} minHeight='100%'>
						<Can
							anyOf={[
								Permissions.PermissionWrite,
								Permissions.UsersWrite,
								Permissions.RolesWrite,
								Permissions.RealmsWrite,
							]}
						>
							<Tooltip title='Настройки доступа' disableInteractive>
								<NavButton component={NavLink} to={AppRoutes.Accesses}>
									<ShieldIcon sx={{ fontSize: 20, fill: 'currentColor', mr: 1 }} />
									{showLabels && 'Доступ'}
								</NavButton>
							</Tooltip>
						</Can>

						<Can anyOf={[Permissions.SpendingRead]}>
							<Tooltip title='Отчёт по расходу за период' disableInteractive>
								<NavButton component={NavLink} to={AppRoutes.Report}>
									<NotebookIcon fontSize={20} fill='currentcolor' mr={1} />
									{showLabels && 'Отчёты'}
								</NavButton>
							</Tooltip>
						</Can>

						<Divider orientation='vertical' flexItem sx={{ my: '12px !important' }} />

						<Tooltip title='Выйти из аккаунта' disableInteractive>
							<NavButton onClick={signOutHandler} sx={{ color: 'text.secondary' }}>
								<LogoutIcon fontSize={20} fill='currentcolor' mr={1} />
								{showLabels && 'Выйти'}
							</NavButton>
						</Tooltip>
					</Stack>
				)}
			</Toolbar>
		</AppBar>
	)
}
