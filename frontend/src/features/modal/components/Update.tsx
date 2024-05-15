import { useEffect, useState } from 'react'
import { Button, Stack, useTheme } from '@mui/material'
import { toast } from 'react-toastify'

import type { IFetchError } from '@/app/types/error'
import type { IUpdateDataItem } from '@/features/table/types/data'
import type { IBaseInfForm } from '@/components/forms/BaseInfForm/type'
import type { IControlForm } from '@/components/forms/ControlForm/type'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { useGetByIdQuery, useUpdateMutation } from '@/features/table/tableApiSlice'
import { getContextMenu, setContextMenu } from '@/features/table/tableSlice'
import { useGetAmountTypesQuery, useGetReagentTypesQuery } from '@/features/table/modules/Types/typesApiSlice'
import { Dialog } from '@/components/Dialog/Dialog'
import { Stepper } from '@/components/Stepper/Stepper'
import { BaseInfForm } from '@/components/forms/BaseInfForm/BaseInfForm'
import { ControlForm } from '@/components/forms/ControlForm/ControlForm'
import { Confirm } from '@/components/Confirm/Confirm'
import { FileDeleteIcon } from '@/components/Icons/FileDeleteIcon'
import { Fallback } from '@/components/Fallback/Fallback'
import { changeModalIsOpen, getModalState } from '../modalSlice'
import { Titles } from '../titles'
import { AmountType } from './AmountType'

export const Update = () => {
	const modal = useAppSelector(getModalState('edit'))

	const dispatch = useAppDispatch()

	const closeHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'edit', isOpen: false }))
		dispatch(setContextMenu())
	}

	return (
		<Dialog
			title={Titles.Edit}
			body={<UpdateForm />}
			open={modal.isOpen}
			onClose={closeHandler}
			maxWidth='md'
			fullWidth
		/>
	)
}

const steps = [
	{ id: 'base', label: 'Основные сведения' },
	{ id: 'control', label: 'Проведение входного контроля' },
]

const UpdateForm = () => {
	const [activeStep, setActiveStep] = useState(0)
	const [base, setBase] = useState<IBaseInfForm>()
	const [control, setControl] = useState<IControlForm>()
	const contextMenu = useAppSelector(getContextMenu)

	const dispatch = useAppDispatch()

	const { palette } = useTheme()

	const { isLoading: isLoadingRTypes } = useGetReagentTypesQuery(null)
	const { isLoading: isLoadingATypes } = useGetAmountTypesQuery(null)

	const { data, isLoading } = useGetByIdQuery(contextMenu?.active || '', { skip: !contextMenu?.active })

	const [update] = useUpdateMutation()

	useEffect(() => {
		if (!data) return

		setBase({
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
		})
		setControl({
			receiptDate: data.data.receiptDate,
			amount: data.data.amount.toString(),
			amountType: data.data.amountTypeId,
			controlDate: data.data.controlDate,
			protocol: data.data.protocol,
			result: data.data.result,
		})
	}, [data])

	const closeHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'edit', isOpen: false }))
		dispatch(setContextMenu())
	}

	const cancelHandler = () => {
		if (activeStep == 0) {
			closeHandler()
		} else {
			// localStorage.setItem(localKeys.control, JSON.stringify(data))
			setActiveStep(prev => prev - 1)
		}
	}

	const saveBaseHandler = (data: IBaseInfForm, isShouldUpdate?: boolean) => {
		console.log('base', data, isShouldUpdate)
		if (isShouldUpdate) {
			setBase(data)
		}
		setActiveStep(1)
	}
	const saveControlHandler = async (form: IControlForm, isShouldUpdate?: boolean) => {
		console.log('control', form, isShouldUpdate)
		if (!data || !base) return

		const newData: IUpdateDataItem = {
			id: data.data.id,
			typeId: base.type,
			name: base.name,
			uname: base.uname,
			document: base.document,
			purity: base.purity,
			dateOfManufacture: base.dateOfManufacture,
			consignment: base.consignment,
			manufacturer: base.manufacturer,
			shelfLife: +base.shelfLife,
			place_closet: base.place_closet,
			place_shelf: +base.place_shelf,
			receiptDate: form.receiptDate,
			amount: +form.amount,
			amountTypeId: form.amountType,
			controlDate: form.controlDate,
			protocol: form.protocol,
			result: form.result,
		}

		try {
			await update(newData).unwrap()

			toast.success('Реактив обновлен')
			setActiveStep(0)
			closeHandler()
		} catch (error) {
			const fetchError = error as IFetchError
			toast.error(fetchError.data.message, { autoClose: false })
		}
	}

	const deleteHandler = async () => {}

	return (
		<Stack mt={-2.5} overflow={'hidden'}>
			<Stack direction={'row'} width={'100%'} alignItems={'center'} mb={2.5}>
				<Stepper steps={steps} active={activeStep} sx={{ width: '100%' }} />

				<Confirm
					width='64px'
					onClick={deleteHandler}
					buttonComponent={
						<Button variant='outlined' color='error'>
							<FileDeleteIcon fontSize={20} fill={palette.error.main} />
						</Button>
					}
					confirmText='Вы уверены, что хотите удалить реактив?'
				/>
			</Stack>
			<AmountType />

			{isLoadingRTypes || isLoadingATypes || isLoading || !base || !control ? (
				<Fallback marginTop={5} marginBottom={3} />
			) : (
				<Stack
					flexDirection={'row'}
					width={'200%'}
					height={activeStep == 0 ? 669 : 333}
					sx={{ transform: `translateX(-${activeStep * 50}%)`, transition: 'all .3s ease-in-out' }}
				>
					<Stack width={'100%'}>
						<BaseInfForm
							defaultValues={base}
							submitLabel='Далее'
							onSubmit={saveBaseHandler}
							onCancel={cancelHandler}
						/>
					</Stack>
					<Stack width={'100%'}>
						<ControlForm
							defaultValues={control}
							submitLabel='Сохранить'
							onSubmit={saveControlHandler}
							cancelLabel='Назад'
							onCancel={cancelHandler}
						/>
					</Stack>
				</Stack>
			)}
		</Stack>
	)
}
