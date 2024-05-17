import { useEffect } from 'react'
import { Button, Stack, TextField, Typography } from '@mui/material'
import { useForm } from 'react-hook-form'
import { toast } from 'react-toastify'

import type { IFetchError } from '@/app/types/error'
import type { CreateNote, NoteForm } from '@/features/table/modules/Note/notes'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { getContextMenu, setContextMenu } from '@/features/table/tableSlice'
import { useGetByIdQuery } from '@/features/table/tableApiSlice'
import {
	useCreateNotesMutation,
	useGetNotesQuery,
	useUpdateNotesMutation,
} from '@/features/table/modules/Note/noteApiSlice'
import { Dialog } from '@/components/Dialog/Dialog'
import { Fallback } from '@/components/Fallback/Fallback'
import { changeModalIsOpen, getModalState } from '../modalSlice'
import { Titles } from '../titles'

export const Note = () => {
	const modal = useAppSelector(getModalState('notes'))

	const dispatch = useAppDispatch()

	const closeHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'notes', isOpen: false }))
		dispatch(setContextMenu())
	}

	return (
		<Dialog
			title={Titles.Notes}
			body={<NotesForm />}
			open={modal.isOpen}
			onClose={closeHandler}
			maxWidth='sm'
			fullWidth
		/>
	)
}

const defaultValues: NoteForm = {
	comment: '',
	note: '',
}

const NotesForm = () => {
	const contextMenu = useAppSelector(getContextMenu)
	const dispatch = useAppDispatch()

	const { data: reagent, isLoading: isLoadingReagent } = useGetByIdQuery(contextMenu?.active || '', {
		skip: !contextMenu?.active,
	})
	const { data, isLoading } = useGetNotesQuery(reagent?.data.id || '', { skip: !reagent?.data.id })

	const [create] = useCreateNotesMutation()
	const [update] = useUpdateNotesMutation()

	const methods = useForm<NoteForm>({ values: defaultValues })

	useEffect(() => {
		if (data?.data.length) methods.reset({ comment: data.data[0].comment, note: data.data[0].note })
	}, [data, methods])

	const closeHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'notes', isOpen: false }))
		dispatch(setContextMenu())
	}

	const saveHandler = async (form: NoteForm) => {
		if (!reagent || !data) return

		const note: CreateNote = {
			reagentId: reagent.data.id,
			comment: form.comment,
			note: form.note,
		}

		try {
			if (data.data.length && data.data[0].id) await update({ ...note, id: data.data[0].id }).unwrap()
			else await create(note).unwrap()
			closeHandler()
		} catch (error) {
			console.log(error)
			const fetchError = error as IFetchError
			toast.error(fetchError.data.message, { autoClose: false })
		}
	}

	return (
		<Stack mt={-3} component={'form'} onSubmit={methods.handleSubmit(saveHandler)}>
			{isLoading || isLoadingReagent ? <Fallback marginTop={5} marginBottom={3} /> : null}

			<Stack mb={2}>
				<Typography fontSize={'1.4rem'} textAlign={'center'}>
					{reagent?.data.name}
				</Typography>
			</Stack>

			<TextField {...methods.register('comment')} label={'Комментарии '} multiline minRows={4} sx={{ mb: 1.5 }} />
			<TextField {...methods.register('note')} label={'Примечание '} multiline minRows={4} />

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
