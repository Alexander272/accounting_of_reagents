import { Box } from '@mui/material'

import { ReportHeader } from '@/features/report/components/ReportHeader'
import { ReportTable } from '@/features/report/components/ReportTable'
import { useSpendingReport } from '@/features/report/hooks/useSpendingReport'
import { PageBox } from '@/components/PageBox/PageBox'

export default function Report() {
	const { period, items, isFetching, isUninitialized, canGenerate, canDownload, generate, download } =
		useSpendingReport()

	return (
		<PageBox>
			<ReportHeader
				period={period}
				canGenerate={canGenerate}
				canDownload={canDownload}
				onGenerate={generate}
				onDownload={download}
				total={items.length}
			/>
			<Box
				borderRadius={3}
				width={'100%'}
				border={'1px solid rgba(0, 0, 0, 0.12)'}
				height={'fit-content'}
				minHeight={600}
				display={'flex'}
				flexDirection={'column'}
				sx={{ backgroundColor: '#fff', userSelect: 'none' }}
			>
				<ReportTable items={items} isFetching={isFetching} isUninitialized={isUninitialized} />
			</Box>
		</PageBox>
	)
}
