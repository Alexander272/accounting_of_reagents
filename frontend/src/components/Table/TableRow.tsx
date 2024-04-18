import type { CSSProperties, FC, PropsWithChildren } from 'react'

import { TableRowContainer } from './style'

type Props = {
	width?: number
	height?: number
	sx?: CSSProperties
}

export const TableRow: FC<PropsWithChildren<Props>> = ({ children, ...style }) => {
	return <TableRowContainer {...style}>{children}</TableRowContainer>
}
