import { FC, useEffect } from 'react'
import {
	alpha,
	Divider,
	FormControl,
	MenuItem,
	Select,
	type SelectChangeEvent,
	type SxProps,
	Theme,
	Tooltip,
} from '@mui/material'

import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { getUserRealms, setRole } from '@/features/user/userSlice'
import { setPage } from '@/features/table/tableSlice'
import { getRealm, setRealm } from '../realmSlice'
import { DashboardIcon } from '@/components/Icons/DashboardIcon'

type Props = {
	sx?: SxProps<Theme>
}

export const ActiveRealm: FC<Props> = ({ sx }) => {
	const realm = useAppSelector(getRealm)
	const dispatch = useAppDispatch()

	const realms = useAppSelector(getUserRealms)

	useEffect(() => {
		if (!realms) return
		const founded = realms.find(e => e.realmId === realm?.id)
		if (founded) return
		if (realms.length && realms[0].realm) dispatch(setRealm(realms[0].realm))
	}, [realms, dispatch, realm])

	const changeHandler = async (event: SelectChangeEvent) => {
		const value = realms.find(e => e.realmId === event.target.value)
		if (!value) return

		dispatch(setRealm(value.realm!))
		dispatch(setRole(value.role?.name || ''))
		dispatch(setPage(1))
	}

	if ((realms.length || 0) < 2) return null
	return (
		<>
			<Divider orientation='vertical' flexItem sx={{ my: 1.5, mx: 2 }} />

			<Tooltip title='Рабочее пространство' disableInteractive enterDelay={500}>
				<FormControl size='small' sx={sx}>
					<Select
						value={realm?.id || ''}
						onChange={changeHandler}
						displayEmpty
						startAdornment={<DashboardIcon sx={{ fontSize: 18, color: 'text.secondary' }} />}
						sx={{
							height: 38,
							fontSize: 15,
							fontWeight: 500,
							color: 'text.primary',
							borderRadius: 2,
							bgcolor: theme => alpha(theme.palette.primary.main, 0.06),
							'.MuiOutlinedInput-notchedOutline': { border: 0 },
							'&:hover': { bgcolor: theme => alpha(theme.palette.primary.main, 0.1) },
							'.MuiSelect-select': { py: 0, pl: 1.5 },
						}}
					>
						<MenuItem value='' disabled>
							Выберите область
						</MenuItem>
						{realms.map(item => (
							<MenuItem key={item.id} value={item.realmId}>
								{item.realm?.name || 'No name'}
							</MenuItem>
						))}
					</Select>
				</FormControl>
			</Tooltip>
		</>
	)
}
