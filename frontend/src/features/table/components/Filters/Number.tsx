import { FC } from 'react'
import { FormControl, InputLabel, MenuItem, Select, TextField } from '@mui/material'
import { Controller, useFormContext } from 'react-hook-form'

import type { IFilter } from '../../types/data'

type Props = {
	index: number
}

export const NumberFilter: FC<Props> = ({ index }) => {
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
