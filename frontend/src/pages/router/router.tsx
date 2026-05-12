import { createBrowserRouter, type RouteObject } from 'react-router-dom'
import { AppRoutes } from '@/constants/routes'
import { Layout } from '@/components/Layout/Layout'
import { NotFound } from '@/pages/notFound/NotFoundLazy'
import { Home } from '@/pages/home/HomeLazy'
import { Auth } from '@/pages/auth/AuthLazy'
import { Accesses } from '@/pages/accesses/AccessesLazy'
import { Role } from '@/pages/accesses/role/RoleLazy'
import { Permissions } from '@/pages/accesses/permissions/PermsLazy'
import { Users } from '@/pages/accesses/users/UsersLazy'
import { Realms } from '@/pages/accesses/realms/RealmsLazy'
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
				element: <PrivateRoute />,
				children: [
					{
						index: true,
						element: <Home />,
					},

					{
						path: AppRoutes.Accesses,
						element: <Accesses />,
						children: [
							{
								index: true,
								element: <Realms />,
							},
							{
								path: AppRoutes.UserAccess,
								element: <Users />,
							},
							{
								path: AppRoutes.RoleAccess,
								element: <Role />,
							},
							{
								path: AppRoutes.Permissions,
								element: <Permissions />,
							},
						],
					},
				],
			},
		],
	},
]

export const router = createBrowserRouter(config)
