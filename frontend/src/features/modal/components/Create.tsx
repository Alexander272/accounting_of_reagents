import { FC, useRef, useState } from 'react'
import { Button, Stack, Tooltip } from '@mui/material'
import { toast } from 'react-toastify'
import dayjs from 'dayjs'

import type { IFetchError } from '@/app/types/error'
import type { ICreateDataItem } from '@/features/table/types/data'
import type { IBaseInfForm } from '@/components/forms/BaseInfForm/type'
import type { IControlForm } from '@/components/forms/ControlForm/type'
import { useAppDispatch, useAppSelector } from '@/hooks/redux'
import { localKeys } from '@/features/table/constants/localKeys'
import { useGetAmountTypesQuery, useGetReagentTypesQuery } from '@/features/table/modules/Types/typesApiSlice'
import { useCreateMutation } from '@/features/table/tableApiSlice'
import { BaseInfForm } from '@/components/forms/BaseInfForm/BaseInfForm'
import { Dialog } from '@/components/Dialog/Dialog'
import { Stepper } from '@/components/Stepper/Stepper'
import { Fallback } from '@/components/Fallback/Fallback'
import { RefreshIcon } from '@/components/Icons/RefreshIcon'
import { ControlForm } from '@/components/forms/ControlForm/ControlForm'
import { changeModalIsOpen, getModalState } from '../modalSlice'
import { Titles } from '../titles'

export const Create = () => {
	const modal = useAppSelector(getModalState('create'))

	const dispatch = useAppDispatch()

	const closeHandler = () => {
		dispatch(changeModalIsOpen({ variant: 'create', isOpen: false }))
	}

	return (
		<Dialog
			title={Titles.Create}
			body={<CreateForm closing={!modal.isOpen} />}
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

const BaseData: IBaseInfForm = {
	type: '',
	name: '',
	uname: '',
	document: '',
	purity: '',
	dateOfManufacture: dayjs().unix(),
	consignment: '',
	manufacturer: '',
	shelfLife: '',
	place_closet: '',
	place_shelf: '',
}

const ControlData: IControlForm = {
	receiptDate: dayjs().unix(),
	amount: '',
	amountType: '',
	controlDate: dayjs().unix(),
	protocol: '',
	result: true,
}

type FormProps = {
	closing?: boolean
}

const CreateForm: FC<FormProps> = ({ closing }) => {
	const { data: rTypes, isLoading: isLoadingRTypes } = useGetReagentTypesQuery(null)
	BaseData.type = rTypes?.data[0]?.id || ''

	const { data: aTypes, isLoading: isLoadingATypes } = useGetAmountTypesQuery(null)
	ControlData.amountType = aTypes?.data[0]?.id || ''

	const [activeStep, setActiveStep] = useState(0)
	const isForceSave = useRef(false)

	const [base, setBase] = useState<IBaseInfForm>(
		JSON.parse(localStorage.getItem(localKeys.baseInfo) || 'null') || BaseData
	)
	const [control, setControl] = useState<IControlForm>(
		JSON.parse(localStorage.getItem(localKeys.control) || 'null') || ControlData
	)

	const dispatch = useAppDispatch()

	const [create] = useCreateMutation()

	const cancelHandler = () => {
		if (activeStep == 0) {
			dispatch(changeModalIsOpen({ variant: 'create', isOpen: false }))
		} else {
			isForceSave.current = true
			// localStorage.setItem(localKeys.control, JSON.stringify(data))
			setActiveStep(prev => prev - 1)
		}
	}

	const saveBaseHandler = (data: IBaseInfForm, isShouldUpdate?: boolean) => {
		console.log('base', data, isShouldUpdate)
		if (isShouldUpdate) {
			localStorage.setItem(localKeys.baseInfo, JSON.stringify(data))
			setBase(data)
		}
		if (!closing) setActiveStep(1)
	}
	const saveControlHandler = async (data: IControlForm, isShouldUpdate?: boolean) => {
		console.log('control', data, isShouldUpdate)
		if (isShouldUpdate) {
			setControl(data)
			if (closing || isForceSave.current) localStorage.setItem(localKeys.control, JSON.stringify(data))
		}
		isForceSave.current = false
		if (!closing) {
			const data: ICreateDataItem = {
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
				receiptDate: control.receiptDate,
				amount: +control.amount,
				amountTypeId: control.amountType,
				controlDate: control.controlDate,
				protocol: control.protocol,
				result: control.result,
			}

			try {
				await create(data).unwrap()

				localStorage.removeItem(localKeys.baseInfo)
				localStorage.removeItem(localKeys.control)
				setBase(BaseData)
				setControl(ControlData)
				toast.success('Реактив создан')
				setActiveStep(0)
			} catch (error) {
				const fetchError = error as IFetchError
				toast.error(fetchError.data.message, { autoClose: false })
			}
		}
	}

	const deleteHandler = () => {
		console.log('delete')
		if (activeStep == 0) {
			localStorage.removeItem(localKeys.baseInfo)
			setBase(BaseData)
		}
		if (activeStep == 1) {
			localStorage.removeItem(localKeys.control)
			setControl(ControlData)
		}
	}

	return (
		<Stack mt={-2.5} overflow={'hidden'}>
			<Stack direction={'row'} width={'100%'} alignItems={'center'} mb={2.5}>
				<Stepper steps={steps} active={activeStep} sx={{ width: '100%' }} />

				<Tooltip title='Очистить'>
					<span>
						<Button variant='outlined' color='gray' onClick={deleteHandler}>
							<RefreshIcon fontSize={18} />
						</Button>
					</span>
				</Tooltip>
			</Stack>

			{isLoadingRTypes || isLoadingATypes ? (
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
							forceSave={closing}
							submitLabel='Далее'
							onSubmit={saveBaseHandler}
							onCancel={cancelHandler}
						/>
					</Stack>
					<Stack width={'100%'}>
						<ControlForm
							defaultValues={control}
							forceSave={isForceSave.current || closing}
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
