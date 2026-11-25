import { Box } from '@mui/material'

import type { ISpendingReportItem } from '../types'
import { ReportTableHead } from './ReportTableHead'
import { ReportTableBody } from './ReportTableBody'
import { Table } from '@/components/Table/Table'
import { Fallback } from '@/components/Fallback/Fallback'
import { NoRowsOverlay } from '@/features/table/components/NoRowsOverlay/components/NoRowsOverlay'

type Props = {
	items: ISpendingReportItem[]
	isFetching: boolean
	isUninitialized: boolean
}

export const ReportTable = ({ items, isFetching, isUninitialized }: Props) => (
	<Box sx={{ maxHeight: 600 }}>
		{isFetching ? <Fallback marginTop={5} /> : null}

		{!isFetching && !isUninitialized && !items.length ? <NoRowsOverlay /> : null}

		{!isFetching && items.length ? (
			<Table sx={{ marginTop: 0 }}>
				<ReportTableHead />
				<ReportTableBody items={items} />
			</Table>
		) : null}
	</Box>
)
