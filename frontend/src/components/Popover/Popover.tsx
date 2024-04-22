import { Menu, MenuProps } from '@mui/material'
import { FC, PropsWithChildren } from 'react'

interface Props extends MenuProps {}

export const Popover: FC<PropsWithChildren<Props>> = ({ children, open, onClose, anchorEl, ...props }) => {
	return (
		<Menu
			open={open}
			onClose={onClose}
			anchorEl={anchorEl}
			transformOrigin={{ horizontal: 'center', vertical: 'top' }}
			anchorOrigin={{ horizontal: 'center', vertical: 'bottom' }}
			MenuListProps={{
				role: 'listbox',
				disableListWrap: true,
			}}
			slotProps={{
				paper: {
					elevation: 0,
					sx: {
						overflow: 'visible',
						filter: 'drop-shadow(0px 2px 8px rgba(0,0,0,0.32))',
						mt: 1.5,
						paddingX: 2,
						paddingBottom: 2,
						maxWidth: 400,
						width: '100%',
						'&:before': {
							content: '""',
							display: 'block',
							position: 'absolute',
							top: 0,
							left: '50%',
							width: 10,
							height: 10,
							bgcolor: 'background.paper',
							transform: 'translate(-50%, -50%) rotate(45deg)',
							zIndex: 0,
						},
					},
				},
			}}
			{...props}
		>
			{children}
		</Menu>
	)
}
