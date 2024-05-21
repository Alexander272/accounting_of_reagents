import { SvgIcon, SxProps, Theme } from '@mui/material'
import { FC } from 'react'

export const ExchangeRefreshIcon: FC<SxProps<Theme>> = style => {
	return (
		<SvgIcon sx={style}>
			<svg
				xmlns='http://www.w3.org/2000/svg'
				viewBox='0 0 102.52 122.88'
				enableBackground={'new 0 0 102.52 122.88'}
				xmlSpace='preserve'
			>
				<path
					fillRule='evenodd'
					clipRule='evenodd'
					d='m56.9 9.2-8.75 44.47-8.87-13.37C20.17 47.97 9.44 60.62 7.85 80.09c-15.7-27.44-6.16-52.04 13.73-66.45L12.52 0 56.9 9.2zM45.61 113.68l8.75-44.47 8.87 13.37c19.11-7.67 29.83-20.32 31.43-39.79 15.7 27.44 6.16 52.04-13.73 66.45l9.05 13.64-44.37-9.2z'
				/>
			</svg>
		</SvgIcon>
	)
}
