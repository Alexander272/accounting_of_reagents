import { FC } from 'react'
import { FormControl, InputLabel, MenuItem, Select, Stack, TextField } from '@mui/material'
import { DatePicker } from '@mui/x-date-pickers'
import { Controller, useFormContext } from 'react-hook-form'
import dayjs from 'dayjs'

import type { Field } from '../type'
import type { IBaseInfForm } from './type'
import { typesApiSlice } from '@/features/table/modules/Types/typesApiSlice'
import { Titles } from './titles'

const fields: Field<keyof IBaseInfForm>[] = [
	{
		key: 'type',
		type: 'List',
		label: Titles.Type,
		useGetOptions: () => {
			const { data } = typesApiSlice.endpoints.getReagentTypes.useQueryState(null, undefined)
			return data?.data || []
		},
	},
	{ key: 'name', type: 'String', label: Titles.Name + ' *', multiline: true, minRows: 1, rules: { required: true } },
	{ key: 'uname', type: 'String', label: Titles.UName, multiline: true, minRows: 1 },
	{ key: 'document', type: 'String', label: Titles.Doc, multiline: true, minRows: 1 },
	{ key: 'purity', type: 'String', label: Titles.Purity },
	{ key: 'dateOfManufacture', type: 'Date', label: Titles.DateOfManufacture + ' *', rules: { required: true } },
	{ key: 'consignment', type: 'String', label: Titles.Consignment },
	{ key: 'manufacturer', type: 'String', label: Titles.Manufacturer },
	{ key: 'shelfLife', type: 'Number', label: Titles.ShelfLife + ' *', rules: { required: true } },
	{ key: 'place_closet', type: 'String', label: Titles.Place.Closet + ' *', rules: { required: true } },
	{ key: 'place_shelf', type: 'Number', label: Titles.Place.Shelf },
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

	const {
		register,
		control,
		formState: { errors },
	} = useFormContext<IBaseInfForm>()

	return (
		<Stack spacing={2}>
			{fields.map(f => {
				switch (f.type) {
					case 'List':
						return (
							<Controller
								key={f.key}
								control={control}
								name={f.key}
								render={({ field, fieldState: { error } }) => (
									<FormControl>
										<InputLabel id={f.key}>{f.label}</InputLabel>

										<Select labelId={f.key} label={f.label} error={Boolean(error)} {...field}>
											{(f.options || []).map(o => (
												<MenuItem key={o.id} value={o.id}>
													{o.description || o.name}
												</MenuItem>
											))}
										</Select>
									</FormControl>
								)}
							/>
						)

					case 'Date':
						return (
							<Controller
								key={f.key}
								control={control}
								name={f.key}
								rules={f.rules}
								render={({ field, fieldState: { error } }) => (
									<DatePicker
										{...field}
										value={dayjs(+field.value * 1000)}
										onChange={value => field.onChange(value?.startOf('d').unix())}
										label={f.label}
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
						)

					case 'Number':
						return (
							<TextField
								key={f.key}
								label={f.label}
								type='number'
								disabled={disabled}
								error={Boolean(errors[f.key])}
								{...register(f.key, { ...f.rules })}
							/>
						)

					case 'String':
						return (
							<TextField
								key={f.key}
								label={f.label}
								disabled={disabled}
								multiline={f.multiline}
								minRows={f.minRows}
								error={Boolean(errors[f.key])}
								{...register(f.key, { ...f.rules })}
							/>
						)
				}
			})}
		</Stack>
	)
}
