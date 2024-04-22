import { Button, Stack, useTheme } from '@mui/material'

import { Search } from './Search/Search'
import { Setting } from './Setting/Setting'
import { Filters } from './Filters/Filters'
import { PlusIcon } from '@/components/Icons/PlusIcon'

export const TableHeader = () => {
	const { palette } = useTheme()

	return (
		<Stack direction={'row'} alignItems={'center'} justifyContent={'space-between'} mt={1} mb={0.5} mx={2}>
			<Search />

			<Stack direction={'row'} alignItems={'center'} spacing={2}>
				<Setting />
				<Filters />

				<Button variant='outlined'>
					<PlusIcon fontSize={12} mr={1} fill={palette.primary.main} /> Добавить
				</Button>
			</Stack>
		</Stack>
	)
}
