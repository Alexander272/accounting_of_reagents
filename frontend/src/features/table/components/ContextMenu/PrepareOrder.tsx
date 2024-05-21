import { ListItemIcon, MenuItem } from '@mui/material'
import { toast } from 'react-toastify'

import type { IFetchError } from '@/app/types/error'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { usePrepareOrderMutation } from '../../tableApiSlice'
import { getSelected, setContextMenu } from '../../tableSlice'

export const PrepareOrder = () => {
	const selected = useAppSelector(getSelected)
	const dispatch = useAppDispatch()

	const [prepareOrder] = usePrepareOrderMutation()

	const orderHandler = async () => {
		const list = Object.keys(selected)

		try {
			await prepareOrder({ list: list }).unwrap()
			toast.success('Заказ сформирован')
		} catch (error) {
			const fetchError = error as IFetchError
			toast.error(fetchError.data.message, { autoClose: false })
		} finally {
			dispatch(setContextMenu())
		}
	}

	return (
		<MenuItem onClick={orderHandler}>
			<ListItemIcon></ListItemIcon>
			Сформировать заказ
		</MenuItem>
	)
}
