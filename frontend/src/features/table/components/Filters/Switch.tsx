import { FC } from 'react'
import { FormControlLabel, Switch } from '@mui/material'
import { Controller, useFormContext } from 'react-hook-form'

import type { IFilter } from '../../types/data'
import type { IColumn } from '../../types/table'

type Props = {
	index: number
	columns: IColumn[]
}

export const SwitchFilter: FC<Props> = ({ index, columns }) => {
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
