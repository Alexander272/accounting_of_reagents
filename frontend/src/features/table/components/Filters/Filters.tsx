import { FC, useEffect, useRef, useState } from 'react'
import {
	Button,
	CircularProgress,
	FormControl,
	FormControlLabel,
	IconButton,
	InputLabel,
	MenuItem,
	OutlinedInput,
	Select,
	SelectChangeEvent,
	Stack,
	Switch,
	TextField,
	Typography,
	useTheme,
} from '@mui/material'
import { Controller, FormProvider, useFieldArray, useForm, useFormContext } from 'react-hook-form'
import { DatePicker } from '@mui/x-date-pickers'
import dayjs from 'dayjs'

import type { CompareTypes, IFilter } from '../../types/data'
import type { IFullFilter } from '../../types/table'
import type { IReagentType } from '../../modules/Types/types'
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
	field: `${columns[1].key}@${columns[1].filter}`,
	fieldType: columns[1].filter as 'string',
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
	['switch', 'eq'],
	['list', 'in'],
])
const defaultValues = new Map([
	['string', ''],
	['date', dayjs().unix().toString()],
	['number', ''],
	['switch', 'false'],
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
		<Stack direction={'row'} spacing={1} alignItems={'center'}>
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
									methods.setValue(`filters.${index}.value`, defaultValues.get(newType) || '')
								}
								field.onChange(event.target.value)
							}}
							labelId={`filters.${index}.field`}
							label={'Колонка'}
							error={Boolean(error)}
						>
							{columns.map(c => (
								<MenuItem
									key={c.key}
									value={`${c.key}@${typeof c.filter == 'string' ? c.filter : c.filter?.type}`}
								>
									{c.label}
								</MenuItem>
							))}
						</Select>
					</FormControl>
				)}
			/>

			{type == 'string' && <StringFilter index={index} />}
			{type == 'number' && <NumberFilter index={index} />}
			{type == 'date' && <DateFilter index={index} />}
			{type == 'switch' && <SwitchFilter index={index} />}
			{type == 'list' && <ListFilter index={index} />}

			{index != 0 && (
				<IconButton onClick={removeHandler}>
					<TimesIcon fontSize={18} padding={0.4} />
				</IconButton>
			)}
		</Stack>
	)
}

type FilterProps = {
	index: number
}
const StringFilter: FC<FilterProps> = ({ index }) => {
	const methods = useFormContext<{ filters: IFilter[] }>()

	return (
		<>
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
							<MenuItem key='con' value='con'>
								Содержит
							</MenuItem>
							<MenuItem key='like' value='like'>
								Равен
							</MenuItem>
							<MenuItem key='start' value='start'>
								Начинается с
							</MenuItem>
							<MenuItem key='end' value='end'>
								Заканчивается на
							</MenuItem>
						</Select>
					</FormControl>
				)}
			/>

			<TextField
				label='Значение'
				{...methods.register(`filters.${index}.value`, { required: true })}
				error={Boolean(methods.formState.errors?.filters && methods.formState.errors?.filters[index]?.value)}
				fullWidth
			/>
		</>
	)
}
const NumberFilter: FC<FilterProps> = ({ index }) => {
	const methods = useFormContext<{ filters: IFilter[] }>()

	return (
		<>
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
							<MenuItem key='n_eq' value='eq'>
								Равно
							</MenuItem>
							<MenuItem key='n_gte' value='gte'>
								Больше или равно
							</MenuItem>
							<MenuItem key='n_lte' value='lte'>
								Меньше или равно
							</MenuItem>
						</Select>
					</FormControl>
				)}
			/>

			<TextField
				label='Значение'
				type='number'
				{...methods.register(`filters.${index}.value`, { required: true })}
				error={Boolean(methods.formState.errors?.filters && methods.formState.errors?.filters[index]?.value)}
				fullWidth
			/>
		</>
	)
}
const DateFilter: FC<FilterProps> = ({ index }) => {
	const methods = useFormContext<{ filters: IFilter[] }>()

	return (
		<>
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
							<MenuItem key='d_eq' value='eq'>
								Равна
							</MenuItem>
							<MenuItem key='d_gte' value='gte'>
								Больше или равна
							</MenuItem>
							<MenuItem key='d_lte' value='lte'>
								Меньше или равна
							</MenuItem>
						</Select>
					</FormControl>
				)}
			/>

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
		</>
	)
}

const SwitchFilter: FC<FilterProps> = ({ index }) => {
	const methods = useFormContext<{ filters: IFilter[] }>()
	const field = methods.watch(`filters.${index}.field`).split('@')[0]
	const filter = columns.find(c => c.key == field)?.filter as {
		type: 'switch'
		options: { true: string; false: string }
	}

	return (
		<Controller
			control={methods.control}
			name={`filters.${index}.value`}
			render={({ field }) => (
				<FormControlLabel
					label={filter?.options[(field.value as 'true') || false]}
					control={
						<Switch
							checked={field.value == 'true'}
							onChange={event => field.onChange(event.target.checked.toString())}
						/>
					}
				/>
			)}
		/>
	)
}
const ListFilter: FC<FilterProps> = ({ index }) => {
	const methods = useFormContext<{ filters: IFilter[] }>()
	const field = methods.watch(`filters.${index}.field`).split('@')[0]
	const value = methods.watch(`filters.${index}.value`)
	const filter = columns.find(c => c.key == field)?.filter as IFullFilter
	const { data, isLoading }: { data?: { data: IReagentType[] }; isLoading: boolean } =
		filter.getOptions && filter.getOptions(null)

	useEffect(() => {
		if (data && value == '') methods.setValue(`filters.${index}.value`, data.data[0].name)
	}, [value, methods, index, data])

	if (isLoading) return <CircularProgress size={20} />
	return (
		<Controller
			control={methods.control}
			name={`filters.${index}.value`}
			// rules={{ required: true }}
			render={({ field, fieldState: { error } }) => (
				<FormControl fullWidth>
					<InputLabel id={`filters.${index}.value`}>Значение</InputLabel>
					<Select
						multiple
						labelId={`filters.${index}.value`}
						value={field.value.split('|')}
						error={Boolean(error)}
						onChange={event =>
							field.onChange(
								typeof event.target.value === 'string'
									? event.target.value
									: event.target.value.join('|')
							)
						}
						input={<OutlinedInput label='Значение' />}
					>
						{data?.data.map(r => (
							<MenuItem key={r.id} value={r.name}>
								{r.description || r.name}
							</MenuItem>
						))}
					</Select>
				</FormControl>
			)}
		/>
	)
}
