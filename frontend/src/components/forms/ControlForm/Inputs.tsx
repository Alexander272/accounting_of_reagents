import { FC, MouseEvent } from 'react'
import {
	Button,
	Divider,
	FormControl,
	InputLabel,
	MenuItem,
	Select,
	SelectChangeEvent,
	Stack,
	TextField,
} from '@mui/material'
import { DatePicker } from '@mui/x-date-pickers'
import { Controller, useFormContext } from 'react-hook-form'
import dayjs from 'dayjs'

import type { Field } from '../type'
import type { IControlForm } from './type'
import { typesApiSlice } from '@/features/table/modules/Types/typesApiSlice'
import { Titles } from './titles'
import { useAppDispatch } from '@/hooks/redux'
import { changeModalIsOpen } from '@/features/modal/modalSlice'

const fields: Field<keyof IControlForm>[] = [
	{ key: 'receiptDate', type: 'Date', label: Titles.ReceiptDate + ' *', rules: { required: true, min: 1000000000 } },
	{ key: 'amount', type: 'Number', label: Titles.Amount + ' *', rules: { required: true, min: 1 } },
	{
		key: 'amountType',
		type: 'List',
		label: Titles.AmountType,
		useGetOptions: () => {
			const { data } = typesApiSlice.endpoints.getAmountTypes.useQueryState(null, undefined)
			return data?.data || []
		},
	},
	{ key: 'controlDate', type: 'Date', label: Titles.Date, rules: { min: 1000000000 } },
	{ key: 'protocol', type: 'String', label: Titles.Protocol },
	{
		key: 'result',
		type: 'List',
		label: Titles.Result,
		options: [
			{ id: 'true', name: 'Соответствует' },
			{ id: 'false', name: 'Несоответствует' },
		],
	},
]

type Props = {
	disabled?: boolean
}

export const Inputs: FC<Props> = ({ disabled }) => {
	fields.forEach(f => {
		if (f.type == 'List' && f.useGetOptions) {
			f.options = f.useGetOptions()
		}
	})

	const dispatch = useAppDispatch()

	const {
		register,
		control,
		formState: { errors },
	} = useFormContext<IControlForm>()

	const editHandler = (event: MouseEvent<HTMLButtonElement>) => {
		event.preventDefault()
		event.stopPropagation()
		console.log('edit')
		dispatch(changeModalIsOpen({ variant: 'amountType', isOpen: true }))
	}

	return (
		<Stack spacing={2}>
			<Controller
				key={fields[0].key}
				control={control}
				name={fields[0].key}
				rules={fields[0].rules}
				render={({ field, fieldState: { error } }) => (
					<DatePicker
						{...field}
						value={dayjs(+field.value * 1000)}
						onChange={value => field.onChange(value?.startOf('d').unix())}
						label={fields[0].label}
						showDaysOutsideCurrentMonth
						fixedWeekNumber={6}
						disabled={disabled}
						slotProps={{
							textField: {
								error: Boolean(error),
							},
						}}
					/>
				)}
			/>

			<Stack direction={'row'}>
				<TextField
					key={fields[1].key}
					label={fields[1].label}
					type='number'
					disabled={disabled}
					error={Boolean(errors[fields[1].key])}
					{...register(fields[1].key, { ...fields[1].rules })}
					sx={{
						flexBasis: '72%',
						'& fieldset': {
							borderTopRightRadius: 0,
							borderBottomRightRadius: 0,
						},
					}}
				/>
				<Controller
					key={fields[2].key}
					control={control}
					name={fields[2].key}
					render={({ field, fieldState: { error } }) => (
						<FormControl sx={{ flexBasis: '28%', ml: '-1px' }}>
							<InputLabel id={fields[2].key}>{fields[2].label}</InputLabel>
							<Select
								labelId={fields[2].key}
								label={fields[2].label}
								error={Boolean(error)}
								value={field.value}
								onChange={field.onChange}
								fullWidth
								sx={{
									borderTopLeftRadius: 0,
									borderBottomLeftRadius: 0,
								}}
							>
								{((fields[2].type == 'List' && fields[2].options) || []).map(o => (
									<MenuItem key={o.id} value={o.id}>
										{o.description || o.name}
									</MenuItem>
								))}
								<Divider />
								<MenuItem sx={{ padding: 0 }}>
									<Button
										onClick={editHandler}
										fullWidth
										sx={{ textTransform: 'capitalize', padding: '4px 8px' }}
									>
										Изменить
									</Button>
								</MenuItem>
							</Select>
						</FormControl>
					)}
				/>
			</Stack>

			<Controller
				key={fields[3].key}
				control={control}
				name={fields[3].key}
				rules={fields[3].rules}
				render={({ field, fieldState: { error } }) => (
					<DatePicker
						{...field}
						value={dayjs(+field.value * 1000)}
						onChange={value => field.onChange(value?.startOf('d').unix())}
						label={fields[3].label}
						showDaysOutsideCurrentMonth
						fixedWeekNumber={6}
						disabled={disabled}
						slotProps={{
							textField: {
								error: Boolean(error),
							},
						}}
					/>
				)}
			/>

			{fields[4].type == 'String' ? (
				<TextField
					key={fields[4].key}
					label={fields[4].label}
					disabled={disabled}
					multiline={fields[4].multiline}
					minRows={fields[4].minRows}
					error={Boolean(errors[fields[4].key])}
					{...register(fields[4].key, { ...fields[4].rules })}
				/>
			) : null}

			{fields[5].type == 'List' ? (
				<Controller
					key={fields[5].key}
					control={control}
					name={fields[5].key}
					render={({ field, fieldState: { error } }) => (
						<FormControl>
							<InputLabel id={fields[5].key}>{fields[5].label}</InputLabel>

							<Select
								labelId={fields[5].key}
								label={fields[5].label}
								error={Boolean(error)}
								value={field.value.toString()}
								onChange={(event: SelectChangeEvent) => field.onChange(event.target.value == 'true')}
							>
								{((fields[5].type == 'List' && fields[5].options) || []).map(o => (
									<MenuItem key={o.id} value={o.id}>
										{o.name}
									</MenuItem>
								))}
							</Select>
						</FormControl>
					)}
				/>
			) : null}
		</Stack>
	)
}
