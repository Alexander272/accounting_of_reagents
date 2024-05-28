import type { PropsWithChildren } from 'react'
import { Navigate, useLocation } from 'react-router-dom'

import { AppRoutes } from '@/constants/routes'
import { useAppSelector } from '@/hooks/redux'
import { getMenu, getToken } from '@/features/user/userSlice'
import { Forbidden } from '../forbidden/ForbiddenLazy'

// проверка авторизации пользователя
export default function PrivateRoute({ children }: PropsWithChildren) {
	const token = useAppSelector(getToken)
	const menu = useAppSelector(getMenu)
	const location = useLocation()

	if (!token) return <Navigate to={AppRoutes.AUTH} state={{ from: location }} />
	if (!menu || !menu.length) return <Forbidden />

	return children
}
