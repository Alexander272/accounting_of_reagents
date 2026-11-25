import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { Button, ButtonGroup, Stack, Tooltip, Typography } from '@mui/material'
import { DatePicker } from '@mui/x-date-pickers'
import dayjs, { type Dayjs } from 'dayjs'
import quarterOfYear from 'dayjs/plugin/quarterOfYear'

import type { IReportPeriod } from '../types'
import { DateFormat } from '@/constants/date'
import { AppRoutes } from '@/constants/routes'
import { DownloadIcon } from '@/components/Icons/DownloadIcon'
import { LeftArrowIcon } from '@/components/Icons/LeftArrowIcon'

dayjs.extend(quarterOfYear)

type Props = {
	period: IReportPeriod
	canGenerate: boolean
	canDownload: boolean
	onGenerate: (period: IReportPeriod) => void
	onDownload: () => void
	total: number
}

type PeriodType = 'month' | 'prev' | 'quarter'

interface DateRange {
	from: Dayjs
	to: Dayjs
}

const PERIOD_PRESETS: { type: PeriodType; label: string }[] = [
	{ type: 'month', label: 'Этот месяц' },
	{ type: 'prev', label: 'Прошлый месяц' },
	{ type: 'quarter', label: 'Квартал' },
]

const PERIOD_GENERATORS: Record<PeriodType, () => DateRange> = {
	// Текущий месяц (с 1-го по последнее число)
	month: () => ({
		from: dayjs().startOf('month'),
		to: dayjs().endOf('month'),
	}),
	// Прошлый месяц (целиком)
	prev: () => {
		const prevMonth = dayjs().subtract(1, 'month')
		return {
			from: prevMonth.startOf('month'),
			to: prevMonth.endOf('month'),
		}
	},
	// Текущий квартал
	quarter: () => ({
		from: dayjs().startOf('quarter'),
		to: dayjs().endOf('quarter'),
	}),
}

export const ReportHeader = ({ period, canGenerate, canDownload, onGenerate, onDownload, total }: Props) => {
	const navigate = useNavigate()

	const [activeType, setActiveType] = useState<PeriodType | 'custom'>('month')
	const [from, setFrom] = useState<Dayjs | null>(() => PERIOD_GENERATORS.month().from)
	const [to, setTo] = useState<Dayjs | null>(() => PERIOD_GENERATORS.month().to)

	const isValidPeriod = Boolean(from?.isValid() && to?.isValid() && !from.isAfter(to, 'day'))

	const backHandler = () => navigate(AppRoutes.Home)

	const generateHandler = () => {
		if (!from || !to || !isValidPeriod) return
		onGenerate({ from, to })
	}

	const handleDateChange = (field: 'from' | 'to') => (value: Dayjs | null) => {
		setActiveType('custom')
		if (field === 'from') setFrom(value)
		else setTo(value)
	}

	const selectPeriodHandler = (type: PeriodType) => {
		const range = PERIOD_GENERATORS[type]()
		setActiveType(type)
		setFrom(range.from)
		setTo(range.to)
	}

	return (
		<Stack>
			<Stack direction={'row'} spacing={2}>
				<Tooltip title='Назад'>
					<Button onClick={backHandler} sx={{ minWidth: 40 }}>
						<LeftArrowIcon fontSize={16} />
					</Button>
				</Tooltip>

				<Stack sx={{ ml: 0.5 }}>
					<Typography variant='h6' sx={{ fontWeight: 'bold' }}>
						Расход реактивов
					</Typography>
					<Typography sx={{ color: '#6b778c' }}>
						за период с {period.from.format(DateFormat)} по {period.to.format(DateFormat)}
					</Typography>
				</Stack>
			</Stack>

			<Stack
				direction={'row'}
				alignItems={'center'}
				justifyContent={'space-between'}
				flexWrap={'wrap'}
				gap={2}
				sx={{
					borderRadius: 3,
					paddingX: 2,
					paddingY: 2,
					width: '100%',
					border: '1px solid rgba(0, 0, 0, 0.12)',
					height: 'fit-content',
					backgroundColor: '#fff',
					my: 1,
				}}
			>
				<Stack direction={'row'} alignItems={'center'} flexWrap={'wrap'} gap={2}>
					<DatePicker
						name='from'
						label='Период с'
						value={from}
						maxDate={to ?? undefined}
						onChange={handleDateChange('from')}
						format={DateFormat}
					/>
					<DatePicker
						name='to'
						label='Период по'
						value={to}
						minDate={from ?? undefined}
						onChange={handleDateChange('to')}
						format={DateFormat}
					/>

					<ButtonGroup sx={{ height: 40 }}>
						{PERIOD_PRESETS.map(({ type, label }) => (
							<Button
								key={type}
								onClick={() => selectPeriodHandler(type)}
								color={activeType === type ? 'secondary' : 'gray'}
								sx={{ textTransform: 'inherit' }}
							>
								{label}
							</Button>
						))}
					</ButtonGroup>

					<Button
						onClick={generateHandler}
						variant='outlined'
						disabled={!canGenerate || !isValidPeriod}
						sx={{ textTransform: 'inherit', height: 40 }}
					>
						Сформировать
					</Button>
				</Stack>

				<Button onClick={onDownload} disabled={!canDownload} sx={{ textTransform: 'inherit' }}>
					<DownloadIcon
						sx={theme => ({
							fontSize: 14,
							mr: 1,
							fill: !canDownload ? theme.palette.action.disabled : theme.palette.primary.main,
						})}
					/>
					Скачать CSV
				</Button>
			</Stack>

			<Typography sx={{ mb: 1, ml: 0.5, color: '#6b778c' }}>
				Позиций:{' '}
				<Typography component={'b'} color='text.primary' sx={{ fontWeight: 'bold' }}>
					{total || '-'}
				</Typography>
			</Typography>
		</Stack>
	)
}
