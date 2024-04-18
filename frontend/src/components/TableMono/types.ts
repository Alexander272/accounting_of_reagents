export interface IColumn {
	name: string
	header: string
	width?: number
	visible?: boolean
	group?: string
	allowSorting?: boolean
	align?: 'center' | 'right' | 'left'
	formatter?: (value: unknown) => string
}

export interface IGroup {
	name: string
	header: string
	children?: IColumn[]
}
