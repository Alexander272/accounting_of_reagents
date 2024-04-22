export interface IDataItem {
	id: string
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
