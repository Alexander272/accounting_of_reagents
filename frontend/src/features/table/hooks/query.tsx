import { useAppSelector } from '@/hooks/redux'
import { getTablePage, getTableSize, getTableSort } from '../tableSlice'
import { useGetAllQuery } from '../tableApiSlice'

export const useGetAllData = () => {
	const page = useAppSelector(getTablePage)
	const size = useAppSelector(getTableSize)

	const sort = useAppSelector(getTableSort)
	// const filter = useAppSelector(getTableFilter)

	const query = useGetAllQuery({ page, size, sort }, { pollingInterval: 5 * 60000 })

	return query
}
