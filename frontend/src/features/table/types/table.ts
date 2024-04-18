export interface IHeadColumn {
	key: string
	label: string
	width?: number
	align?: 'center' | 'right' | 'left'
	allowsSorting?: boolean
	children?: IHeadColumn[]
}

export interface IColumn {
	key: string
	label: string
	width?: number
	align?: 'center' | 'right' | 'left'
	formatter?: (value: unknown) => string
}

export interface ISort {
	[x: string]: 'DESC' | 'ASC'
}
