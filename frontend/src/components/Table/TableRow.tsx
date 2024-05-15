import type { CSSProperties, FC, MouseEvent, PropsWithChildren } from 'react'

import { TableRowContainer } from './style'

type Props = {
	onClick?: () => void
	onContext?: (event: MouseEvent<HTMLDivElement>) => void
	width?: number
	height?: number
	sx?: CSSProperties
}

export const TableRow: FC<PropsWithChildren<Props>> = ({ children, onClick, onContext, width, height, sx }) => {
	return (
		<TableRowContainer onClick={onClick} onContextMenu={onContext} width={width} height={height} styles={sx}>
			{children}
		</TableRowContainer>
	)
}
