import styled from '@emotion/styled'
import { alpha, Button, styled as muiStyled } from '@mui/material'
import { Link } from 'react-router-dom'

export const NavLink = styled(Link)`
	text-decoration: none;
	font-size: 1.2rem;
	color: var(--primary-color);
	height: 100%;
	display: flex;
	align-items: center;
	justify-content: center;
	position: relative;

	&:after {
		content: '';
		position: absolute;
		bottom: 0px;
		height: 2px;
		width: 100%;
		max-width: 0%;
		background-color: var(--primary-color);
		transition: all 0.3s ease-in-out;
	}

	&:hover {
		&:after {
			max-width: 100%;
		}
	}
`

export const NavButton = muiStyled(Button)(({ theme }) => ({
	cursor: 'pointer',
	height: '48px',
	display: 'flex',
	alignItems: 'center',
	justifyContent: 'center',
	position: 'relative',
	borderRadius: 8,
	transition: 'all 0.15s ease-in-out',

	color: theme.palette.text.primary,
	textTransform: 'none',
	fontSize: 15,
	fontWeight: 500,
	padding: theme.spacing(0, 1.5),
	'& .MuiButton-startIcon': { marginRight: theme.spacing(0.75) },
	'&:hover': { backgroundColor: alpha(theme.palette.primary.main, 0.06) },
	'&.active': {
		height: '100%',
		borderRadius: 0,
		color: theme.palette.primary.main,
		backgroundColor: 'transparent', // убираем заливку
		boxShadow: `inset 0 -2px 0 ${theme.palette.primary.main}`,
	},
})) as typeof Button

export const NavButtonMy = muiStyled('div')(({ theme }) => ({
	cursor: 'pointer',
	fontSize: '1.2rem',
	color: theme.palette.primary.main,
	height: '100%',
	display: 'flex',
	alignItems: 'center',
	justifyContent: 'center',
	position: 'relative',

	'&:after': {
		content: '""',
		position: 'absolute',
		bottom: 0,
		height: 2,
		width: '100%',
		maxWidth: '0%',
		backgroundColor: theme.palette.primary.main,
		transition: 'all 0.3s ease-in-out',
	},

	'&:hover': {
		'&:after': {
			maxWidth: '100%',
		},
	},
}))

// export const NavButton = styled.p`
// 	cursor: pointer;
// 	font-size: 1.2rem;
// 	color: var(--primary-color);
// 	height: 100%;
// 	display: flex;
// 	align-items: center;
// 	justify-content: center;
// 	position: relative;

// 	&:after {
// 		content: '';
// 		position: absolute;
// 		bottom: 0px;
// 		height: 2px;
// 		width: 100%;
// 		max-width: 0%;
// 		background-color: var(--primary-color);
// 		transition: all 0.3s ease-in-out;
// 	}

// 	&:hover {
// 		&:after {
// 			max-width: 100%;
// 		}
// 	}
// `
