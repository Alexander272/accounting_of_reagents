import { FC } from 'react'
import { FormControl, InputLabel, MenuItem, Select, TextField } from '@mui/material'
import { Controller, useFormContext } from 'react-hook-form'

import type { IFilter } from '../../types/data'

type Props = {
	index: number
}

export const StringFilter: FC<Props> = ({ index }) => {
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
