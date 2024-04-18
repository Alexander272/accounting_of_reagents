import { createTheme } from '@mui/material/styles'

import { type IScrollbarParameters, generateScrollbarStyles } from '@/utils/generateScrollbarStyles'

const scrollbarParameters: IScrollbarParameters = {
	borderRadius: '5px',
	scrollbarBgColor: '#f2f2f2',
	scrollbarHeight: '.5rem',
	scrollbarWidth: '.5rem',
	thumbColor: '#00000020',
	thumbColorActive: '#00000050',
	thumbColorHover: '#00000030',
}

export const theme = createTheme({
	palette: {
		primary: {
			main: '#05287f',
			light: '#eaeefc',
		},
		background: {
			default: '#fafafa',
			paper: '#FFF',
		},
	},
	components: {
		MuiCssBaseline: {
			// styleOverrides: defaultStyles.concat(
			// 	// FontFace.generateFontParameters(fontList),
			// 	generateScrollbarStyles(scrollbarParameters)
			// ),
			styleOverrides: generateScrollbarStyles(scrollbarParameters),
		},
		MuiAppBar: {
			styleOverrides: {
				root: {
					border: 'none',
					boxShadow: '0 0 3px #0000004f',
					background: '#fff',
				},
			},
			defaultProps: {
				elevation: 0,
			},
		},
		MuiTooltip: {
			defaultProps: {
				disableInteractive: true,
				arrow: true,
				// enterDelay: 800,
			},
			styleOverrides: {
				tooltip: {
					fontSize: '0.9rem',
					backgroundColor: '#000000de',
				},
				arrow: {
					color: '#000000de',
				},
			},
		},
		MuiPaper: {
			styleOverrides: {
				root: {
					borderRadius: 8,
				},
			},
		},
		MuiSelect: {
			defaultProps: {
				size: 'small',
			},
			styleOverrides: {
				root: {
					borderRadius: 12,
				},
			},
		},
		MuiInputLabel: {
			defaultProps: {
				size: 'small',
			},
		},
		MuiTextField: {
			defaultProps: {
				size: 'small',
				autoComplete: 'off',
			},
			styleOverrides: {
				root: {
					// borderRadius: 12,
					'& fieldset': {
						transition: 'all 0.3s ease-in-out',
					},
				},
			},
		},
		MuiOutlinedInput: {
			styleOverrides: {
				root: {
					borderRadius: 12,
				},
			},
		},
		MuiButton: {
			styleOverrides: {
				root: {
					borderRadius: 12,
				},
			},
		},
		MuiBadge: {
			defaultProps: {
				anchorOrigin: {
					vertical: 'top',
					horizontal: 'left',
				},
			},
		},
	},
})
