import { SvgIcon, SxProps, Theme } from '@mui/material'
import { FC } from 'react'

export const ArrowGoesIcon: FC<SxProps<Theme>> = style => {
	return (
		<SvgIcon sx={style}>
			<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 122.883 85.208' xmlSpace='preserve'>
				<path
					fillRule='evenodd'
					clipRule='evenodd'
					d='M122.883 28.086 93.668 0l-.004 18.078H41.932v20.019h51.732l.004 18.08 29.215-28.091zM0 57.118l29.215-28.087.002 18.078h51.734v20.019H29.217l-.002 18.08L0 57.118z'
				/>
			</svg>
		</SvgIcon>
	)
}
