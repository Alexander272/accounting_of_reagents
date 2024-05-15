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
	isShow?: boolean
	allowSearch?: boolean
	formatter?: (value: unknown) => string
}

export interface IContextMenu {
	active: string
	coords: ICoordinates
}

export interface ICoordinates {
	mouseX: number
	mouseY: number
}
