import styled from '@emotion/styled'

export const TableContainer = styled.div``

export const TableHeadContainer = styled.div``

type RowProps = {
	columnWidth?: number | number[]
}
export const TableRowContainer = styled.div<RowProps>`
	display: grid;
`

export const TableCellContainer = styled.div``
