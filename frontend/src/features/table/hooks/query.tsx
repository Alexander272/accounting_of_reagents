import { useAppSelector } from '@/hooks/redux'
import { useGetAllQuery } from '../tableApiSlice'
import { getSearch, getTablePage, getTableSize, getTableSort } from '../tableSlice'

export const useGetAllData = () => {
	const page = useAppSelector(getTablePage)
	const size = useAppSelector(getTableSize)

	const search = useAppSelector(getSearch)
	const sort = useAppSelector(getTableSort)
	// const filter = useAppSelector(getTableFilter)

	const query = useGetAllQuery({ page, size, search, sort }, { pollingInterval: 5 * 60000 })

	return query
}
