import { FC } from 'react'
import { Button, Stack, Typography } from '@mui/material'
import { toast } from 'react-toastify'

import type { IFetchError } from '@/app/types/error'
import type { IUpdateDataItem } from '@/features/table/types/data'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { getContextMenu, setContextMenu } from '@/features/table/tableSlice'
import { useGetByIdQuery, useUpdateMutation } from '@/features/table/tableApiSlice'
import { useGetReagentTypesQuery } from '@/features/table/modules/Types/typesApiSlice'
import { Dialog } from '@/components/Dialog/Dialog'
import { Fallback } from '@/components/Fallback/Fallback'
import { changeModalIsOpen, getModalState } from '../modalSlice'
import { Titles } from '../titles'

export const Seizure: FC = () => {
	const modal = useAppSelector(getModalState('seizure'))

	const dispatch = useAppDispatch()

	const closeHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'seizure', isOpen: false }))
		dispatch(setContextMenu())
	}

	return (
		<Dialog
			title={Titles.Seizure}
			body={<SeizureForm />}
			open={modal.isOpen}
			onClose={closeHandler}
			maxWidth='sm'
			fullWidth
		/>
	)
}

const SeizureForm: FC = () => {
	const contextMenu = useAppSelector(getContextMenu)
	const dispatch = useAppDispatch()

	const { data: reagent, isLoading } = useGetByIdQuery(contextMenu?.active || '', { skip: !contextMenu?.active })
	const { data: rTypes, isLoading: isLoadingRTypes } = useGetReagentTypesQuery(null)

	const [update] = useUpdateMutation()

	const closeHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'seizure', isOpen: false }))
		dispatch(setContextMenu())
	}

	const saveHandler = async () => {
		const newType = rTypes?.data.find(t => t.name == 'Э')
		if (!reagent || !newType) return

		const newData: IUpdateDataItem = { ...reagent.data, typeId: newType.id, seizure: 'изъят' }
		try {
			await update(newData).unwrap()
			closeHandler()
		} catch (error) {
			const fetchError = error as IFetchError
			toast.error(fetchError.data.message, { autoClose: false })
		}
	}

	return (
		<Stack mt={-3}>
			{isLoading || isLoadingRTypes ? <Fallback marginTop={5} marginBottom={3} /> : null}

			<Typography>Изъять {reagent?.data.name}</Typography>

			<Stack direction={'row'} spacing={3} mt={2}>
				<Button onClick={closeHandler} variant='outlined' fullWidth>
					Отмена
				</Button>
				<Button onClick={saveHandler} variant='contained' fullWidth>
					Да
				</Button>
			</Stack>
		</Stack>
	)
}
