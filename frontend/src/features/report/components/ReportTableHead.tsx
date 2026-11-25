import { ReportColumnAligns, ReportColumnWidths, ReportHeaders } from '../constants/table'
import { TableHead } from '@/components/Table/TableHead'
import { TableRow } from '@/components/Table/TableRow'
import { TableCell } from '@/components/Table/TableCell'
import { CellText } from '@/components/CellText/CellText'

export const ReportTableHead = () => (
	<TableHead sx={{ borderRadius: 0 }}>
		<TableRow height={40}>
			{ReportHeaders.map((label, i) => (
				<TableCell key={label} width={ReportColumnWidths[i]} align={ReportColumnAligns[i]}>
					<CellText value={label} />
				</TableCell>
			))}
		</TableRow>
	</TableHead>
)
