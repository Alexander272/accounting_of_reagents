import { FC, PropsWithChildren } from 'react'

import { TableContainer } from './style'

type Props = unknown

export const Table: FC<PropsWithChildren<Props>> = ({ children }) => {
	return <TableContainer>{children}</TableContainer>
}
