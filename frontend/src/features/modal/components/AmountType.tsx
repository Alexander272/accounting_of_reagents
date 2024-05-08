import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { Dialog } from '@/components/Dialog/Dialog'
import { changeModalIsOpen, getModalState } from '../modalSlice'
import { Titles } from '../titles'

export const AmountType = () => {
	const modal = useAppSelector(getModalState('amountType'))

	const dispatch = useAppDispatch()

	const closeHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'amountType', isOpen: false }))
	}

	return (
		<Dialog
			title={Titles.AmountType}
			body={<></>}
			open={modal.isOpen}
			onClose={closeHandler}
			maxWidth='md'
			fullWidth
		/>
	)
}
