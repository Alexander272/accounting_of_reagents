import { FixedSizeList } from 'react-window'

import type { ISpendingReportItem } from '../types'
import { ReportRow } from './ReportRow'
import { TableBody } from '@/components/Table/TableBody'
import { RowHeight } from '@/constants/defaultValues'
import { ReportColumnWidths } from '../constants/table'

type Props = {
	items: ISpendingReportItem[]
}

export const ReportTableBody = ({ items }: Props) => (
	<TableBody>
		<FixedSizeList
			overscanCount={10}
			height={580}
			itemCount={items.length}
			itemSize={RowHeight}
			width={ReportColumnWidths.reduce((ac, cur) => ac + cur, 0)}
		>
			{({ index, style }) => <ReportRow item={items[index]} sx={style} />}
		</FixedSizeList>
	</TableBody>
)
