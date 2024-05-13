import { FC, useCallback } from 'react'
import { Button, Divider, IconButton, Stack, TextField, useTheme } from '@mui/material'
import { FieldErrors, UseFormRegister, useFieldArray, useForm } from 'react-hook-form'
import { DndContext, DragEndEvent, PointerSensor, closestCenter, useSensor, useSensors } from '@dnd-kit/core'
import { SortableContext, useSortable, verticalListSortingStrategy } from '@dnd-kit/sortable'
import { CSS } from '@dnd-kit/utilities'
import { toast } from 'react-toastify'

import type { IFetchError } from '@/app/types/error'
import type { IAmountType } from '@/features/table/modules/Types/types'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { useEditAmountTypeMutation, useGetAmountTypesQuery } from '@/features/table/modules/Types/typesApiSlice'
import { Dialog } from '@/components/Dialog/Dialog'
import { Fallback } from '@/components/Fallback/Fallback'
import { DragIcon } from '@/components/Icons/DragIcon'
import { TimesIcon } from '@/components/Icons/TimesIcon'
import { CirclePlusIcon } from '@/components/Icons/CirclePlusIcon'
import { changeModalIsOpen, getModalState } from '../modalSlice'
import { Titles } from '../titles'

export const AmountType = () => {
	const modal = useAppSelector(getModalState('amountType'))

	const { data, isFetching } = useGetAmountTypesQuery(null)

	const dispatch = useAppDispatch()

	const closeHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'amountType', isOpen: false }))
	}

	return (
		<Dialog
			title={Titles.AmountType}
			body={isFetching ? <Fallback /> : <CreateForm data={data?.data || []} />}
			open={modal.isOpen}
			onClose={closeHandler}
			maxWidth='md'
			fullWidth
		/>
	)
}

type Form = {
	data: IAmountType[]
	deleted: string[]
}

type FormProps = {
	data: IAmountType[]
}
const CreateForm: FC<FormProps> = ({ data }) => {
	const methods = useForm<Form>({ values: { data: data, deleted: [] } })
	const { fields, append, remove, move } = useFieldArray({ control: methods.control, name: 'data' })

	const [edit] = useEditAmountTypeMutation()

	const sensors = useSensors(useSensor(PointerSensor))

	const { palette } = useTheme()

	const dispatch = useAppDispatch()

	const closeHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'amountType', isOpen: false }))
	}

	const saveHandler = methods.handleSubmit(async data => {
		console.log('data', data.data)
		console.log('dirty', methods.formState.isDirty, methods.formState.dirtyFields)
		console.log(methods.formState)
		// methods.formState.dirtyFields.form?.forEach((f, i) => {
		// 	console.log('dirty field (index ' + i + ')', f)
		// })

		if (!methods.formState.isDirty) return

		try {
			await edit(data).unwrap()
			toast.success('Единицы измерения обновлены')
		} catch (error) {
			const fetchError = error as IFetchError
			toast.error(fetchError.data.message, { autoClose: false })
		}
	})

	const dragEndHandler = useCallback(
		(event: DragEndEvent) => {
			// console.log(event)
			const { active, over } = event
			if (active.id === over?.id) return

			const startIndex = active.data.current?.sortable.index
			const endIndex = over?.data.current?.sortable.index
			if (startIndex == undefined || endIndex == undefined) return

			move(startIndex, endIndex)
			// const values = methods.getValues()

			for (let i = Math.min(startIndex, endIndex); i <= Math.max(startIndex, endIndex); i++) {
				// update(i, { ...values.form[i], number: i + 1 })
				methods.setValue(`data.${i}.number`, i + 1)
				// methods.formState.dirtyFields.form[i].number = true
			}
		},
		[move, methods]
	)

	const addNewHandler = () => {
		append({ id: '', name: '', description: '', number: fields.length + 1 })
	}

	const removeHandler = (index: number) => {
		const values = methods.getValues()
		methods.setValue(`deleted.${values.deleted.length}`, values.data[index].id)
		remove(index)
	}

	return (
		<Stack component={'form'} onSubmit={saveHandler}>
			<Stack spacing={1.5} ml={-2} alignItems={'flex-start'}>
				<DndContext sensors={sensors} collisionDetection={closestCenter} onDragEnd={dragEndHandler}>
					<SortableContext items={fields} strategy={verticalListSortingStrategy}>
						{fields.map((d, i) => (
							<Item
								key={d.id}
								id={d.id}
								index={i}
								errors={methods.formState.errors}
								register={methods.register}
								remove={removeHandler}
							/>
						))}
					</SortableContext>
				</DndContext>

				<Button
					onClick={addNewHandler}
					variant='outlined'
					sx={{ ml: '46px!important', width: 250, textTransform: 'inherit' }}
				>
					<CirclePlusIcon fill={palette.primary.main} fontSize={16} mr={1} />
					Добавить значение
				</Button>
			</Stack>

			<Divider sx={{ mt: 2, mb: 2 }} />
			<Stack direction={'row'} spacing={3}>
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

type ItemProps = {
	id: string
	index: number
	register: UseFormRegister<Form>
	remove: (index: number) => void
	errors?: FieldErrors<Form>
}
const Item: FC<ItemProps> = ({ id, index, register, remove, errors }) => {
	const { attributes, listeners, setNodeRef, transform, transition } = useSortable({ id })

	const removeHandler = () => {
		remove(index)
	}

	return (
		<Stack
			ref={setNodeRef}
			direction={'row'}
			spacing={1}
			alignItems={'center'}
			width={'100%'}
			sx={{
				transform: CSS.Transform.toString(transform),
				transition: transition,
			}}
		>
			<IconButton {...attributes} {...listeners} sx={{ cursor: 'grab' }}>
				<DragIcon fill={'#a8a8a8'} fontSize={24} />
			</IconButton>

			<TextField
				label='Обозначение'
				{...register(`data.${index}.name`, { required: true })}
				error={Boolean(errors?.data && errors.data[index]?.name)}
				fullWidth
				sx={{ maxWidth: 250 }}
			/>
			<TextField
				label='Название'
				{...register(`data.${index}.description`, { required: true })}
				error={Boolean(errors?.data && errors.data[index]?.description)}
				fullWidth
			/>

			{index != 0 && (
				<IconButton onClick={removeHandler}>
					<TimesIcon fontSize={18} padding={0.4} />
				</IconButton>
			)}
		</Stack>
	)
}
