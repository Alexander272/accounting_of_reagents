import { CSSProperties, FC, MouseEvent } from 'react'

import type { IDataItem } from '../types/data'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { ColWidth } from '@/constants/defaultValues'
import { TableRow } from '@/components/Table/TableRow'
import { TableCell } from '@/components/Table/TableCell'
import { CellText } from '@/components/CellText/CellText'
import { Columns } from '../columns'
import { getContextMenu, getSelected, setContextMenu, setSelected } from '../tableSlice'

type Props = {
	data: IDataItem
	sx?: CSSProperties
}

export const DataTableRow: FC<Props> = ({ data, sx }) => {
	//TODO цвет я наверное с сервера буду получать
	const selected = useAppSelector(getSelected)
	const contextMenu = useAppSelector(getContextMenu)
	const dispatch = useAppDispatch()

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

	const background = contextMenu?.active == data.id ? '#c6d6ff' : selected[data.id] ? '#dde6fd' : data.background

	return (
		<TableRow
			onClick={selectHandler}
			onContext={contextHandler}
			sx={{
				...sx,
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
