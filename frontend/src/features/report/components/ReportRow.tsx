import { CSSProperties } from 'react'
import { Typography } from '@mui/material'

import type { ISpendingReportItem } from '../types'
import { ReportColumnAligns, ReportColumnWidths } from '../constants/table'
import { formatNumber } from '../utils/format'
import { TableRow } from '@/components/Table/TableRow'
import { TableCell } from '@/components/Table/TableCell'
import { CellText } from '@/components/CellText/CellText'

type Props = {
	item: ISpendingReportItem
	sx?: CSSProperties
}

export const ReportRow = ({ item, sx }: Props) => {
	const remainderPercent = (item.remainder / item.amountAtStart) * 100
	const percentColor = remainderPercent > 40 ? '#6b778c' : remainderPercent > 20 ? '#dd9000' : '#c52424'

	return (
		<TableRow height={37} hover sx={{ ...sx }}>
			<TableCell width={ReportColumnWidths[0]} align={ReportColumnAligns[0]}>
				<CellText value={item.name || '—'} align={ReportColumnAligns[0]} />
			</TableCell>
			<TableCell width={ReportColumnWidths[1]} align={ReportColumnAligns[1]}>
				<CellText value={item.uname || '—'} align={ReportColumnAligns[1]} />
			</TableCell>
			<TableCell width={ReportColumnWidths[2]} align={ReportColumnAligns[2]}>
				<CellText value={item.consignment || '—'} align={ReportColumnAligns[2]} />
			</TableCell>
			<TableCell width={ReportColumnWidths[3]} align={ReportColumnAligns[3]}>
				<Typography sx={{ fontVariantNumeric: 'tabular-nums' }}>{formatNumber(item.amountAtStart)}</Typography>
				<Typography sx={{ color: '#6b778c', ml: 0.5, pr: 0.75 }}>{item.unit}</Typography>
			</TableCell>
			<TableCell width={ReportColumnWidths[4]} align={ReportColumnAligns[4]}>
				<Typography sx={{ fontVariantNumeric: 'tabular-nums', fontWeight: 'bold' }}>
					{formatNumber(item.spent)}
				</Typography>
				<Typography sx={{ color: '#6b778c', ml: 0.5, pr: 0.75 }}>{item.unit}</Typography>
			</TableCell>
			<TableCell width={ReportColumnWidths[5]} align={ReportColumnAligns[5]}>
				<Typography sx={{ fontVariantNumeric: 'tabular-nums', pl: 0.75 }}>
					{formatNumber(item.remainder)}
				</Typography>
				<Typography sx={{ color: '#6b778c', ml: 0.5 }}>{item.unit}</Typography>
				<Typography sx={{ color: percentColor, ml: 0.5, width: 52, pr: 0.75 }}>
					{/* ·{' '} */}
					{new Intl.NumberFormat('ru-RU', { maximumFractionDigits: 0 }).format(remainderPercent)}%
				</Typography>
			</TableCell>
		</TableRow>
	)
}
