import { FC, useState } from 'react'
import { Autocomplete, FormControl, InputLabel, MenuItem, Select, Stack, TextField } from '@mui/material'
import { DatePicker } from '@mui/x-date-pickers'
import { Controller, useFormContext } from 'react-hook-form'
import dayjs from 'dayjs'

import type { AutoCompleteField, Field } from '../type'
import type { IBaseInfForm } from './type'
import { typesApiSlice } from '@/features/table/modules/Types/typesApiSlice'
import { Titles } from './titles'
import { useLazyGetUniqueFieldQuery } from '@/features/table/tableApiSlice'

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
	{
		key: 'name',
		type: 'AutoComplete',
		label: Titles.Name + ' *',
		multiline: true,
		minRows: 1,
		rules: { required: true, validate: value => Boolean(value.trim()) },
	},
	{ key: 'uname', type: 'AutoComplete', label: Titles.UName, multiline: true, minRows: 1 },
	{ key: 'document', type: 'AutoComplete', label: Titles.Doc, multiline: true, minRows: 1 },
	{ key: 'purity', type: 'AutoComplete', label: Titles.Purity },
	{
		key: 'dateOfManufacture',
		type: 'Date',
		label: Titles.DateOfManufacture + ' *',
		rules: { required: true, min: 1000000000 },
	},
	{ key: 'consignment', type: 'String', label: Titles.Consignment },
	{ key: 'manufacturer', type: 'AutoComplete', label: Titles.Manufacturer },
	{ key: 'shelfLife', type: 'Number', label: Titles.ShelfLife + ' *', rules: { required: true, min: 1 } },
	{ key: 'place_closet', type: 'AutoComplete', label: Titles.Place.Closet + ' *', rules: { required: true } },
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

	const { control } = useFormContext<IBaseInfForm>()

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
							<Controller
								key={f.key}
								control={control}
								name={f.key}
								rules={f.rules}
								render={({ field, fieldState: { error } }) => (
									<TextField
										{...field}
										label={f.label}
										type='number'
										disabled={disabled}
										error={Boolean(error)}
									/>
								)}
							/>
						)

					case 'String':
						return (
							<Controller
								key={f.key}
								control={control}
								name={f.key}
								rules={f.rules}
								render={({ field, fieldState: { error } }) => (
									<TextField
										{...field}
										label={f.label}
										disabled={disabled}
										error={Boolean(error)}
										multiline={f.multiline}
										minRows={f.minRows}
									/>
								)}
							/>
						)

					case 'AutoComplete':
						return <AutocompleteField f={f} disabled={disabled} />
				}
			})}
		</Stack>
	)
}

const AutocompleteField: FC<{ f: AutoCompleteField<keyof IBaseInfForm>; disabled?: boolean }> = ({ f, disabled }) => {
	const { control } = useFormContext<IBaseInfForm>()
	const [options, setOptions] = useState<string[]>([])

	const [getUnique, { isLoading }] = useLazyGetUniqueFieldQuery()

	const focusHandler = async () => {
		const data = await getUnique(f.key).unwrap()
		setOptions(data.data || [])
	}

	return (
		<Controller
			key={f.key}
			control={control}
			name={f.key}
			rules={f.rules}
			render={({ field, fieldState: { error } }) => (
				<Autocomplete
					value={field.value.toString() || ''}
					freeSolo
					disableClearable
					autoComplete
					options={options}
					loading={isLoading}
					loadingText='Поиск похожих значений...'
					noOptionsText='Ничего не найдено'
					// getOptionLabel={o => o.name}
					onChange={(_event, value) => {
						field.onChange(value)
					}}
					onFocus={focusHandler}
					renderInput={params => (
						<TextField
							{...params}
							label={f.label}
							error={Boolean(error)}
							disabled={disabled}
							onChange={field.onChange}
							inputRef={field.ref}
						/>
					)}
				/>
			)}
		/>
	)
}
