import { FC } from 'react'
import { Stepper as BaseStepper, Step, StepButton, StepLabel, SxProps, Theme } from '@mui/material'
import { CustomConnector } from './stepper.style'
import { CustomStepIcon } from './Icon'

type Step = {
	id: string
	completed?: boolean
	label: string
}

type Props = {
	active?: number
	steps: Step[]
	onClick?: (index: number) => void
	sx?: SxProps<Theme>
}

export const Stepper: FC<Props> = ({ active, steps, onClick, sx }) => {
	const selectHandler = (index: number) => () => {
		onClick && onClick(index)
	}

	return (
		<BaseStepper
			alternativeLabel
			nonLinear={Boolean(onClick)}
			activeStep={active}
			connector={<CustomConnector />}
			sx={sx}
		>
			{steps.map((step, i) => (
				<Step key={step.id} completed={step.completed}>
					<StepButton onClick={selectHandler(i)} sx={{ borderRadius: 4, margin: 0, padding: 0 }}>
						<StepLabel StepIconComponent={CustomStepIcon}>{step.label}</StepLabel>
					</StepButton>
				</Step>
			))}
		</BaseStepper>
	)
}
