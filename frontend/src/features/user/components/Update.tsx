import { useEffect, type FC } from 'react'
import {
	Dialog,
	DialogTitle,
	DialogContent,
	DialogActions,
	Button,
	TextField,
	MenuItem,
	IconButton,
	Typography,
	Select,
	Stack,
} from '@mui/material'
import { Controller, useFieldArray, useForm } from 'react-hook-form'
import { toast } from 'react-toastify'

import type { IFetchError } from '@/app/types/error'
import type { IUserData, IUserDataDTO } from '@/features/user/types/user'
import { useGetRolesQuery } from '@/features/user/roleApiSlice'
import { useUpdateUserMutation } from '../userApiSlice'
import { TimesIcon } from '@/components/Icons/TimesIcon'
import { Fallback } from '@/components/Fallback/Fallback'
import { BoxFallback } from '@/components/Fallback/BoxFallback'
import { useGetAllRealmsQuery } from '@/features/realms/realmsApiSlice'

type Props = {
	user: IUserData
	onClose: () => void
}

export const UpdateModal: FC<Props> = ({ user, onClose }) => {
	const { data: roles, isFetching: isFetchingRoles } = useGetRolesQuery()
	const { data: realms, isFetching: isFetchingRealms } = useGetAllRealmsQuery()

	const { control, handleSubmit, watch, reset } = useForm<IUserDataDTO>({
		defaultValues: {
			...user,
			realms: [],
		},
	})
	const { fields } = useFieldArray({ control, name: 'realms' })

	console.log('user', user)

	const [update, { isLoading }] = useUpdateUserMutation()

	useEffect(() => {
		if (realms?.data && user) {
			const userRealmsMap = new Map(user.realms.map(r => [r.realmId, r]))

			const initialRealms = realms.data.map(realm => {
				const userBinding = userRealmsMap.get(realm.id)
				return {
					id: userBinding?.id || '',
					userId: user.id,
					realmId: realm.id,
					realm: realm,
					roleId: userBinding?.roleId || '',
					role: userBinding?.role,
					isActive: userBinding?.isActive ?? true,
					createdAt: userBinding?.createdAt || '',
				}
			})

			reset({ ...user, realms: initialRealms })
		}
	}, [realms, user, reset])

	const saveHandler = handleSubmit(async form => {
		console.log('save user', form)
		const payload = {
			...form,
			realms: form.realms.filter(r => r.roleId || r.createdAt),
		}

		try {
			await update(payload).unwrap()
			toast.success('Пользователь обновлен')
			onClose()
		} catch (error) {
			const fetchError = error as IFetchError
			toast.error(fetchError.data.message, { autoClose: false })
		}
	})

	if (isFetchingRoles || isFetchingRealms) return <Fallback />
	return (
		<Dialog
			open={Boolean(user)}
			onClose={onClose}
			fullWidth
			maxWidth='sm'
			PaperProps={{
				sx: {
					borderRadius: '16px',
					p: 1,
				},
			}}
		>
			<DialogTitle sx={{ m: 0, p: 2, display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
				<Typography variant='h6' component='div' sx={{ fontWeight: 'bold' }}>
					Добавить пользователя
				</Typography>
				<IconButton onClick={onClose} sx={{ color: 'text.secondary' }}>
					<TimesIcon fontSize={16} />
				</IconButton>
			</DialogTitle>

			{isLoading && <BoxFallback />}

			{/* Body */}
			<DialogContent dividers sx={{ borderTop: '1px solid #f0f0f0', borderBottom: '1px solid #f0f0f0', py: 3 }}>
				<Stack direction={'row'} spacing={2} mb={2} sx={{ width: '100%' }}>
					<Stack sx={{ width: '50%' }}>
						<Typography variant='caption' sx={{ fontWeight: 600, mb: 1, display: 'block' }}>
							Имя
						</Typography>
						<Controller
							control={control}
							name='firstName'
							disabled
							render={({ field }) => <TextField {...field} fullWidth />}
						/>
					</Stack>

					<Stack sx={{ width: '50%' }}>
						<Typography variant='caption' sx={{ fontWeight: 600, mb: 1, display: 'block' }}>
							Фамилия
						</Typography>
						<Controller
							control={control}
							name='lastName'
							disabled
							render={({ field }) => <TextField {...field} fullWidth />}
						/>
					</Stack>
				</Stack>

				<Stack mb={2}>
					<Typography variant='caption' sx={{ fontWeight: 600, mb: 1, display: 'block' }}>
						Email
					</Typography>
					<Controller
						control={control}
						name='email'
						disabled
						render={({ field }) => <TextField {...field} fullWidth />}
					/>
				</Stack>

				{fields.map((field, index) => (
					<Stack key={field.id} direction={'row'} spacing={2} sx={{ width: '100%', mb: 2 }}>
						<Stack width={'25%'}>
							<Typography variant='caption' sx={{ fontWeight: 600, mb: 1, display: 'block' }}>
								Область
							</Typography>
							<Typography pt={'9px'}>{field.realm?.name}</Typography>
						</Stack>

						<Stack sx={{ flex: 1 }}>
							<Typography variant='caption' sx={{ fontWeight: 600, mb: 1, display: 'block' }}>
								Роль
							</Typography>
							<Controller
								control={control}
								name={`realms.${index}.roleId`}
								render={({ field: selectField }) => (
									<Select {...selectField} fullWidth displayEmpty>
										<MenuItem value=''>
											<em>Нет роли</em>
										</MenuItem>
										{roles?.data
											.filter(role => role.isEditable)
											.map(role => (
												<MenuItem key={role.id} value={role.id}>
													{role.name}
												</MenuItem>
											))}
									</Select>
								)}
							/>
						</Stack>

						{/* Статус показываем только если роль выбрана */}
						<Stack sx={{ flex: 1 }}>
							<Typography variant='caption' sx={{ fontWeight: 600, mb: 1, display: 'block' }}>
								Статус
							</Typography>
							<Controller
								control={control}
								name={`realms.${index}.isActive`}
								render={({ field: statusField }) => (
									<Select
										value={statusField.value ? 'active' : 'inactive'}
										onChange={e => statusField.onChange(e.target.value === 'active')}
										fullWidth
										disabled={!watch(`realms.${index}.roleId`)} // Блокируем, если роль не выбрана
									>
										<MenuItem value={'active'}>Активный</MenuItem>
										<MenuItem value={'inactive'}>Неактивный</MenuItem>
									</Select>
								)}
							/>
						</Stack>
					</Stack>
				))}
			</DialogContent>

			{/* Footer */}
			<DialogActions sx={{ p: 2, gap: 1 }}>
				<Button
					onClick={onClose}
					variant='outlined'
					sx={{ borderRadius: '8px', textTransform: 'none', color: 'text.primary', borderColor: '#ddd' }}
				>
					Отмена
				</Button>
				<Button
					onClick={saveHandler}
					variant='contained'
					sx={{ borderRadius: '8px', textTransform: 'none', px: 3 }}
				>
					Сохранить
				</Button>
			</DialogActions>
		</Dialog>
	)
}
