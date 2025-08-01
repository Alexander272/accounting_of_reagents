import { FC, useEffect } from 'react'
import { CircularProgress, FormControl, InputLabel, MenuItem, OutlinedInput, Select } from '@mui/material'
import { Controller, useFormContext } from 'react-hook-form'

import type { IFilter } from '../../types/data'
import type { IColumn, IFullFilter } from '../../types/table'
import type { IReagentType } from '../../modules/Types/types'

type Props = {
	index: number
	columns: IColumn[]
}

export const ListFilter: FC<Props> = ({ index, columns }) => {
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
			rules={{ required: true }}
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
