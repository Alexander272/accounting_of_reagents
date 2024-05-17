import { useEffect, useRef } from 'react'
import { Button, Stack, TextField, Typography } from '@mui/material'
import { toast } from 'react-toastify'

import type { IFetchError } from '@/app/types/error'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { useGetByIdQuery, useUpdateMutation } from '@/features/table/tableApiSlice'
import { getContextMenu, setContextMenu } from '@/features/table/tableSlice'
import { Dialog } from '@/components/Dialog/Dialog'
import { Fallback } from '@/components/Fallback/Fallback'
import { changeModalIsOpen, getModalState } from '../modalSlice'
import { Titles } from '../titles'

export const Disposal = () => {
	const modal = useAppSelector(getModalState('disposal'))

	const dispatch = useAppDispatch()

	const closeHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'disposal', isOpen: false }))
		dispatch(setContextMenu())
	}

	return (
		<Dialog
			title={Titles.Disposal}
			body={<DisposalForm />}
			open={modal.isOpen}
			onClose={closeHandler}
			maxWidth='sm'
			fullWidth
		/>
	)
}

const DisposalForm = () => {
	const input = useRef<HTMLInputElement>()
	const contextMenu = useAppSelector(getContextMenu)
	const dispatch = useAppDispatch()

	const { data: reagent, isLoading } = useGetByIdQuery(contextMenu?.active || '', { skip: !contextMenu?.active })

	const [update] = useUpdateMutation()

	useEffect(() => {
		if (reagent && input.current) input.current.value = reagent.data.disposal
	}, [reagent])

	const closeHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'disposal', isOpen: false }))
		dispatch(setContextMenu())
	}

	const saveHandler = async () => {
		console.log(input.current?.value)
		if (!reagent || !input.current) return
		const newData = { ...reagent.data, disposal: input.current.value }

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
			{isLoading ? <Fallback marginTop={5} marginBottom={3} /> : null}

			<Stack mb={2}>
				<Typography fontSize={'1.4rem'} textAlign={'center'}>
					{reagent?.data.name}
				</Typography>
			</Stack>

			<TextField inputRef={input} label={'Сведения об утилизации'} multiline minRows={4} />

			<Stack direction={'row'} spacing={3} mt={2}>
				<Button onClick={closeHandler} variant='outlined' fullWidth>
					Отмена
				</Button>
				<Button onClick={saveHandler} variant='contained' fullWidth>
					Сохранить
				</Button>
			</Stack>
		</Stack>
	)
}
