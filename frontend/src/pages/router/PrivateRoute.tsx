import type { PropsWithChildren } from 'react'
import { Navigate, useLocation } from 'react-router-dom'

import { useAppSelector } from '@/hooks/redux'
import { getToken } from '@/features/user/userSlice'
import { AppRoutes } from '@/constants/routes'

// проверка авторизации пользователя
export default function PrivateRoute({ children }: PropsWithChildren) {
	const token = useAppSelector(getToken)
	const location = useLocation()

	if (!token) return <Navigate to={AppRoutes.AUTH} state={{ from: location }} />

	return children
}
