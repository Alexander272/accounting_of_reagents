import { FC } from 'react'
import { Stack, TextField } from '@mui/material'
import { useFormContext } from 'react-hook-form'

import type { Field } from '../type'
import type { IBaseInfForm } from './type'
import { Titles } from './titles'

const fields: Field<keyof IBaseInfForm>[] = [
	{ key: 'type', type: 'Radio', label: Titles.Type, options: [] },
	{ key: 'name', type: 'String', label: Titles.Name, multiline: true, minRows: 1 },
	{ key: 'uname', type: 'String', label: Titles.UName, multiline: true, minRows: 1 },
	{ key: 'document', type: 'String', label: Titles.Doc, multiline: true, minRows: 1 },
	{ key: 'purity', type: 'String', label: Titles.Purity },
	{ key: 'dateOfManufacture', type: 'Date', label: Titles.DateOfManufacture },
	{ key: 'consignment', type: 'String', label: Titles.Consignment },
	{ key: 'manufacturer', type: 'String', label: Titles.Manufacturer },
	{ key: 'shelfLife', type: 'Number', label: Titles.ShelfLife },
	{ key: 'place_closet', type: 'String', label: Titles.Place.Closet },
	{ key: 'place_shelf', type: 'Number', label: Titles.Place.Shelf },
]

type Props = {
	disabled?: boolean
}

export const Inputs: FC<Props> = ({ disabled }) => {
	const {
		register,
		formState: { errors },
	} = useFormContext<IBaseInfForm>()

	return (
		<Stack spacing={2}>
			{fields.map(f => {
				switch (f.type) {
					case 'Radio':
						//TODO сделать компонент
						return
					case 'Number':
						return (
							<TextField
								key={f.key}
								label={f.label}
								type='number'
								disabled={disabled}
								error={Boolean(errors[f.key])}
								{...register(f.key)}
							/>
						)

					case 'String':
						return (
							<TextField
								key={f.key}
								label={f.label}
								type='number'
								disabled={disabled}
								multiline={f.multiline}
								minRows={f.minRows}
								error={Boolean(errors[f.key])}
								{...register(f.key)}
							/>
						)
				}
			})}
		</Stack>
	)
}
