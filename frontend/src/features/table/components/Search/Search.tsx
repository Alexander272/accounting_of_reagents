import { type ChangeEvent, type FC } from 'react'
import { InputAdornment, TextField } from '@mui/material'

import { useAppDispatch } from '@/hooks/redux'
import { useDebounce } from '@/hooks/useDebounce'
import { SearchIcon } from '@/components/Icons/SearchIcon'
import { setSearch } from '@/features/table/tableSlice'
import { Setting } from './Setting'

export const Search: FC = () => {
	const dispatch = useAppDispatch()

	const searchHandler = useDebounce(v => {
		dispatch(setSearch(v as string))
	}, 700)

	const changeValueHandler = (event: ChangeEvent<HTMLInputElement>) => {
		// setValue(event.target.value)
		searchHandler(event.target.value)
	}

	return (
		<>
			<TextField
				// value={value}
				onChange={changeValueHandler}
				InputProps={{
					startAdornment: (
						<InputAdornment position='start'>
							<SearchIcon fontSize={16} />
						</InputAdornment>
					),
					endAdornment: (
						<InputAdornment position='end'>
							<Setting />
						</InputAdornment>
					),
				}}
				placeholder='Поиск...'
				sx={{ width: 350 }}
			/>
		</>
	)
}
