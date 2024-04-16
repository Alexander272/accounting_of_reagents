import type { RouteObject } from 'react-router-dom'
import { createBrowserRouter } from 'react-router-dom'

import { AppRoutes } from '@/constants/routes'
import { Layout } from '@/components/Layout/Layout'
import { NotFound } from '@/pages/notFound/NotFoundLazy'
import { Home } from '@/pages/home/HomeLazy'
import { Auth } from '@/pages/auth/AuthLazy'
import PrivateRoute from './PrivateRoute'

const config: RouteObject[] = [
	{
		element: <Layout />,
		errorElement: <NotFound />,
		children: [
			{
				path: AppRoutes.AUTH,
				element: <Auth />,
			},
			{
				path: AppRoutes.HOME,
				element: (
					<PrivateRoute>
						<Home />
					</PrivateRoute>
				),
			},
		],
	},
]

export const router = createBrowserRouter(config)
