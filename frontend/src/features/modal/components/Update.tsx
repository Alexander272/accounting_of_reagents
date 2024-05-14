import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { changeModalIsOpen, getModalState } from '../modalSlice'
import { Dialog } from '@/components/Dialog/Dialog'
import { Titles } from '../titles'

export const Update = () => {
	const modal = useAppSelector(getModalState('edit'))

	const dispatch = useAppDispatch()

	const closeHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'edit', isOpen: false }))
	}

	return (
		<Dialog title={Titles.Edit} body={<></>} open={modal.isOpen} onClose={closeHandler} maxWidth='md' fullWidth />
	)
}
