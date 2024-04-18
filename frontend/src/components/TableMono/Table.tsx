import { FC } from 'react'

import type { IColumn, IGroup } from './types'
import { TableContainer } from './style'
import { TableHeader } from './TableHeader'

type Props = {
	columns: IColumn[]
	groups?: IGroup[]
}

export const Table: FC<Props> = ({ columns, groups }) => {
	return (
		<TableContainer>
			<TableHeader columns={columns} groups={groups} />
		</TableContainer>
	)
}
