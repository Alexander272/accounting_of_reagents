export interface IExtending {
	id: string
	reagentId: string
	date: number
	period: number
}

export type CreateExtending = Omit<IExtending, 'id'>

export interface IExtendingForm {
	date: number
	period: string
}
