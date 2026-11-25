import { FC, PropsWithChildren } from 'react'

import type { TableCellAlign } from './types'
import { TableCellContainer } from './style'

type Props = {
	width?: number
	colSpan?: number
	rowSpan?: number
	isActive?: boolean
	align?: TableCellAlign
	onClick?: () => void
}

export const TableCell: FC<PropsWithChildren<Props>> = ({ children, width, isActive, align, onClick }) => {
	return (
		<TableCellContainer width={width} isActive={isActive} align={align} onClick={onClick}>
			{children}
		</TableCellContainer>
	)
}
