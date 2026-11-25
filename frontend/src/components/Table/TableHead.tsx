import { CSSProperties, FC, PropsWithChildren } from 'react'

import { TableHeadContainer } from './style'

type Props = {
	sx?: CSSProperties
}

export const TableHead: FC<PropsWithChildren<Props>> = ({ children, sx }) => {
	return <TableHeadContainer styles={sx}>{children}</TableHeadContainer>
}
