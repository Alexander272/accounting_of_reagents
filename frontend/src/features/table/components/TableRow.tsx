import { CSSProperties, FC, MouseEvent } from 'react'

import type { IDataItem } from '../types/data'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { ColWidth } from '@/constants/defaultValues'
import { TableRow } from '@/components/Table/TableRow'
import { TableCell } from '@/components/Table/TableCell'
import { CellText } from '@/components/CellText/CellText'
import { Columns } from '../columns'
import { getContextMenu, getSelected, setContextMenu, setSelected } from '../tableSlice'
import { useTheme } from '@mui/material'

type Props = {
	data: IDataItem
	sx?: CSSProperties
}

export const DataTableRow: FC<Props> = ({ data, sx }) => {
	const selected = useAppSelector(getSelected)
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

	let background = data.background
	if (selected[data.id]) background = palette.rowActive.light
	if (contextMenu?.active == data.id) background = palette.rowActive.main
	// const background =
	// 	contextMenu?.active == data.id
	// 		? palette.rowActive.main
	// 		: selected[data.id]
	// 		? palette.rowActive.light
	// 		: data.background

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
