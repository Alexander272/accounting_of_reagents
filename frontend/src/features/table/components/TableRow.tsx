import { CSSProperties, FC } from 'react'

import { ColWidth } from '@/constants/defaultValues'
import { TableRow } from '@/components/Table/TableRow'
import { TableCell } from '@/components/Table/TableCell'
import { CellText } from '@/components/CellText/CellText'
import type { IDataItem } from '../types/data'
import { Columns } from '../columns'

type Props = {
	data: IDataItem
	sx?: CSSProperties
}

export const DataTableRow: FC<Props> = ({ data, sx }) => {
	//TODO цвет я наверное с сервера буду получать
	return (
		<TableRow sx={sx}>
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
