import type { CSSProperties, FC, PropsWithChildren } from 'react'

import { TableRowContainer } from './style'

type Props = {
	onClick?: () => void
	width?: number
	height?: number
	sx?: CSSProperties
}

export const TableRow: FC<PropsWithChildren<Props>> = ({ children, onClick, width, height, sx }) => {
	return (
		<TableRowContainer onClick={onClick} width={width} height={height} styles={sx}>
			{children}
		</TableRowContainer>
	)
}
