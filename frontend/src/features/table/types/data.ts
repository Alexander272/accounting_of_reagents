import { ISort } from './table'

export interface IDataItem {
	id: string
}

export interface IParams {
	page?: number
	size?: number
	sort: ISort
}
