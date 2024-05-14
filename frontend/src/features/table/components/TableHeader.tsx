import { Button, Stack, useTheme } from '@mui/material'

import { useAppDispatch } from '@/hooks/redux'
import { Create } from '@/features/modal/components/Create'
import { changeModalIsOpen } from '@/features/modal/modalSlice'
import { PlusIcon } from '@/components/Icons/PlusIcon'
import { Search } from './Search/Search'
// import { Setting } from './Setting/Setting'
import { Filters } from './Filters/Filters'

export const TableHeader = () => {
	const { palette } = useTheme()

	const dispatch = useAppDispatch()

	const createHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'create', isOpen: true }))
	}

	return (
		<Stack direction={'row'} alignItems={'center'} justifyContent={'space-between'} mt={1} mb={0.5} mx={2}>
			<Search />

			<Stack direction={'row'} alignItems={'center'} spacing={2}>
				{/* <Setting /> */}
				<Filters />

				<Button onClick={createHandler} variant='outlined'>
					<PlusIcon fontSize={12} mr={1} fill={palette.primary.main} /> Добавить
				</Button>
				<Create />
			</Stack>
		</Stack>
	)
}
