import { Table } from '@/components/Table/Table'
import { DataTableHead } from './TableHead'
import { DataTableBody } from './TableBody'
import { TableFooter } from './TableFooter'
import { TableHeader } from './TableHeader'
import { ContextMenu } from './ContextMenu/ContextMenu'

export const DataTable = () => {
	return (
		<>
			<TableHeader />
			<Table>
				<DataTableHead />
				<DataTableBody />
				<ContextMenu />
			</Table>
			<TableFooter />
		</>
	)
}
