export interface IHeadColumn {
	key: string
	label: string
	width?: number
	align?: 'center' | 'right' | 'left'
	allowsSorting?: boolean
	children?: IHeadColumn[]
}

export type FilterType = 'number' | 'string' | 'date' | 'switch' | 'list'
export interface IFullFilter {
	type: FilterType
	options?: unknown
	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	getOptions?: (arg: null) => any
}
export interface IColumn {
	key: string
	label: string
	width?: number
	align?: 'center' | 'right' | 'left'
	isShow?: boolean
	allowSearch?: boolean
	filter?: FilterType | IFullFilter
	formatter?: (value: unknown) => string
}

import type { IDataItem } from './data'

export interface IContextMenu {
	active: IDataItem
	coords: ICoordinates
}

export interface ICoordinates {
	mouseX: number
	mouseY: number
}

export interface ISelect {
	[id: string]: boolean
}
