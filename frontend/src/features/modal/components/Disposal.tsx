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
import { Controller, useForm } from 'react-hook-form'

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

type DisposalForm = { disposal: string }

const DisposalForm = () => {
	const contextMenu = useAppSelector(getContextMenu)
	const dispatch = useAppDispatch()

	const { data: reagent, isLoading } = useGetByIdQuery(contextMenu?.active || '', { skip: !contextMenu?.active })

	const [update] = useUpdateMutation()

	const methods = useForm<DisposalForm>({ values: { disposal: reagent?.data.disposal || '' } })

	const closeHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'disposal', isOpen: false }))
		dispatch(setContextMenu())
	}

	const saveHandler = async (form: DisposalForm) => {
		console.log(form)
		if (!reagent || !form) return
		const newData = { ...reagent.data, disposal: form.disposal }

		try {
			await update(newData).unwrap()
			closeHandler()
		} catch (error) {
			const fetchError = error as IFetchError
			toast.error(fetchError.data.message, { autoClose: false })
		}
	}

	return (
		<Stack mt={-3} component={'form'} onSubmit={methods.handleSubmit(saveHandler)}>
			{isLoading ? <Fallback marginTop={5} marginBottom={3} /> : null}

			<Stack mb={2}>
				<Typography fontSize={'1.4rem'} textAlign={'center'}>
					{reagent?.data.name}
				</Typography>
			</Stack>

			<Controller
				control={methods.control}
				name='disposal'
				rules={{ required: true }}
				render={({ field }) => <TextField {...field} label={'Сведения об утилизации'} multiline minRows={4} />}
			/>
			{/* <TextField inputRef={input} label={'Сведения об утилизации'} multiline minRows={4} /> */}

			<Stack direction={'row'} spacing={3} mt={2}>
				<Button onClick={closeHandler} variant='outlined' fullWidth>
					Отмена
				</Button>
				<Button type='submit' variant='contained' fullWidth>
					Сохранить
				</Button>
			</Stack>
		</Stack>
	)
}
