import { CSSProperties, FC } from 'react'

import type { IDataItem } from '../types/data'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { ColWidth } from '@/constants/defaultValues'
import { TableRow } from '@/components/Table/TableRow'
import { TableCell } from '@/components/Table/TableCell'
import { CellText } from '@/components/CellText/CellText'
import { Columns } from '../columns'
import { getSelected, setSelected } from '../tableSlice'

type Props = {
	data: IDataItem
	sx?: CSSProperties
}

export const DataTableRow: FC<Props> = ({ data, sx }) => {
	//TODO цвет я наверное с сервера буду получать
	const selected = useAppSelector(getSelected)
	const dispatch = useAppDispatch()

	const selectHandler = () => {
		dispatch(setSelected(data.id))
	}

	return (
		<TableRow onClick={selectHandler} sx={{ ...sx, background: selected[data.id] ? '#dde6fd' : data.background }}>
			{Columns.map(c => (
				<TableCell key={data.id + c.key} width={c.width || ColWidth}>
					<CellText
						value={
							c.formatter
								? c.formatter(data[c.key as keyof IDataItem])
								: data[c.key as keyof IDataItem]?.toString() || '-'
						}
					/>
				</TableCell>
			))}
		</TableRow>
	)
}
