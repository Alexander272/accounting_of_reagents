import { FC, useRef, useState } from 'react'
import {
	Button,
	FormControl,
	IconButton,
	InputLabel,
	MenuItem,
	Select,
	SelectChangeEvent,
	Stack,
	TextField,
	Typography,
	useTheme,
} from '@mui/material'
import { Controller, FormProvider, useFieldArray, useForm, useFormContext } from 'react-hook-form'
import { DatePicker } from '@mui/x-date-pickers'
import dayjs from 'dayjs'

import type { CompareTypes, IFilter } from '../../types/data'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { Popover } from '@/components/Popover/Popover'
import { Badge } from '@/components/Badge/Badge'
import { FilterIcon } from '@/components/Icons/FilterIcon'
import { TimesIcon } from '@/components/Icons/TimesIcon'
import { PlusIcon } from '@/components/Icons/PlusIcon'
import { CheckIcon } from '@/components/Icons/CheckIcon'
import { getFilters, setFilters } from '../../tableSlice'
import { Columns } from '../../columns'

const columns = Columns.filter(c => c.filter)

const defaultValue: IFilter = {
	field: `${columns[0].key}@${columns[0].filter}`,
	fieldType: columns[0].filter as 'string',
	compareType: 'con',
	value: '',
}

export const Filters = () => {
	const [open, setOpen] = useState(false)
	const anchor = useRef(null)

	const filters = useAppSelector(getFilters)
	const dispatch = useAppDispatch()

	const { palette } = useTheme()

	const methods = useForm<{ filters: IFilter[] }>({ values: { filters: [defaultValue] } })
	const { fields, append, remove } = useFieldArray({ control: methods.control, name: 'filters' })

	const toggleHandler = () => setOpen(prev => !prev)

	const closeHandler = () => {
		toggleHandler()
	}

	const addNewHandler = () => {
		append(defaultValue)
	}
	const removeHandler = (index: number) => {
		remove(index)
	}

	const resetHandler = () => {
		dispatch(setFilters([]))
		closeHandler()
	}
	const applyHandler = methods.handleSubmit(form => {
		// const result = Object.groupBy(form.filters, ({ field }) => field)
		console.log('form', form)

		const groupedMap = new Map<string, IFilter[]>()
		for (const e of form.filters) {
			e.field = e.field.split('@')[0]
			let thisList = groupedMap.get(e.field)
			if (thisList === undefined) {
				thisList = []
				groupedMap.set(e.field, thisList)
			}
			thisList.push(e)
		}
		const filters: IFilter[] = []
		groupedMap.forEach(v => filters.push(...v))

		// const filters = form.filters.sort((a, b) => a.field.localeCompare(b.field))
		// console.log(filters)

		dispatch(setFilters(filters))
		closeHandler()
	})

	return (
		<>
			<Button
				ref={anchor}
				onClick={toggleHandler}
				variant='outlined'
				color='gray'
				sx={{ minWidth: 30, paddingX: 1.5 }}
			>
				<Badge color='primary' variant={filters.length < 2 ? 'dot' : 'standard'} badgeContent={filters.length}>
					<FilterIcon fill={palette.gray.main} fontSize={16} mr={1} />
				</Badge>
				Фильтр
			</Button>

			<Popover
				open={open}
				onClose={closeHandler}
				anchorEl={anchor.current}
				paperSx={{
					maxWidth: 700,
					'&:before': {
						content: '""',
						display: 'block',
						position: 'absolute',
						top: 0,
						left: '66%',
						width: 10,
						height: 10,
						bgcolor: 'background.paper',
						transform: 'translate(-50%, -50%) rotate(45deg)',
						zIndex: 0,
					},
				}}
			>
				<Stack direction={'row'} mb={2.5} justifyContent={'space-between'} alignItems={'center'}>
					<Typography fontSize={'1.1rem'}>Применить фильтр</Typography>

					<Stack direction={'row'} spacing={1} height={34}>
						<Button onClick={addNewHandler} variant='outlined' sx={{ minWidth: 40, padding: '5px 14px' }}>
							<PlusIcon fill={palette.primary.main} fontSize={14} />
						</Button>

						<Button onClick={resetHandler} variant='outlined' color='gray' sx={{ minWidth: 40 }}>
							<TimesIcon fill={palette.gray.main} fontSize={12} />
						</Button>

						<Button onClick={applyHandler} variant='contained' sx={{ minWidth: 40, padding: '6px 12px' }}>
							<CheckIcon fill={palette.common.white} fontSize={20} />
						</Button>
					</Stack>
				</Stack>

				<Stack spacing={1.5}>
					<FormProvider {...methods}>
						{fields.map((f, i) => (
							<FilterItem key={f.id} index={i} remove={removeHandler} />
						))}
					</FormProvider>
				</Stack>
			</Popover>
		</>
	)
}

const compareTypes = new Map([
	['string', 'con'],
	['date', 'eq'],
	['number', 'eq'],
])

type FilterItemProps = {
	index: number
	remove: (index: number) => void
}
const FilterItem: FC<FilterItemProps> = ({ index, remove }) => {
	const methods = useFormContext<{ filters: IFilter[] }>()
	const type = methods.watch(`filters.${index}.fieldType`)

	const removeHandler = () => remove(index)

	return (
		<Stack direction={'row'} spacing={1}>
			<Controller
				control={methods.control}
				name={`filters.${index}.field`}
				render={({ field, fieldState: { error } }) => (
					<FormControl fullWidth sx={{ maxWidth: 170 }}>
						<InputLabel id={`filters.${index}.field`}>Колонка</InputLabel>

						<Select
							value={field.value}
							onChange={(event: SelectChangeEvent) => {
								const newType = event.target.value.split('@')[1]
								if (newType != type) {
									methods.setValue(`filters.${index}.fieldType`, event.target.value.split('@')[1])
									methods.setValue(
										`filters.${index}.compareType`,
										(compareTypes.get(newType) || 'con') as CompareTypes
									)
								}
								field.onChange(event.target.value)
							}}
							labelId={`filters.${index}.field`}
							label={'Колонка'}
							error={Boolean(error)}
						>
							{columns.map(c => (
								<MenuItem key={c.key} value={`${c.key}@${c.filter}`}>
									{c.label}
								</MenuItem>
							))}
						</Select>
					</FormControl>
				)}
			/>

			<Controller
				name={`filters.${index}.compareType`}
				control={methods.control}
				rules={{ required: true }}
				render={({ field, fieldState: { error } }) => (
					<FormControl fullWidth sx={{ maxWidth: 170 }}>
						<InputLabel id={`filters.${index}.compareType`}>Условие</InputLabel>

						<Select
							{...field}
							error={Boolean(error)}
							labelId={`filters.${index}.compareType`}
							label='Условие'
						>
							{type == 'number' && NumberOptions}
							{type == 'date' && DateOptions}
							{type == 'string' && StringOptions}
						</Select>
					</FormControl>
				)}
			/>

			{type == 'string' && (
				<TextField
					label='Значение'
					{...methods.register(`filters.${index}.value`, { required: true })}
					error={Boolean(
						methods.formState.errors?.filters && methods.formState.errors?.filters[index]?.value
					)}
					fullWidth
				/>
			)}
			{type == 'number' && (
				<TextField
					label='Значение'
					type='number'
					{...methods.register(`filters.${index}.value`, { required: true })}
					error={Boolean(
						methods.formState.errors?.filters && methods.formState.errors?.filters[index]?.value
					)}
					fullWidth
				/>
			)}
			{type == 'date' && (
				<Controller
					control={methods.control}
					name={`filters.${index}.value`}
					rules={{ required: true }}
					render={({ field, fieldState: { error } }) => (
						<DatePicker
							{...field}
							value={field.value ? dayjs(+field.value * 1000) : null}
							onChange={value => field.onChange(value?.startOf('d').unix())}
							label={'Значение'}
							showDaysOutsideCurrentMonth
							fixedWeekNumber={6}
							slotProps={{
								textField: {
									fullWidth: true,
									error: Boolean(error),
								},
							}}
						/>
					)}
				/>
			)}

			{index != 0 && (
				<IconButton onClick={removeHandler}>
					<TimesIcon fontSize={18} padding={0.4} />
				</IconButton>
			)}
		</Stack>
	)
}

const NumberOptions = [
	<MenuItem key='n_eq' value='eq'>
		Равно
	</MenuItem>,
	<MenuItem key='n_gte' value='gte'>
		Больше или равно
	</MenuItem>,
	<MenuItem key='n_lte' value='lte'>
		Меньше или равно
	</MenuItem>,
	// <MenuItem key='n_range' value='range'>
	// 	В диапазоне
	// </MenuItem>,
]

const DateOptions = [
	<MenuItem key='d_eq' value='eq'>
		Равна
	</MenuItem>,
	<MenuItem key='d_gte' value='gte'>
		Больше или равна
	</MenuItem>,
	<MenuItem key='d_lte' value='lte'>
		Меньше или равна
	</MenuItem>,
	// <MenuItem value='range'>В диапазоне</MenuItem>,
]

const StringOptions = [
	<MenuItem key='con' value='con'>
		Содержит
	</MenuItem>,
	<MenuItem key='like' value='like'>
		Равен
	</MenuItem>,
	<MenuItem key='start' value='start'>
		Начинается с
	</MenuItem>,
	<MenuItem key='end' value='end'>
		Заканчивается на
	</MenuItem>,
]
