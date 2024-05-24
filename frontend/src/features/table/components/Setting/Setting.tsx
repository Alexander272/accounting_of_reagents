import { useRef, useState } from 'react'
import { Button, FormControlLabel, IconButton, Stack, Switch, Typography, useTheme } from '@mui/material'
import { Controller, useForm } from 'react-hook-form'

import type { ISelect } from '../../types/table'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { SettingIcon } from '@/components/Icons/SettingIcon'
import { TimesIcon } from '@/components/Icons/TimesIcon'
import { CheckIcon } from '@/components/Icons/CheckIcon'
import { Popover } from '@/components/Popover/Popover'
import { Columns } from '../../columns'
import { getHidden, setHidden } from '../../tableSlice'

export const Setting = () => {
	const [open, setOpen] = useState(false)
	const anchor = useRef(null)

	const { palette } = useTheme()

	const hidden = useAppSelector(getHidden)
	const dispatch = useAppDispatch()

	const methods = useForm<ISelect>({
		values: Columns.reduce((a, v) => ({ ...a, [v.key]: hidden[v.key] || false }), {}),
	})

	const toggleHandler = () => setOpen(prev => !prev)
	const closeHandler = () => toggleHandler()

	const resetHandler = () => {
		dispatch(setHidden())
		closeHandler()
	}
	const applyHandler = methods.handleSubmit(form => {
		const newHidden: ISelect = {}
		Object.keys(form).forEach(k => {
			if (form[k]) newHidden[k] = true
		})

		dispatch(setHidden(newHidden))
		closeHandler()
	})

	return (
		<>
			<IconButton ref={anchor} onClick={toggleHandler}>
				<SettingIcon fontSize={20} />
			</IconButton>

			<Popover open={open} onClose={closeHandler} anchorEl={anchor.current}>
				<Stack direction={'row'} mb={2.5} justifyContent={'space-between'} alignItems={'center'}>
					<Typography fontSize={'1.1rem'}>Видимые колонки</Typography>

					<Stack direction={'row'} spacing={1} height={34}>
						<Button onClick={resetHandler} variant='outlined' color='gray' sx={{ minWidth: 40 }}>
							<TimesIcon fill={palette.gray.main} fontSize={12} />
						</Button>

						<Button onClick={applyHandler} variant='contained' sx={{ minWidth: 40, padding: '6px 12px' }}>
							<CheckIcon fill={palette.common.white} fontSize={20} />
						</Button>
					</Stack>
				</Stack>

				<Stack spacing={1} maxHeight={450} overflow={'auto'}>
					{Columns.map(c => (
						<Controller
							key={c.key}
							control={methods.control}
							name={c.key}
							render={({ field }) => (
								<FormControlLabel
									label={c.label}
									sx={{
										color: !field.value ? 'inherit' : '#505050',
										transition: '.2s color ease-in-out',
									}}
									control={
										<Switch
											checked={!field.value}
											onChange={event => field.onChange(!event.target.checked)}
										/>
									}
								/>
							)}
						/>
					))}
				</Stack>
			</Popover>
		</>
	)
}
