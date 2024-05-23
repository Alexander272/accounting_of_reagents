export type ReagentType = 'main' | 'auxiliary' | 'experimental' | 'standard' | 'titling'

export interface IDataItem {
	id: string
	type: ReagentType
	name: string
	uname: string
	document: string
	purity: string
	dateOfManufacture: number
	consignment: string
	manufacturer: string
	shelfLife: number
	place_closet: string
	place_shelf: number
	incomingControl_receiptDate: number
	incomingControl_amount: number
	incomingControl_date: number
	incomingControl_protocol: string
	incomingControl_result: boolean
	spending: string
	extending_date: number
	extending_period: number
	seizureInformation: string
	disposalInformation: string
	comments: string
	notes: string
	hasRunOut: boolean
	isOverdue: boolean
	itemStyle?: IItemStyles
}
export interface IItemStyles {
	background?: string
	textColor?: string
}

export interface IParams {
	page?: number
	size?: number
	sort: ISort
	search?: ISearch
	filters?: IFilter[]
}

export interface ISort {
	[x: string]: 'DESC' | 'ASC'
}

export type CompareTypes = 'con' | 'start' | 'end' | 'like' | 'in' | 'eq' | 'gte' | 'lte'
export interface IFilter {
	field: string
	fieldType: string
	compareType: CompareTypes
	value: string
}

export interface ISearch {
	value: string
	fields: string[]
}

export interface ICreateDataItem {
	typeId: string
	name: string
	uname: string
	document: string
	purity: string
	dateOfManufacture: number
	consignment: string
	manufacturer: string
	shelfLife: number
	place_closet: string
	place_shelf: number
	receiptDate: number
	amount: number
	amountTypeId: string
	controlDate: number
	protocol: string
	result: boolean
}

export interface IUpdateDataItem extends ICreateDataItem {
	id: string
	seizure: string
	disposal: string
}
