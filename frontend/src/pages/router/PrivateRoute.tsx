import { Navigate, Outlet, useLocation } from 'react-router-dom'

import { AppRoutes } from '@/constants/routes'
import { useAppSelector } from '@/hooks/redux'
import { getPermissions, getToken } from '@/features/user/userSlice'
import { Forbidden } from '../forbidden/ForbiddenLazy'

// проверка авторизации пользователя
export default function PrivateRoute() {
	const token = useAppSelector(getToken)
	const perms = useAppSelector(getPermissions)
	const location = useLocation()

	if (!token) return <Navigate to={AppRoutes.AUTH} state={{ from: location }} />
	if (!perms) return <Forbidden />

	return <Outlet />
}
