import { FC } from 'react'
import { SvgIcon, SxProps, Theme } from '@mui/material'

export const FileIcon: FC<SxProps<Theme>> = style => {
	return (
		<SvgIcon sx={style}>
			<svg
				xmlns='http://www.w3.org/2000/svg'
				shapeRendering='geometricPrecision'
				textRendering='geometricPrecision'
				imageRendering='optimizeQuality'
				fillRule='evenodd'
				clipRule='evenodd'
				viewBox='0 0 396 511.464'
			>
				<path
					fillRule='nonzero'
					d='M66.392 0h195.691a9.85 9.85 0 017.86 3.907l123.399 133.496a9.793 9.793 0 012.604 6.666l.054 301.01c0 18.215-7.494 34.803-19.511 46.834-12.071 12.065-28.666 19.551-46.881 19.551H66.392c-18.249 0-34.864-7.473-46.895-19.504C7.48 479.943 0 463.342 0 445.079V66.385c0-18.269 7.473-34.877 19.49-46.894C31.514 7.467 48.129 0 66.392 0zm46.929 237.005c-5.445 0-9.867-4.422-9.867-9.867 0-5.446 4.422-9.868 9.867-9.868h169.358c5.445 0 9.867 4.422 9.867 9.868 0 5.445-4.422 9.867-9.867 9.867H113.321zm0 81.359c-5.445 0-9.867-4.422-9.867-9.867 0-5.446 4.422-9.867 9.867-9.867h160.386c5.445 0 9.867 4.421 9.867 9.867 0 5.445-4.422 9.867-9.867 9.867H113.321zm0 87.117c-5.445 0-9.867-4.422-9.867-9.867 0-5.446 4.422-9.868 9.867-9.868h136.012c5.446 0 9.868 4.422 9.868 9.868 0 5.445-4.422 9.867-9.868 9.867H113.321zM264.884 27.364V99.5c1.16 15.897 6.748 28.429 16.622 36.872 10.016 8.565 24.759 13.346 44.128 13.679l50.631-.041v-2.15L264.884 27.364zm111.381 142.38l-50.787-.04c-24.251-.38-43.24-6.809-56.762-18.372-14.032-11.99-21.919-29.133-23.472-50.442l-.095-1.349V19.735H66.392c-12.817 0-24.482 5.249-32.945 13.712-8.45 8.45-13.712 20.121-13.712 32.938v378.694c0 12.804 5.262 24.468 13.719 32.925 8.463 8.463 20.135 13.726 32.938 13.726h263.216c12.803 0 24.475-5.263 32.945-13.713 8.45-8.47 13.712-20.141 13.712-32.938V169.744z'
				/>
			</svg>
		</SvgIcon>
	)
}