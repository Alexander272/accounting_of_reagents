import { ListItemIcon, Menu, MenuItem } from '@mui/material'

import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { ModalVariants, changeModalIsOpen } from '@/features/modal/modalSlice'
import { Update } from '@/features/modal/components/Update'
import { Spending } from '@/features/modal/components/Spending'
import { Extending } from '@/features/modal/components/Extending'
import { Disposal } from '@/features/modal/components/Disposal'
import { Seizure } from '@/features/modal/components/Seizure'
import { Note } from '@/features/modal/components/Note'
import { EditIcon } from '@/components/Icons/EditIcon'
import { FileIcon } from '@/components/Icons/FileIcon'
import { ClockIcon } from '@/components/Icons/ClockIcon'
import { DisposalIcon } from '@/components/Icons/DisposalIcon'
import { NotebookIcon } from '@/components/Icons/NotebookIcon'
import { ExchangeRefreshIcon } from '@/components/Icons/ExchangeRefreshIcon'
import { getContextMenu, setContextMenu } from '../../tableSlice'
import { PrepareOrder } from './PrepareOrder'
import { CreateOnBase } from './CreateOnBase'

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
				<MenuItem onClick={openFormHandler('notes')}>
					<ListItemIcon>
						<NotebookIcon fontSize={20} fill={'#363636'} />
					</ListItemIcon>
					Комментарии
				</MenuItem>
				<PrepareOrder />
				<CreateOnBase />
				<MenuItem onClick={openFormHandler('edit')}>
					<ListItemIcon>
						<EditIcon fontSize={18} fill={'#363636'} />
					</ListItemIcon>
					Редактировать
				</MenuItem>
				<MenuItem onClick={openFormHandler('extending')}>
					<ListItemIcon>
						<ClockIcon fontSize={18} fill={'#363636'} />
					</ListItemIcon>
					Продление срока годности
				</MenuItem>
				<MenuItem onClick={openFormHandler('seizure')}>
					<ListItemIcon>
						{/* <ChangeExchangeIcon fontSize={18} fill={'#363636'} /> */}
						<ExchangeRefreshIcon fontSize={18} fill={'#363636'} />
					</ListItemIcon>
					Сведения об изъятии
				</MenuItem>
				<MenuItem onClick={openFormHandler('disposal')}>
					<ListItemIcon>
						<DisposalIcon fontSize={18} fill={'#363636'} />
					</ListItemIcon>
					Сведения об утилизации
				</MenuItem>

				{/* Сформировать заказ? */}
			</Menu>

			<Update />
			<Spending />
			<Extending />
			<Disposal />
			<Note />
			<Seizure />
		</>
	)
}
