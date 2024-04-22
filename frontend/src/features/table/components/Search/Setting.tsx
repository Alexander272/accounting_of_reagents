import { useRef, useState } from 'react'
import { Checkbox, FormControlLabel, IconButton, Stack, useTheme } from '@mui/material'
import { Controller, FormProvider, useForm } from 'react-hook-form'

import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { SettingIcon } from '@/components/Icons/SettingIcon'
import { Popover } from '@/components/Popover/Popover'
import { Columns } from '../../columns'
import { getSearch, setSearchFields } from '../../tableSlice'

type FieldList = {
	[x: string]: boolean
}

export const Setting = () => {
	const [open, setOpen] = useState(false)
	const anchor = useRef(null)

	const { palette } = useTheme()

	const { fields } = useAppSelector(getSearch)
	const dispatch = useAppDispatch()

	const methods = useForm<FieldList>({ defaultValues: fields?.reduce((a, v) => ({ ...a, [v]: true }), {}) })

	const toggleHandler = () => setOpen(prev => !prev)

	const mouseDownSettingHandler = (event: React.MouseEvent<HTMLButtonElement>) => {
		event.preventDefault()
	}

	const saveHandler = (data: FieldList) => {
		console.log('submit', data)
		const values = Object.keys(data).reduce((ac: string[], k) => {
			if (data[k]) ac.push(k)
			return ac
		}, [])
		dispatch(setSearchFields(values))
		toggleHandler()
	}

	const closeHandler = () => {
		methods.handleSubmit(saveHandler)()
	}

	//TODO надо перед циклом отфильтровать список колонок (чтобы убрать из выбора все скрытые и все типы полей отличные от строк и чисел)
	// можно наверное сделать в state массив со списком скрытых полей

	return (
		<>
			<IconButton ref={anchor} onClick={toggleHandler} onMouseDown={mouseDownSettingHandler} edge='end'>
				<SettingIcon fontSize={20} />
			</IconButton>

			<Popover open={open} onClose={closeHandler} anchorEl={anchor.current}>
				<FormProvider {...methods}>
					<Stack spacing={1} sx={{ maxHeight: 300, overflow: 'auto', userSelect: 'none', pr: 1 }}>
						{Columns.map(c => (
							<Controller
								key={c.key}
								control={methods.control}
								name={c.key}
								render={({ field }) => (
									<FormControlLabel
										label={c.label}
										control={<Checkbox checked={field.value || false} />}
										onChange={field.onChange}
										sx={{
											transition: 'all 0.3s ease-in-out',
											borderRadius: 3,
											':hover': { backgroundColor: palette.action.hover },
										}}
									/>
								)}
							/>
						))}
					</Stack>
				</FormProvider>
			</Popover>
		</>
	)
}
