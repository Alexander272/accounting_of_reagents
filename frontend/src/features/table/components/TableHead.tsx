import { TableCell } from '@/components/table/TableCell'
import { TableHead } from '@/components/table/TableHead'
import { TableRow } from '@/components/table/TableRow'
import { Titles } from '../constants/titles'

export const DataTableHead = () => {
	return (
		<TableHead>
			<TableRow>
				<TableCell rowSpan={2}>{Titles.Type}</TableCell>
				<TableCell rowSpan={2}>{Titles.Name}</TableCell>
				<TableCell rowSpan={2}>{Titles.UName}</TableCell>
				<TableCell rowSpan={2}>{Titles.Doc}</TableCell>
				<TableCell rowSpan={2}>{Titles.Purity}</TableCell>
				<TableCell rowSpan={2}>{Titles.DateOfManufacture}</TableCell>
				<TableCell rowSpan={2}>{Titles.Consignment}</TableCell>
				<TableCell rowSpan={2}>{Titles.Manufacturer}</TableCell>
				<TableCell rowSpan={2}>{Titles.ShelfLife}</TableCell>
				<TableCell colSpan={2}>{Titles.Place.Main}</TableCell>
				<TableCell></TableCell>
			</TableRow>
			<TableRow>
				<TableCell>{Titles.Place.Closet}</TableCell>
				<TableCell>{Titles.Place.Shelf}</TableCell>
			</TableRow>
		</TableHead>
	)
}
