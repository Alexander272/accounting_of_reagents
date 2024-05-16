import { FC } from 'react'
import { SxProps, Theme, Box } from '@mui/material'

import StyledGridOverlay from '../StyledGridOverlay'

import EmptyIcon from './EmptyIcon'

export const NoRowsOverlay: FC<SxProps<Theme>> = style => {
	return (
		<StyledGridOverlay sx={style}>
			<EmptyIcon />
			<Box mt={1} color='text.secondary'>
				Не найдено ни одной позиции
			</Box>
		</StyledGridOverlay>
	)
}
