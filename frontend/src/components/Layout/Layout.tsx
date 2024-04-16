import { Suspense } from 'react'
import { Outlet } from 'react-router-dom'
import { Box } from '@mui/material'

import { Fallback } from '@/components/Fallback/Fallback'
import { LayoutHeader } from './LayoutHeader'

export const Layout = () => {
	return (
		<Box minHeight='100vh' height='100vh' display='flex' flexDirection='column' pb={4}>
			<LayoutHeader />
			<Suspense fallback={<Fallback />}>
				<Outlet />
			</Suspense>
		</Box>
	)
}
