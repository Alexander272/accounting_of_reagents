import type { IReportPeriod, ISpendingReportItem } from '../types'
import { ReportHeaders } from '../constants/table'
import { csvCell } from './format'
import { saveAs } from '@/utils/saveAs'

const withUnit = (value: number, unit: string) => `${value}${unit ? ` ${unit}` : ''}`

export const downloadReportCsv = (items: ISpendingReportItem[], period: IReportPeriod | null) => {
	if (!items.length || !period) return

	const rows = [
		[...ReportHeaders],
		...items.map(item => [
			item.name,
			item.uname,
			item.consignment,
			withUnit(item.amountAtStart, item.unit),
			withUnit(item.spent, item.unit),
			withUnit(item.remainder, item.unit),
		]),
	]

	const content = rows.map(row => row.map(csvCell).join(';')).join('\r\n')
	const blob = new Blob([`\uFEFF${content}`], { type: 'text/csv;charset=utf-8;' })

	saveAs(blob, `расход-реактивов_${period.from.format('YYYY-MM-DD')}_${period.to.format('YYYY-MM-DD')}.csv`)
}
