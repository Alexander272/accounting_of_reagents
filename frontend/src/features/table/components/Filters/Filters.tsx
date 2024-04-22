import { useRef, useState } from 'react'
import { Button, Stack, useTheme } from '@mui/material'

import { useAppSelector } from '@/hooks/redux'
import { FilterIcon } from '@/components/Icons/FilterIcon'
import { Popover } from '@/components/Popover/Popover'
import { getFilters } from '../../tableSlice'

export const Filters = () => {
	const [open, setOpen] = useState(false)
	const anchor = useRef(null)

	const filters = useAppSelector(getFilters)

	const { palette } = useTheme()

	const toggleHandler = () => setOpen(prev => !prev)

	const closeHandler = () => {}

	return (
		<>
			<Button
				ref={anchor}
				onClick={toggleHandler}
				variant='outlined'
				color='gray'
				sx={{ minWidth: 30, paddingX: 1.5 }}
			>
				<FilterIcon fill={palette.gray.main} fontSize={16} mr={1} />
				Фильтр
			</Button>

			<Popover open={open} onClose={closeHandler} anchorEl={anchor.current}>
				{filters.map((f, i) => (
					<Stack key={f.field + i} direction={'row'} spacing={2}>
						{/* <Select></Select> */}
					</Stack>
				))}
			</Popover>
		</>
	)
}
