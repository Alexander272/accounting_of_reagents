import { RootState } from '@/app/store'
import { PayloadAction, createSlice } from '@reduxjs/toolkit'

export type ModalVariants = 'create' | 'edit' | 'amountType' | 'spending' | 'extending'

interface IModalOptions {
	isOpen: boolean
	// content?: string
}

type IModalState = Record<ModalVariants, IModalOptions>

interface IChangeModalAction extends IModalOptions {
	variant: ModalVariants
}

const initialState: IModalState = {
	create: { isOpen: false },
	edit: { isOpen: false },
	amountType: { isOpen: false },
	spending: { isOpen: false },
	extending: { isOpen: false },
}

const modalSlice = createSlice({
	name: 'modal',
	initialState,
	reducers: {
		changeModalIsOpen: (state, action: PayloadAction<IChangeModalAction>) => {
			const { variant, isOpen } = action.payload
			state[variant].isOpen = isOpen
		},

		resetModal: () => initialState,
	},
})

export const getModalState = (variant: ModalVariants) => (state: RootState) => state.modal[variant]

export const modalPath = modalSlice.name
export const modalReducer = modalSlice.reducer

export const { changeModalIsOpen, resetModal } = modalSlice.actions
