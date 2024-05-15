import { Button, List, ListItem, ListSubheader, Stack, TextField, Typography, useTheme } from '@mui/material'
import { Controller, useForm } from 'react-hook-form'
import { DatePicker } from '@mui/x-date-pickers'
import { toast } from 'react-toastify'
import dayjs from 'dayjs'

import type { IFetchError } from '@/app/types/error'
import type { CreateSpending, ISpendingForm } from '@/features/table/types/spending'
import { DateFormat } from '@/constants/date'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { useCreateSpendingMutation, useGetSpendingQuery } from '@/features/table/spendingApiSlice'
import { useGetByIdQuery } from '@/features/table/tableApiSlice'
import { getContextMenu, setContextMenu } from '@/features/table/tableSlice'
import { Dialog } from '@/components/Dialog/Dialog'
import { Fallback } from '@/components/Fallback/Fallback'
import { PlusIcon } from '@/components/Icons/PlusIcon'
import { changeModalIsOpen, getModalState } from '../modalSlice'
import { Titles } from '../titles'
import { useGetAmountTypesQuery } from '@/features/table/modules/Types/typesApiSlice'

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

const SpendingForm = () => {
	const contextMenu = useAppSelector(getContextMenu)

	const { palette } = useTheme()

	const { data: aTypes } = useGetAmountTypesQuery(null)
	const { data: reagent, isLoading: isLoadingReagent } = useGetByIdQuery(contextMenu?.active || '', {
		skip: !contextMenu?.active,
	})
	const { data, isLoading } = useGetSpendingQuery(contextMenu?.active || '', { skip: !contextMenu?.active })

	const [create] = useCreateSpendingMutation()

	const methods = useForm<ISpendingForm>({
		values: { date: dayjs().startOf('d').unix(), amount: '' },
	})

	const submitHandler = async (data: ISpendingForm) => {
		console.log(data)
		if (!reagent) return

		const spending: CreateSpending = {
			reagentId: reagent.data.id,
			date: data.date,
			amount: +data.amount,
		}

		try {
			await create(spending).unwrap()
		} catch (error) {
			const fetchError = error as IFetchError
			toast.error(fetchError.data.message, { autoClose: false })
		}
	}

	const type = aTypes?.data.find(a => a.id == reagent?.data.amountTypeId)

	return (
		<Stack mt={-3} mb={-1.5} overflow={'hidden'}>
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
					rules={{ required: true }}
					render={({ field, fieldState: { error } }) => (
						<DatePicker
							{...field}
							value={dayjs(+field.value * 1000)}
							onChange={value => field.onChange(value?.startOf('d').unix())}
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
					label={'Количество ' + (type?.name ? `(${type?.name})` : '')}
					type='number'
					fullWidth
					error={Boolean(methods.formState.errors['amount'])}
					{...methods.register('amount', { required: true })}
				/>

				<Button type='submit' variant='outlined' sx={{ minWidth: 40 }}>
					<PlusIcon fontSize={16} fill={palette.primary.main} />
				</Button>
			</Stack>

			<List
				subheader={<ListSubheader>Расходование</ListSubheader>}
				sx={{ width: '100%', maxHeight: 350, overflow: 'auto' }}
			>
				{data?.data.map(d => (
					<ListItem key={d.id}>
						{dayjs(d.date * 1000).format(DateFormat)} - {d.amount} {type?.name}
					</ListItem>
				))}
				{!data?.data.length && (
					<ListItem sx={{ justifyContent: 'center', color: '#7a7a7a' }}>Расход не найден</ListItem>
				)}
			</List>
		</Stack>
	)
}
