import { FixedSizeList } from 'react-window'

import { ColWidth, RowHeight, Size } from '@/constants/defaultValues'
import { useAppSelector } from '@/hooks/redux'
import { Fallback } from '@/components/Fallback/Fallback'
import { TableBody } from '@/components/Table/TableBody'
import { useGetAllData } from '../hooks/query'
import { getTableSize } from '../tableSlice'
import { Columns } from '../columns'
import { NoRowsOverlay } from './NoRowsOverlay/components/NoRowsOverlay'
import { DataTableRow } from './TableRow'

export const DataTableBody = () => {
	const size = useAppSelector(getTableSize)

	const { data, isFetching } = useGetAllData()

	//TODO NoRowsOverlay занимает только часть экрана и при скроле не двигается
	if (!data || !data.total) return <NoRowsOverlay />

	return (
		<TableBody>
			{isFetching && (
				<Fallback
					position={'absolute'}
					top={'50%'}
					left={'50%'}
					transform={'translate(-50%, -50%)'}
					height={160}
					width={160}
					borderRadius={3}
					zIndex={15}
					backgroundColor={'#fafafa'}
				/>
			)}

			<FixedSizeList
				overscanCount={10}
				height={RowHeight * Size}
				itemCount={data.data.length > (size || Size) ? size || Size : data.data.length}
				itemSize={RowHeight}
				width={Columns.reduce((ac, cur) => ac + (cur.width || ColWidth), 12)}
			>
				{({ index, style }) => <DataTableRow data={data.data[index]} sx={style} />}
			</FixedSizeList>
		</TableBody>
	)
}
