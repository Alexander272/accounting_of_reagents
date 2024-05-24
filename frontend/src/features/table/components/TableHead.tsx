import { ColWidth, RowHeight } from '@/constants/defaultValues'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { TableCell } from '@/components/Table/TableCell'
import { TableHead } from '@/components/Table/TableHead'
import { TableRow } from '@/components/Table/TableRow'
import { TableGroup } from '@/components/Table/TableGroup'
import { CellText } from '@/components/CellText/CellText'
import { Badge } from '@/components/Badge/Badge'
import { SortUpIcon } from '@/components/Icons/SortUpIcon'
import { SortDownIcon } from '@/components/Icons/SortDownIcon'
import { getHidden, getTableSort, setSort } from '../tableSlice'
import { Columns, HeaderColumns } from '../columns'

export const DataTableHead = () => {
	const sort = useAppSelector(getTableSort)
	const hidden = useAppSelector(getHidden)

	const width = Columns.reduce((ac, cur) => ac + (hidden[cur.key] ? 0 : cur.width || ColWidth), 12)
	const height = 2 * RowHeight

	const dispatch = useAppDispatch()

	const setSortHandler = (field: string) => () => {
		dispatch(setSort(field))
	}

	const renderHeader = () => {
		const header: JSX.Element[] = []

		HeaderColumns.forEach(c => {
			if (c.children) {
				const subhead: JSX.Element[] = []
				c.children.forEach(c => {
					if (!hidden[c.key]) {
						subhead.push(
							<TableCell key={c.key} width={c.width || ColWidth} isActive onClick={setSortHandler(c.key)}>
								<CellText value={c.label} />

								<Badge
									color='primary'
									badgeContent={Object.keys(sort).findIndex(k => k == c.key) + 1}
									invisible={Object.keys(sort).length < 2}
								>
									{!sort[c.key] || sort[c.key] == 'ASC' ? (
										<SortUpIcon fontSize={16} fill={sort[c.key] ? 'black' : '#adadad'} />
									) : null}
									{sort[c.key] == 'DESC' ? (
										<SortDownIcon fontSize={16} fill={sort[c.key] ? 'black' : '#adadad'} />
									) : null}
								</Badge>
							</TableCell>
						)
					}
				})

				if (subhead.length > 0) {
					header.push(
						<TableGroup key={c.key}>
							<TableRow>
								<TableCell key={c.key}>
									<CellText value={c.label} />
								</TableCell>
							</TableRow>
							<TableRow>{subhead}</TableRow>
						</TableGroup>
					)
				}
			} else if (!hidden[c.key]) {
				header.push(
					<TableCell key={c.key} width={c.width || ColWidth} isActive onClick={setSortHandler(c.key)}>
						<CellText value={c.label} />

						<Badge
							color='primary'
							badgeContent={Object.keys(sort).findIndex(k => k == c.key) + 1}
							invisible={Object.keys(sort).length < 2}
						>
							{!sort[c.key] || sort[c.key] == 'ASC' ? (
								<SortUpIcon fontSize={16} fill={sort[c.key] ? 'black' : '#adadad'} />
							) : null}
							{sort[c.key] == 'DESC' ? (
								<SortDownIcon fontSize={16} fill={sort[c.key] ? 'black' : '#adadad'} />
							) : null}
						</Badge>
					</TableCell>
				)
			}
		})

		return header
	}

	return (
		<TableHead>
			<TableRow width={width} height={height} sx={{ padding: '0 6px' }}>
				{renderHeader()}
			</TableRow>
		</TableHead>
	)
}
