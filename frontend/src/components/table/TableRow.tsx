import { FC, PropsWithChildren } from 'react'

import { TableRowContainer } from './style'

export const TableRow: FC<PropsWithChildren> = ({ children }) => {
	return <TableRowContainer>{children}</TableRowContainer>
}
