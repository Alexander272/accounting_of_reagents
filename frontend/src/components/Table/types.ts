export type TableCellAlign = 'left' | 'center' | 'right'

export interface ITableColumn {
	key: string
	label: string
	width?: number
	allowsSorting?: boolean
	children: ITableColumn[]
}
