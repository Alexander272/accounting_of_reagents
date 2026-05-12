import { useMemo, useState, type FC } from 'react'
import {
	Box,
	Typography,
	Button,
	TextField,
	MenuItem,
	Select,
	Table,
	TableBody,
	TableCell,
	TableContainer,
	TableHead,
	TableRow,
	Paper,
	InputAdornment,
	type SelectChangeEvent,
	useTheme,
	Avatar,
	Chip,
	Tooltip,
	CircularProgress,
} from '@mui/material'
import dayjs from 'dayjs'

import type { IUserData, IUserRealm } from '@/features/user/types/user'
import { getAvatarColor, getInitials } from './utils'
import { stringToHSLA } from '@/utils/colors'
import { useDebounce } from '@/hooks/useDebounce'
import { useGetAllUsersQuery, useSyncUsersMutation } from '@/features/user/userApiSlice'
import { useGetRolesQuery } from '@/features/user/roleApiSlice'
import { BoxFallback } from '@/components/Fallback/BoxFallback'
import { SearchIcon } from '@/components/Icons/SearchIcon'
import { SyncIcon } from '@/components/Icons/SyncIcon'
import { ModifyIcon } from '@/components/Icons/ModifyIcon'
import { StatusBadge } from '../StatusBadge'
import { UpdateModal } from '../../../user/components/Update'

export const Users = () => {
	const { palette } = useTheme()
	const [search, setSearch] = useState('')
	const [roleFilter, setRoleFilter] = useState<string[]>([''])
	const [statusFilter, setStatusFilter] = useState('')

	const [user, setUser] = useState<IUserData | null>(null)

	const debouncedSearch = useDebounce(search, 300)

	const { data, isFetching } = useGetAllUsersQuery()
	const { data: roles, isFetching: isFetchingRoles } = useGetRolesQuery()
	const [sync, { isLoading }] = useSyncUsersMutation()

	const filteredUsers = useMemo(() => {
		if (!data?.data) return []

		const lowSearch = (debouncedSearch as string).toLowerCase().trim()

		return data.data.filter(user => {
			// 1. Поиск (по имени, фамилии или почте)
			const matchesSearch =
				!lowSearch ||
				[user.lastName, user.firstName, user.email].some(field => field?.toLowerCase().includes(lowSearch))

			// 2. Фильтр по ролям
			// Проверяем, есть ли у пользователя хотя бы один realm с выбранной ролью
			const matchesRole =
				roleFilter.length === 0 || roleFilter.includes('')
					? true
					: user.realms?.some(realm => roleFilter.includes(realm.role?.name || ''))

			// 3. Фильтр по статусу
			const matchesStatus =
				!statusFilter || statusFilter === ''
					? true
					: statusFilter === 'active'
						? user.realms?.some(realm => realm.isActive) // Есть хотя бы один активный реалм
						: user.realms?.every(realm => !realm.isActive) // Все реалмы неактивны (или массив пуст)

			return matchesSearch && matchesRole && matchesStatus
		})
	}, [data, debouncedSearch, roleFilter, statusFilter])

	const syncHandler = async () => {
		await sync()
	}

	const roleHandler = (event: SelectChangeEvent<string[]>) => {
		const value = event.target.value
		let newValue = typeof value === 'string' ? value.split(',') : value

		// 1. Если список стал пустым — возвращаем ''
		if (newValue.length === 0) {
			newValue = ['']
		}
		// 2. Если в списке больше одного элемента
		else if (newValue.length > 1) {
			// Если только что добавили что-то к '', то убираем ''
			if (newValue.includes('')) {
				newValue = newValue.filter(v => v !== '')
			}

			// 3. Если выбраны все доступные опции (кроме ''),
			// тут можно добавить условие сравнения с длиной исходного массива ролей
			if (newValue.length === roles?.data.length) {
				newValue = ['']
			}
		}

		setRoleFilter(newValue)
	}

	const userHandler = (user: IUserData | null) => {
		setUser(user)
	}

	return (
		<Box sx={{ p: 3 }}>
			{/* Page Header */}
			<Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', mb: 4 }}>
				<Box>
					<Typography variant='h4' sx={{ fontWeight: 'bold' }}>
						Пользователи
					</Typography>
					<Typography variant='body1' color='text.secondary'>
						Управление учётными записями
					</Typography>
				</Box>
				<Button
					variant='outlined'
					sx={{ borderRadius: '8px', textTransform: 'none', background: '#fff' }}
					onClick={syncHandler}
				>
					{isLoading ? (
						<CircularProgress size={16} sx={{ mr: 1.5 }} />
					) : (
						<SyncIcon fill={palette.primary.main} fontSize={16} mr={1.5} />
					)}
					Синхронизировать
				</Button>
			</Box>

			{isFetching || isFetchingRoles || isLoading ? <BoxFallback /> : null}

			{/* Toolbar */}
			<Box sx={{ display: 'flex', gap: 2, mb: 3, flexWrap: 'wrap' }}>
				<TextField
					placeholder='Поиск по имени или email…'
					size='small'
					value={search}
					onChange={e => setSearch(e.target.value)}
					InputProps={{
						startAdornment: (
							<InputAdornment position='start'>
								<SearchIcon fontSize='small' />
							</InputAdornment>
						),
					}}
					sx={{ flexGrow: 1, minWidth: '200px', background: '#fff' }}
				/>

				<Select
					size='small'
					displayEmpty
					multiple
					value={roleFilter}
					onChange={roleHandler}
					sx={{ width: '400px', background: '#fff' }}
				>
					<MenuItem value='' disabled>
						Все роли
					</MenuItem>
					{roles?.data.map(role => (
						<MenuItem key={role.id} value={role.name}>
							{role.name}
						</MenuItem>
					))}
				</Select>

				<Select
					size='small'
					displayEmpty
					value={statusFilter}
					onChange={e => setStatusFilter(e.target.value)}
					sx={{ width: '300px', background: '#fff' }}
				>
					<MenuItem value=''>Все статусы</MenuItem>
					<MenuItem value='active'>Активные</MenuItem>
					<MenuItem value='inactive'>Неактивные</MenuItem>
				</Select>
			</Box>

			{user ? <UpdateModal user={user} onClose={() => setUser(null)} /> : null}

			{/* Table Container */}
			<TableContainer component={Paper} elevation={0} sx={{ border: '1px solid #eee', borderRadius: 2 }}>
				<Table>
					<TableHead>
						<TableRow sx={{ borderBottom: '1px solid #f3f4f6' }}>
							<TableCell sx={{ color: 'text.secondary', fontSize: '0.875rem' }}>Пользователь</TableCell>
							<TableCell
								align='center'
								sx={{ color: 'text.secondary', fontSize: '0.875rem', width: 200 }}
							>
								Область
							</TableCell>
							<TableCell
								align='center'
								sx={{ color: 'text.secondary', fontSize: '0.875rem', width: 250 }}
							>
								Роль
							</TableCell>
							<TableCell
								align='center'
								sx={{ color: 'text.secondary', fontSize: '0.875rem', width: 200 }}
							>
								Статус
							</TableCell>
							<TableCell
								align='center'
								sx={{ color: 'text.secondary', fontSize: '0.875rem', width: 250 }}
							>
								Создан
							</TableCell>
							<TableCell sx={{ color: 'text.secondary', fontSize: '0.875rem', width: 100 }}>
								Действия
							</TableCell>
						</TableRow>
					</TableHead>
					<TableBody>
						{filteredUsers.map(user => (
							<UserRow key={user.id} u={user} setUser={userHandler} />
						))}
						{!data?.data.length && !isFetching ? (
							<TableRow>
								<TableCell colSpan={6} align='center' sx={{ py: 3, color: 'text.secondary' }}>
									Пользователи не найдены.
								</TableCell>
							</TableRow>
						) : null}
					</TableBody>
				</Table>
			</TableContainer>
		</Box>
	)
}

type RowProps = {
	u: IUserData
	setUser: (user: IUserData | null) => void
}

const UserRow: FC<RowProps> = ({ u, setUser }) => {
	const { palette } = useTheme()

	const editHandler = (e: React.MouseEvent) => {
		e.stopPropagation()
		e.preventDefault()

		setUser(u)
	}

	return (
		<>
			<TableRow key={u.id} hover>
				{/* Пользователь */}
				<TableCell rowSpan={u.realms.length}>
					<Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
						<Avatar
							sx={{
								bgcolor: getAvatarColor(u.id),
								fontSize: '14px',
								width: 36,
								height: 36,
							}}
						>
							{getInitials(u)}
						</Avatar>
						<Box>
							<Typography variant='body2' sx={{ fontWeight: 500 }}>
								{u.firstName && u.lastName ? u.firstName + ' ' + u.lastName : u.username}
							</Typography>
							<Typography variant='caption' color='text.secondary' sx={{ display: 'block' }}>
								{u.email}
							</Typography>
						</Box>
					</Box>
				</TableCell>

				{u.realms.length > 0 && (
					<>
						<RealmRow r={u.realms[0]} />

						{/* Действия */}
						<TableCell rowSpan={u.realms.length} align='center' sx={{ p: 0 }}>
							<Tooltip title='Редактировать пользователя'>
								<Button
									onClick={editHandler}
									sx={{
										minWidth: 60,
										minHeight: 60,
										borderRadius: '6px',
										':hover': { svg: { fill: palette.secondary.main } },
									}}
								>
									<ModifyIcon sx={{ fontSize: 18 }} />
								</Button>
							</Tooltip>
						</TableCell>
					</>
				)}
			</TableRow>
			{u.realms.length > 1
				? [...u.realms].slice(1).map(r => (
						<TableRow key={r.realmId} hover>
							<RealmRow r={r} />
						</TableRow>
					))
				: null}
		</>
	)
}

type RealmRowProps = {
	r: IUserRealm
}
const RealmRow: FC<RealmRowProps> = ({ r }) => {
	const colors = useMemo(() => stringToHSLA(r.role?.slug || ''), [r.role?.slug])
	const realmColors = useMemo(() => stringToHSLA(r.realm?.name || ''), [r.realm?.name])

	if (!r.role)
		return (
			<>
				<TableCell colSpan={4}></TableCell>
			</>
		)

	return (
		<>
			{/* Область */}
			<TableCell align='center'>
				<Chip
					label={r.realm?.name}
					size={'small'}
					style={{
						backgroundColor: realmColors.bg,
						color: realmColors.text,
						border: `1px solid ${realmColors.border}`,
						fontWeight: 500,
						fontSize: '0.75rem',
						height: '20px',
						borderRadius: '6px',
					}}
				/>
			</TableCell>
			{/* Роль */}
			<TableCell align='center'>
				<Chip
					label={r.role?.name}
					size={'small'}
					style={{
						backgroundColor: colors.bg,
						color: colors.text,
						border: `1px solid ${colors.border}`,
						fontWeight: 500,
						fontSize: '0.75rem',
						height: '20px',
						borderRadius: '6px',
					}}
				/>
			</TableCell>

			{/* Статус */}
			<TableCell align='center'>
				<StatusBadge active={r.isActive} label={r.isActive ? 'Активный' : 'Неактивный'} />
			</TableCell>

			{/* Создан */}
			<TableCell align='center' sx={{ color: 'text.secondary', fontSize: '13px' }}>
				{dayjs(r.createdAt).format('dddd, DD MMM YYYY HH:mm')}
			</TableCell>
		</>
	)
}
