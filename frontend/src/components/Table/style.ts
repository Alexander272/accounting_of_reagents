import type { CSSProperties } from 'react'
import styled from '@emotion/styled'
import { lighten } from '@mui/material'

import type { TableCellAlign } from './types'

type RowProps = {
	width?: number
	height?: number
	hover?: boolean
	styles?: CSSProperties
}
export const TableRowContainer = styled.div<RowProps>(
	{
		display: 'flex',
		cursor: 'pointer',
		borderRadius: '8px',
		// padding: '0 6px',
		transition: '.3s all ease-in-out',
	},
	props => {
		const hover = {
			background: props.styles?.background ? lighten(props.styles?.background.toString(), 0.4) : '#e6f4f4',
		}

		return {
			width: props.width && props.width + 'px',
			height: props.height && props.height + 'px',

			...props.styles,

			':hover': props.hover ? hover : undefined,
		}
	},
)

export const TableGroupContainer = styled.div`
	position: relative;

	&:not(:last-of-type)::after {
		content: '';
		width: 1px;
		height: 60%;
		background: #e0e0e0;
		position: absolute;
		top: 20%;
		right: -0.5px;
	}
`

type CellProps = {
	width?: number
	isActive?: boolean
	align?: TableCellAlign
}
const justifyContent: Record<TableCellAlign, string> = {
	left: 'flex-start',
	center: 'center',
	right: 'flex-end',
}
export const TableCellContainer = styled.div<CellProps>`
	width: ${props => (props.width ? props.width + 'px' : '100%')};
	position: relative;
	border-bottom: 1px solid #e0e0e0;
	cursor: ${props => props.isActive && 'pointer'};
	transition: all 0.4s ease-in-out;
	padding: 0 8px;
	display: flex;
	align-items: center;
	justify-content: ${props => justifyContent[props.align ?? 'center']};
	text-align: ${props => props.align ?? 'center'};

	&:hover {
		background-color: ${props => props.isActive && '#e0e0e04d'};
		border-radius: ${props => props.isActive && '12px'};
	}

	&:not(:last-of-type)::after {
		content: '';
		width: 1px;
		height: 60%;
		background: #e0e0e0;
		position: absolute;
		top: 20%;
		right: -0.5px;
	}
`

type HeadProps = {
	styles?: CSSProperties
}
export const TableHeadContainer = styled.div<HeadProps>(
	{
		background: '#f5f5f5',
		borderRadius: '12px',
		width: 'fit-content',
	},
	props => {
		return {
			...props.styles,
		}
	},
)

export const TableBodyContainer = styled.div``

type TableProps = {
	columnWidth?: number | number[]
	styles?: CSSProperties
}
export const TableContainer = styled.div<TableProps>`
	max-width: 100%;
	height: 100%;
	overflow-y: hidden;
	overflow-x: auto;
	position: relative;
	margin-top: 8px;

	${props => ({ ...props.styles })}
`
