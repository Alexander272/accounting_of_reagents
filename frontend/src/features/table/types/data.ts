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
	//TODO дописать оставшиеся поля
}

export interface IParams {
	page?: number
	size?: number
	sort: ISort
	search?: ISearch
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
