import { CSSProperties, FC, MouseEvent } from 'react'
import { useTheme } from '@mui/material'

import type { IDataItem } from '../types/data'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { ColWidth } from '@/constants/defaultValues'
import { TableRow } from '@/components/Table/TableRow'
import { TableCell } from '@/components/Table/TableCell'
import { CellText } from '@/components/CellText/CellText'
import { Columns } from '../columns'
import { getContextMenu, getHidden, getSelected, setContextMenu, setSelected } from '../tableSlice'

type Props = {
	data: IDataItem
	sx?: CSSProperties
}

export const DataTableRow: FC<Props> = ({ data, sx }) => {
	const selected = useAppSelector(getSelected)
	const hidden = useAppSelector(getHidden)
	const contextMenu = useAppSelector(getContextMenu)
	const dispatch = useAppDispatch()

	const { palette } = useTheme()

	const selectHandler = () => {
		dispatch(setSelected(data.id))
	}

	const contextHandler = (event: MouseEvent<HTMLDivElement>) => {
		event.preventDefault()
		const menu = {
			active: data.id,
			coords: { mouseX: event.clientX + 2, mouseY: event.clientY - 6 },
		}
		dispatch(setContextMenu(menu))
	}

	let background = data.itemStyle?.background
	if (selected[data.id]) background = palette.rowActive.light
	if (contextMenu?.active == data.id) background = palette.rowActive.main

	return (
		<TableRow
			onClick={selectHandler}
			onContext={contextHandler}
			hover
			sx={{
				...sx,
				padding: '0 6px',
				background: background,
			}}
		>
			{Columns.map(c => {
				if (hidden[c.key]) return null

				let value = data[c.key as keyof IDataItem]?.toString() || '-'
				if (c.formatter) value = c.formatter(data[c.key as keyof IDataItem])

				return (
					<TableCell key={data.id + c.key} width={c.width || ColWidth}>
						<CellText value={value} color={data.itemStyle?.textColor} />
					</TableCell>
				)
			})}
		</TableRow>
	)
}
