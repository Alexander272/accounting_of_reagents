import { FC } from 'react'
import { Button, Stack } from '@mui/material'
import { FormProvider, useForm } from 'react-hook-form'

import type { IBaseInfForm } from './type'
import { Inputs } from './Inputs'

type Props = {
	defaultValues: IBaseInfForm
	submitLabel?: string
	cancelLabel?: string
	disabled?: boolean
	onCancel?: () => void
	onSubmit: (data: IBaseInfForm, isShouldUpdate?: boolean) => void
}

export const BaseInfForm: FC<Props> = ({ defaultValues, submitLabel, cancelLabel, disabled, onSubmit, onCancel }) => {
	const methods = useForm<IBaseInfForm>({
		values: defaultValues,
	})

	const submitHandler = methods.handleSubmit(data => {
		onSubmit(data, Boolean(Object.keys(methods.formState.dirtyFields).length))
	})

	return (
		<Stack component={'form'} onSubmit={submitHandler}>
			<FormProvider {...methods}>
				<Inputs disabled={disabled} />
			</FormProvider>

			<Stack direction={'row'} spacing={3} mt={4}>
				{onCancel && (
					<Button onClick={onCancel} variant='outlined' fullWidth disabled={disabled}>
						{cancelLabel ? cancelLabel : 'Отменить'}
					</Button>
				)}
				<Button variant='contained' type='submit' fullWidth disabled={disabled}>
					{submitLabel ? submitLabel : 'Сохранить'}
				</Button>
			</Stack>
		</Stack>
	)
}
