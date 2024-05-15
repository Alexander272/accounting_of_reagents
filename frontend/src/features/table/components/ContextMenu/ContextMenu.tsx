import { ListItemIcon, Menu, MenuItem } from '@mui/material'

import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { ModalVariants, changeModalIsOpen } from '@/features/modal/modalSlice'
import { Update } from '@/features/modal/components/Update'
import { Spending } from '@/features/modal/components/Spending'
import { EditIcon } from '@/components/Icons/EditIcon'
import { FileIcon } from '@/components/Icons/FileIcon'
import { getContextMenu, setContextMenu } from '../../tableSlice'

export const ContextMenu = () => {
	const contextMenu = useAppSelector(getContextMenu)
	const dispatch = useAppDispatch()

	const closeHandler = () => {
		dispatch(setContextMenu())
	}

	const openFormHandler = (variant: ModalVariants) => () => {
		// closeHandler()
		dispatch(changeModalIsOpen({ variant, isOpen: true }))
	}

	return (
		<>
			<Menu
				open={Boolean(contextMenu)}
				onClose={closeHandler}
				anchorReference='anchorPosition'
				anchorPosition={
					contextMenu ? { top: contextMenu.coords.mouseY, left: contextMenu.coords.mouseX } : undefined
				}
			>
				<MenuItem onClick={openFormHandler('spending')}>
					<ListItemIcon>
						<FileIcon fontSize={20} fill={'#363636'} />
					</ListItemIcon>
					Расход
				</MenuItem>
				<MenuItem onClick={openFormHandler('edit')}>
					<ListItemIcon>
						<EditIcon fontSize={18} fill={'#363636'} />
					</ListItemIcon>
					Редактировать
				</MenuItem>

				{/* Продление срока годности */}
				{/* Сведения об изъятии */}
				{/* Сведения об утилизации */}
				{/* Комментарии  */}
				{/* Сформировать заказ? */}
			</Menu>

			<Update />
			<Spending />
		</>
	)
}
