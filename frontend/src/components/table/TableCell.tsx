import { FC, PropsWithChildren } from 'react'

import { TableCellContainer } from './style'

type Props = {
	colSpan?: number
	rowSpan?: number
}

export const TableCell: FC<PropsWithChildren<Props>> = ({ children }) => {
	return <TableCellContainer>{children}</TableCellContainer>
}
