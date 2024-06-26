import { SvgIcon, SxProps, Theme } from '@mui/material'
import { FC } from 'react'

export const SortDownIcon: FC<SxProps<Theme>> = style => {
	return (
		<SvgIcon sx={style}>
			<svg
				shapeRendering='geometricPrecision'
				textRendering='geometricPrecision'
				imageRendering='optimizeQuality'
				fillRule='evenodd'
				clipRule='evenodd'
				viewBox='0 0 512 421.65'
			>
				<path
					fillRule='nonzero'
					d='M383.01 0c11.57 0 20.97 9.4 20.97 20.97 0 11.57-9.4 20.97-20.97 20.97l-362.04.44C9.4 42.38 0 32.97 0 21.4 0 9.84 9.4.43 20.97.43L383.01 0zM263.22 309.53c-7.91-7.95-7.86-20.83.08-28.74s20.83-7.86 28.74.08l71.48 71.5.32-206.63c0-11.2 9.09-20.3 20.29-20.3s20.3 9.1 20.3 20.3l-.32 206.4 73.12-73.12c7.94-7.94 20.87-7.94 28.81 0 7.95 7.95 7.95 20.88 0 28.82l-107.9 107.9c-8.02 7.91-20.9 7.87-28.81-.08L263.22 309.53zm-91.6 29.58c11.57 0 20.97 9.4 20.97 20.97 0 11.57-9.4 20.97-20.97 20.97l-150.65.16C9.4 381.21 0 371.81 0 360.24c0-11.56 9.4-20.97 20.97-20.97l150.65-.16zm41.33-170.7c11.57 0 20.97 9.4 20.97 20.97 0 11.57-9.4 20.97-20.97 20.97l-191.98.23C9.4 210.58 0 201.17 0 189.6c0-11.56 9.4-20.97 20.97-20.97l191.98-.22z'
				/>
			</svg>
		</SvgIcon>
	)
}
