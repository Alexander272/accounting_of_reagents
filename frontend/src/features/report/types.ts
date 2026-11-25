import type { Dayjs } from 'dayjs'

export interface IReportParams {
	from: number
	to: number
	realm: string
}

export interface IReportPeriod {
	from: Dayjs
	to: Dayjs
}

export interface ISpendingReportItem {
	name: string
	uname: string
	consignment: string
	amountAtStart: number
	spent: number
	remainder: number
	unit: string
}
