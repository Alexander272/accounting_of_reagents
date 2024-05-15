export interface ISpending {
	id: string
	reagentId: string
	date: number
	amount: number
}

export type CreateSpending = Omit<ISpending, 'id'>

export interface ISpendingForm {
	date: number
	amount: string
}
