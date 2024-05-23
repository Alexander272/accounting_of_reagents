import { FC } from 'react'
import { CircularProgress, ListItemIcon, MenuItem } from '@mui/material'
import { toast } from 'react-toastify'

import type { IFetchError } from '@/app/types/error'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { changeModalIsOpen } from '@/features/modal/modalSlice'
import { CopyIcon } from '@/components/Icons/CopyIcon'
import { getContextMenu, setContextMenu } from '../../tableSlice'
import { useLazyGetByIdQuery } from '../../tableApiSlice'
import { localKeys } from '../../constants/localKeys'

export const CreateOnBase: FC = () => {
	const contextMenu = useAppSelector(getContextMenu)
	const dispatch = useAppDispatch()

	const [get, { isLoading }] = useLazyGetByIdQuery()

	const createHandler = async () => {
		if (!contextMenu?.active) return

		try {
			const data = await get(contextMenu?.active).unwrap()
			const base = {
				type: data.data.typeId,
				name: data.data.name,
				uname: data.data.uname,
				document: data.data.document,
				purity: data.data.purity,
				dateOfManufacture: data.data.dateOfManufacture,
				consignment: data.data.consignment,
				manufacturer: data.data.manufacturer,
				shelfLife: data.data.shelfLife.toString(),
				place_closet: data.data.place_closet,
				place_shelf: data.data.place_shelf.toString(),
			}
			const control = {
				receiptDate: data.data.receiptDate,
				amount: data.data.amount.toString(),
				amountType: data.data.amountTypeId,
				controlDate: data.data.controlDate,
				protocol: data.data.protocol,
				result: data.data.result,
			}
			localStorage.setItem(localKeys.baseInfo, JSON.stringify(base))
			localStorage.setItem(localKeys.control, JSON.stringify(control))
			dispatch(changeModalIsOpen({ variant: 'create', isOpen: true }))
			dispatch(setContextMenu())
		} catch (error) {
			const fetchError = error as IFetchError
			toast.error(fetchError.data.message, { autoClose: false })
		}
	}

	return (
		<MenuItem onClick={createHandler}>
			<ListItemIcon>
				{isLoading ? <CircularProgress size={18} /> : <CopyIcon fontSize={18} fill={'#363636'} />}
			</ListItemIcon>
			Создать на основании
		</MenuItem>
	)
}
