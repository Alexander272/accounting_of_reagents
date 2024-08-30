import { useAppSelector } from '@/hooks/redux'
import { useGetAllQuery } from '../tableApiSlice'
import { getFilters, getSearch, getTablePage, getTableSize, getTableSort } from '../tableSlice'

export const useGetAllData = () => {
	const page = useAppSelector(getTablePage)
	const size = useAppSelector(getTableSize)

	const search = useAppSelector(getSearch)
	const sort = useAppSelector(getTableSort)
	const filters = useAppSelector(getFilters)

	const query = useGetAllQuery(
		{ page, size, search, sort, filters },
		{ pollingInterval: 5 * 60000, skipPollingIfUnfocused: true, refetchOnFocus: true }
	)

	return query
}
