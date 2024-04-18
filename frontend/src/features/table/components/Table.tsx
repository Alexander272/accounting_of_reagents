import { Table } from '@/components/Table/Table'
import { DataTableHead } from './TableHead'
import { DataTableBody } from './TableBody'

export const DataTable = () => {
	return (
		<>
			<Table>
				<DataTableHead />
				<DataTableBody />
			</Table>
		</>
	)
}
