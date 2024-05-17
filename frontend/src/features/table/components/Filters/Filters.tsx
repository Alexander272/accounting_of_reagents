import { useRef, useState } from 'react'
import { Button, Stack, Typography, useTheme } from '@mui/material'

import type { IFilter } from '../../types/data'
import { useAppSelector } from '@/hooks/redux'
import { Popover } from '@/components/Popover/Popover'
import { Badge } from '@/components/Badge/Badge'
import { FilterIcon } from '@/components/Icons/FilterIcon'
import { getFilters } from '../../tableSlice'
import { Columns } from '../../columns'

const defaultValues: IFilter = {
	field: Columns[0].key,
	fieldType: Columns[0].filter as 'string',
	compareType: 'con',
	value: '',
}

export const Filters = () => {
	const [open, setOpen] = useState(false)
	const anchor = useRef(null)

	const filters = useAppSelector(getFilters)

	const { palette } = useTheme()

	const toggleHandler = () => setOpen(prev => !prev)

	const closeHandler = () => {
		toggleHandler()
	}

	return (
		<>
			<Button
				ref={anchor}
				onClick={toggleHandler}
				variant='outlined'
				color='gray'
				sx={{ minWidth: 30, paddingX: 1.5 }}
			>
				<Badge color='primary' variant={filters.length < 2 ? 'dot' : 'standard'} badgeContent={filters.length}>
					<FilterIcon fill={palette.gray.main} fontSize={16} mr={1} />
				</Badge>
				Фильтр
			</Button>

			<Popover
				open={open}
				onClose={closeHandler}
				anchorEl={anchor.current}
				paperSx={{
					maxWidth: 700,
					'&:before': {
						content: '""',
						display: 'block',
						position: 'absolute',
						top: 0,
						left: '66%',
						width: 10,
						height: 10,
						bgcolor: 'background.paper',
						transform: 'translate(-50%, -50%) rotate(45deg)',
						zIndex: 0,
					},
				}}
			>
				<Stack direction={'row'}>
					<Typography fontSize={'1.1rem'}>Применить фильтр</Typography>
				</Stack>

				{filters.map((f, i) => (
					<Stack key={f.field + i} direction={'row'} spacing={2}>
						{/* <Select></Select> */}
					</Stack>
				))}
			</Popover>
		</>
	)
}
