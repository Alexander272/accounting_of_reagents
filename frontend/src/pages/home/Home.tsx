import { DataTable } from '@/features/table/components/Table'
import { PageBox } from '@/styled/PageBox'
import { Box } from '@mui/material'

export default function Home() {
	return (
		<PageBox>
			<Box
				borderRadius={3}
				padding={2}
				width={'100%'}
				border={'1px solid rgba(0, 0, 0, 0.12)'}
				flexGrow={1}
				maxHeight={800}
				display={'flex'}
				flexDirection={'column'}
				sx={{ backgroundColor: '#fff', userSelect: 'none' }}
			>
				<DataTable />
			</Box>
		</PageBox>
	)
}
