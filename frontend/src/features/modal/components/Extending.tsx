import { useState } from 'react'
import { Box, Button, Stack, TextField, Typography, useTheme } from '@mui/material'
import { Controller, useForm } from 'react-hook-form'
import { DatePicker } from '@mui/x-date-pickers'
import { toast } from 'react-toastify'
import dayjs from 'dayjs'

import type { IFetchError } from '@/app/types/error'
import type { CreateExtending, IExtendingForm } from '@/features/table/modules/Extending/extending'
import { DateFormat } from '@/constants/date'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { getContextMenu, setContextMenu } from '@/features/table/tableSlice'
import {
	useCreateExtendingMutation,
	useDeleteExtendingMutation,
	useGetExtendingQuery,
	useUpdateExtendingMutation,
} from '@/features/table/modules/Extending/extendingApiSlice'
import { useGetByIdQuery } from '@/features/table/tableApiSlice'
import { NoRowsOverlay } from '@/features/table/components/NoRowsOverlay/components/NoRowsOverlay'
import { Dialog } from '@/components/Dialog/Dialog'
import { Fallback } from '@/components/Fallback/Fallback'
import { Table } from '@/components/Table/Table'
import { TableHead } from '@/components/Table/TableHead'
import { TableBody } from '@/components/Table/TableBody'
import { TableRow } from '@/components/Table/TableRow'
import { TableCell } from '@/components/Table/TableCell'
import { Confirm } from '@/components/Confirm/Confirm'
import { FileDeleteIcon } from '@/components/Icons/FileDeleteIcon'
import { PlusIcon } from '@/components/Icons/PlusIcon'
import { SaveIcon } from '@/components/Icons/SaveIcon'
import { changeModalIsOpen, getModalState } from '../modalSlice'
import { Titles } from '../titles'

export const Extending = () => {
	const modal = useAppSelector(getModalState('extending'))

	const dispatch = useAppDispatch()

	const closeHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'extending', isOpen: false }))
		dispatch(setContextMenu())
	}

	return (
		<Dialog
			title={Titles.Extending}
			body={<ExtendingForm />}
			open={modal.isOpen}
			onClose={closeHandler}
			maxWidth='sm'
			fullWidth
		/>
	)
}

const defaultValues = {
	date: dayjs().unix(),
	period: '',
}

export const ExtendingForm = () => {
	const [selected, setSelected] = useState<number>()
	const contextMenu = useAppSelector(getContextMenu)

	const { palette } = useTheme()

	const { data: reagent, isLoading: isLoadingReagent } = useGetByIdQuery(contextMenu?.active || '', {
		skip: !contextMenu?.active,
	})
	const { data, isLoading } = useGetExtendingQuery(contextMenu?.active || '', { skip: !contextMenu?.active })

	const [create] = useCreateExtendingMutation()
	const [update] = useUpdateExtendingMutation()
	const [remove] = useDeleteExtendingMutation()

	const methods = useForm<IExtendingForm>({ values: defaultValues })

	const submitHandler = async (form: IExtendingForm) => {
		if (!reagent || !data) return

		const extending: CreateExtending = {
			reagentId: reagent.data.id,
			date: form.date,
			period: +form.period,
		}

		try {
			if (selected != undefined) await update({ ...extending, id: data.data[selected].id }).unwrap()
			else await create(extending).unwrap()

			setSelected(undefined)
			methods.reset(defaultValues)
		} catch (error) {
			const fetchError = error as IFetchError
			toast.error(fetchError.data.message, { autoClose: false })
		}
	}

	const selectHandler = (index: number) => () => {
		if (!data) return
		methods.reset({ date: data.data[index].date, period: data.data[index].period.toString() })

		if (index == selected) setSelected(undefined)
		else setSelected(index)
	}

	const deleteHandler = (index: number) => async () => {
		if (!data) return
		try {
			await remove(data.data[index]).unwrap()
			setSelected(undefined)
			methods.reset(defaultValues)
		} catch (error) {
			const fetchError = error as IFetchError
			toast.error(fetchError.data.message, { autoClose: false })
		}
	}

	return (
		<Stack mt={-3}>
			{isLoading || isLoadingReagent ? <Fallback marginTop={5} marginBottom={3} /> : null}

			<Stack mb={3}>
				<Typography fontSize={'1.4rem'} textAlign={'center'}>
					{reagent?.data.name}
				</Typography>
				{/* <Typography>
					{reagent?.data.manufacturer} {reagent?.data.consignment}
				</Typography> */}
			</Stack>

			<Stack spacing={1} direction={'row'} component={'form'} onSubmit={methods.handleSubmit(submitHandler)}>
				<Controller
					control={methods.control}
					name='date'
					rules={{ required: true }}
					render={({ field, fieldState: { error } }) => (
						<DatePicker
							{...field}
							value={dayjs(+field.value * 1000)}
							onChange={value => field.onChange(value?.unix())}
							label={'Дата расхода'}
							showDaysOutsideCurrentMonth
							fixedWeekNumber={6}
							slotProps={{
								textField: {
									fullWidth: true,
									error: Boolean(error),
								},
							}}
						/>
					)}
				/>
				<TextField
					label={'Период продления, мес'}
					type='number'
					fullWidth
					error={Boolean(methods.formState.errors['period'])}
					{...methods.register('period', { required: true })}
				/>

				<Button type='submit' variant='outlined' sx={{ minWidth: 40 }}>
					{selected == undefined ? (
						<PlusIcon fontSize={16} fill={palette.primary.main} />
					) : (
						<SaveIcon fontSize={16} fill={palette.primary.main} />
					)}
				</Button>
			</Stack>

			<Table>
				<TableHead>
					<TableRow height={40} sx={{ padding: 0 }}>
						<TableCell width={252}>
							<Typography>Дата</Typography>
						</TableCell>
						<TableCell width={252}>
							<Typography>Период продления, мес</Typography>
						</TableCell>
						<TableCell width={48}></TableCell>
					</TableRow>
				</TableHead>
				<Box maxHeight={350} overflow={'auto'}>
					{!data?.data.length && (
						<Box position={'relative'} height={100}>
							<NoRowsOverlay />
						</Box>
					)}

					<TableBody>
						{data?.data.map((d, i) => (
							<TableRow
								key={d.id}
								onClick={selectHandler(i)}
								hover
								height={40}
								sx={{ background: selected == i ? palette.rowActive.main : undefined, padding: 0 }}
							>
								<TableCell width={252}>{dayjs(d.date * 1000).format(DateFormat)}</TableCell>
								<TableCell width={252}>{d.period}</TableCell>
								<TableCell width={48}>
									<Confirm
										buttonComponent={
											<Button color='error' sx={{ minWidth: 40 }}>
												<FileDeleteIcon fontSize={18} />
											</Button>
										}
										onClick={deleteHandler(i)}
										confirmText='Вы уверены, что хотите удалить продление?'
									/>
								</TableCell>
							</TableRow>
						))}
					</TableBody>
				</Box>
			</Table>
		</Stack>
	)
}
