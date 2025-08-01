import { useState } from 'react'
import { Box, Button, Stack, TextField, Typography, useTheme } from '@mui/material'
import { Controller, useForm } from 'react-hook-form'
import { DatePicker } from '@mui/x-date-pickers'
import { toast } from 'react-toastify'
import dayjs from 'dayjs'

import type { IFetchError } from '@/app/types/error'
import type { CreateSpending, ISpendingForm } from '@/features/table/modules/Spending/spending'
import { DateFormat } from '@/constants/date'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import {
	useCreateSpendingMutation,
	useDeleteSpendingMutation,
	useGetSpendingQuery,
	useUpdateSpendingMutation,
} from '@/features/table/modules/Spending/spendingApiSlice'
import { useGetAmountTypesQuery } from '@/features/table/modules/Types/typesApiSlice'
import { useGetByIdQuery } from '@/features/table/tableApiSlice'
import { getContextMenu, setContextMenu } from '@/features/table/tableSlice'
import { NoRowsOverlay } from '@/features/table/components/NoRowsOverlay/components/NoRowsOverlay'
import { Dialog } from '@/components/Dialog/Dialog'
import { Fallback } from '@/components/Fallback/Fallback'
import { Table } from '@/components/Table/Table'
import { TableHead } from '@/components/Table/TableHead'
import { TableRow } from '@/components/Table/TableRow'
import { TableCell } from '@/components/Table/TableCell'
import { TableBody } from '@/components/Table/TableBody'
import { Confirm } from '@/components/Confirm/Confirm'
import { PlusIcon } from '@/components/Icons/PlusIcon'
import { SaveIcon } from '@/components/Icons/SaveIcon'
import { FileDeleteIcon } from '@/components/Icons/FileDeleteIcon'
import { changeModalIsOpen, getModalState } from '../modalSlice'
import { Titles } from '../titles'

export const Spending = () => {
	const modal = useAppSelector(getModalState('spending'))

	const dispatch = useAppDispatch()

	const closeHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'spending', isOpen: false }))
		dispatch(setContextMenu())
	}

	return (
		<Dialog
			title={Titles.Spending}
			body={<SpendingForm />}
			open={modal.isOpen}
			onClose={closeHandler}
			maxWidth='sm'
			fullWidth
		/>
	)
}

const defaultValues = {
	date: 0,
	amount: '',
}

const SpendingForm = () => {
	const [selected, setSelected] = useState<number>()
	const contextMenu = useAppSelector(getContextMenu)

	const { palette } = useTheme()

	const { data: aTypes } = useGetAmountTypesQuery(null)
	const { data: reagent, isLoading: isLoadingReagent } = useGetByIdQuery(contextMenu?.active || '', {
		skip: !contextMenu?.active,
	})
	const { data, isLoading } = useGetSpendingQuery(contextMenu?.active || '', { skip: !contextMenu?.active })

	const [create] = useCreateSpendingMutation()
	const [update] = useUpdateSpendingMutation()
	const [remove] = useDeleteSpendingMutation()

	const methods = useForm<ISpendingForm>({ values: defaultValues })

	const submitHandler = async (form: ISpendingForm) => {
		if (!reagent || !data) return

		const spending: CreateSpending = {
			reagentId: reagent.data.id,
			date: form.date,
			amount: +form.amount,
		}

		try {
			if (selected != undefined) await update({ ...spending, id: data.data[selected].id }).unwrap()
			else await create(spending).unwrap()

			setSelected(undefined)
			methods.reset(defaultValues)
		} catch (error) {
			const fetchError = error as IFetchError
			toast.error(fetchError.data.message, { autoClose: false })
		}
	}

	const selectHandler = (index: number) => () => {
		if (!data) return

		if (index == selected) {
			setSelected(undefined)
			methods.reset(defaultValues)
		} else {
			setSelected(index)
			methods.reset({ date: data.data[index].date, amount: data.data[index].amount.toString() })
		}
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

	const type = aTypes?.data.find(a => a.id == reagent?.data.amountTypeId)

	return (
		<Stack mt={-3} mb={-1.5}>
			{isLoading || isLoadingReagent ? <Fallback marginTop={5} marginBottom={3} /> : null}

			<Stack mb={4}>
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
					rules={{ required: true, min: 100000 }}
					render={({ field, fieldState: { error } }) => (
						<DatePicker
							{...field}
							value={field.value ? dayjs(+field.value * 1000) : null}
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
				<Controller
					control={methods.control}
					name='amount'
					rules={{ required: true, min: 0.00001 }}
					render={({ field, fieldState: { error } }) => (
						<TextField
							label={'Количество ' + (type?.name ? `(${type?.name})` : '')}
							type='number'
							fullWidth
							{...field}
							error={Boolean(error)}
							inputProps={{
								step: 0.00001,
								min: 0,
							}}
						/>
					)}
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
							<Typography>Количество</Typography>
						</TableCell>
						<TableCell width={48}></TableCell>
					</TableRow>
				</TableHead>
				<Box maxHeight={350} overflow={'auto'} position={'relative'} minHeight={100}>
					{!data?.data.length && <NoRowsOverlay />}

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
								<TableCell width={252}>
									{d.amount} {type?.name}
								</TableCell>
								<TableCell width={48}>
									<Confirm
										buttonComponent={
											<Button color='error' sx={{ minWidth: 40 }}>
												<FileDeleteIcon fontSize={18} />
											</Button>
										}
										onClick={deleteHandler(i)}
										confirmText='Вы уверены, что хотите удалить расход?'
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
