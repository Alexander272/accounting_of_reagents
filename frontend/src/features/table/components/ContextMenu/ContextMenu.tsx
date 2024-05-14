import { ListItemIcon, Menu, MenuItem } from '@mui/material'

import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { Update } from '@/features/modal/components/Update'
import { EditIcon } from '@/components/Icons/EditIcon'
import { getContextMenu, setContextMenu } from '../../tableSlice'
import { ModalVariants, changeModalIsOpen } from '@/features/modal/modalSlice'

export const ContextMenu = () => {
	const contextMenu = useAppSelector(getContextMenu)
	const dispatch = useAppDispatch()

	const closeHandler = () => {
		dispatch(setContextMenu())
	}

	const openFormHandler = (variant: ModalVariants) => () => {
		closeHandler()
		dispatch(changeModalIsOpen({ variant, isOpen: true }))
	}

	return (
		<>
			<Menu
				open={Boolean(contextMenu)}
				onClose={closeHandler}
				anchorReference='anchorPosition'
				anchorPosition={contextMenu ? { top: contextMenu.mouseY, left: contextMenu.mouseX } : undefined}
			>
				<MenuItem onClick={openFormHandler('edit')}>
					<ListItemIcon>
						<EditIcon fontSize={20} fill={'#363636'} />
					</ListItemIcon>
					Редактировать
				</MenuItem>

				{/* Продление срока годности */}
				{/* Расход */}
				{/* Сведения об изъятии */}
				{/* Сведения об утилизации */}
				{/* Комментарии  */}
				{/* Сформировать заказ? */}
			</Menu>

			<Update />
		</>
	)
}
